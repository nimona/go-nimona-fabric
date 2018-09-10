package net

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"reflect"
	"strings"
	"sync"
	"time"

	cbor "github.com/ugorji/go/codec"
	"go.uber.org/zap"

	"github.com/nimona/go-nimona/blocks"
	"github.com/nimona/go-nimona/codec"
	"github.com/nimona/go-nimona/log"
	"github.com/nimona/go-nimona/peers"
	"github.com/nimona/go-nimona/storage"
	"github.com/nimona/go-nimona/utils"
)

var (
	// ErrInvalidRequest when received an invalid request block
	ErrInvalidRequest = errors.New("Invalid request")
)

var (
	// ErrAllAddressesFailed for when a peer cannot be dialed
	ErrAllAddressesFailed = errors.New("all addresses failed to dial")
	// ErrNoAddresses for when a peer has no addresses
	ErrNoAddresses = errors.New("no addresses")
	// ErrNotForUs block is not meant for us
	ErrNotForUs = errors.New("block not for us")
)

// Exchange interface for mocking exchange
type Exchange interface {
	Get(ctx context.Context, id string) (interface{}, error)
	GetLocalBlocks() ([]string, error)
	Handle(contentType string, h func(o interface{}) error) error
	Send(ctx context.Context, o interface{}, recipient *blocks.Key, opts ...blocks.MarshalOption) error
	Listen(ctx context.Context, addrress string) (net.Listener, error)
	RegisterDiscoverer(discovery Discoverer)
	// Verify(block *blocks.Block) error
	// Sign(block *blocks.Block, signerPeerInfo *peers.PrivatePeerInfo) error
}

type exchange struct {
	network     Networker
	addressBook *peers.AddressBook
	discovery   Discoverer

	outgoingPayloads chan outBlock
	incoming         chan net.Conn
	outgoing         chan net.Conn
	close            chan bool

	streams    sync.Map
	handlers   []handler
	logger     *zap.Logger
	streamLock utils.Kmutex

	store         storage.Storage
	getRequests   sync.Map
	subscriptions sync.Map
}

type outBlock struct {
	recipient *blocks.Key
	bytes     []byte
}

type incBlock struct {
	peerID  string
	conn    net.Conn
	payload []byte
}

type handler struct {
	contentType string
	handler     func(o interface{}) error
}

// NewExchange creates a exchange on a given network
func NewExchange(addressBook *peers.AddressBook, store storage.Storage) (Exchange, error) {
	ctx := context.Background()

	network, err := NewNetwork(addressBook)
	if err != nil {
		return nil, err
	}

	w := &exchange{
		network:     network,
		addressBook: addressBook,

		outgoingPayloads: make(chan outBlock, 100),
		incoming:         make(chan net.Conn),
		outgoing:         make(chan net.Conn),
		close:            make(chan bool),

		handlers:   []handler{},
		logger:     log.Logger(ctx).Named("exchange"),
		streamLock: utils.NewKmutex(),

		store:       store,
		getRequests: sync.Map{},
	}

	self := w.addressBook.GetLocalPeerInfo()

	go func() {
		for block := range w.outgoingPayloads {
			recipientThumbPrint := block.recipient.Thumbprint()
			if self.Thumbprint() == recipientThumbPrint {
				w.logger.Info("cannot send block to self")
				continue
			}

			// TODO log error and reconsider the async
			// TODO also maybe we need to verify it or something?

			logger := w.logger.With(zap.String("peerID", recipientThumbPrint))

			// try to send the block directly to the recipient
			logger.Debug("getting conn to write block")
			conn, err := w.GetOrDial(ctx, recipientThumbPrint)
			if err != nil {
				logger.Debug("could not get conn to recipient", zap.Error(err))
			} else {
				if err := w.writeBlock(ctx, block.bytes, conn); err != nil {
					// TODO better handling of connection errors
					w.Close(recipientThumbPrint, conn)
					logger.Debug("could not write to recipient", zap.Error(err))
				} else {
					// update peer status
					w.addressBook.PutPeerStatus(recipientThumbPrint, peers.StatusConnected)
					continue
				}
			}

			// else try to send message via their relay addresses
			conn, err = w.getOrDialRelay(ctx, recipientThumbPrint)
			if err != nil {
				logger.Debug("could not get conn to recipient's relay", zap.Error(err))
				continue
			}

			// create forwarded block
			fw := ForwardRequest{
				Recipient: block.recipient,
				Block:     block.bytes,
			}

			fwb, err := blocks.Marshal(fw, blocks.SignWith(self.Key))
			if err != nil {
				panic(err)
				continue
			}

			// try to send the block directly to the recipient
			if err := w.writeBlock(ctx, fwb, conn); err != nil {
				// TODO better handling of connection errors
				// TODO this is a bad close, id is of recipient, conn is of relay
				w.Close(recipientThumbPrint, conn)
				logger.Debug("could not write to relay", zap.Error(err))
				// update peer status
				w.addressBook.PutPeerStatus(recipientThumbPrint, peers.StatusError)
				continue
			}

			// update peer status
			w.addressBook.PutPeerStatus(recipientThumbPrint, peers.StatusCanConnect)
		}
	}()

	return w, nil
}

func (w *exchange) RegisterDiscoverer(discovery Discoverer) {
	w.discovery = discovery

	ctx := context.Background()
	go func() {
		for {
			blocks, err := w.store.List()
			if err != nil {
				time.Sleep(time.Second * 10)
				continue
			}

			for _, block := range blocks {
				if err := w.discovery.PutProviders(ctx, block); err != nil {
					w.logger.Warn("could not announce provider for block", zap.String("id", block))
				}
			}

			time.Sleep(time.Second * 30)
		}
	}()
}

func (w *exchange) Handle(contentType string, h func(o interface{}) error) error {
	w.handlers = append(w.handlers, handler{
		contentType: contentType,
		handler:     h,
	})
	return nil
}

func (w *exchange) Close(peerID string, conn net.Conn) {
	if conn != nil {
		conn.Close()
	}
	w.streams.Range(func(k, v interface{}) bool {
		if k.(string) == peerID {
			w.streams.Delete(k)
		}
		if v.(net.Conn) == conn {
			w.streams.Delete(k)
		}
		return true
	})
}

func (w *exchange) HandleConnection(conn net.Conn) error {
	w.logger.Debug("handling new connection", zap.String("remote", conn.RemoteAddr().String()))

	blockDecoder := cbor.NewDecoder(conn, codec.CborHandler())
	for {
		block := map[string]interface{}{}
		if err := blockDecoder.Decode(&block); err != nil {
			w.logger.Error("could not read block", zap.Error(err))
			w.Close("", conn)
			return err
		}

		blockBytes := []byte{}
		enc := cbor.NewEncoderBytes(&blockBytes, blocks.CborHandler())
		if err := enc.Encode(block); err != nil {
			panic(err)
		}
		if err := w.process(blockBytes, conn); err != nil {
			w.Close("", conn)
			return err
		}
	}
}

// Process incoming block
func (w *exchange) process(blockBytes []byte, conn net.Conn) error {
	ib, err := blocks.Unmarshal(blockBytes, blocks.Verify(), blocks.ReturnBlock())
	if err != nil {
		panic(err)
	}

	block := ib.(*blocks.Block)

	if os.Getenv("DEBUG_BLOCKS") != "" {
		fmt.Println("Processing type", block.Type, "as", reflect.TypeOf(block.Payload))
		fmt.Println("< ---------- inc block / start")
		b, _ := json.MarshalIndent(block, "< ", "  ")
		fmt.Println(string(b))
		fmt.Println("< ---------- inc block / end")
	}

	if os.Getenv("TELEMETRY") == "client" {
		SendBlockEvent(
			false,
			block.Type,
			0, // len(GetRecipientsFromBlockPolicies(block)),
			0, // TODO fix payload size
			len(blockBytes),
		)
	}

	blockID := block.ID()
	if blocks.ShouldPersist(block.Type) {
		if err := w.store.Store(blockID, blockBytes); err != nil {
			if err != storage.ErrExists {
				w.logger.Warn("could not write block", zap.Error(err))
			}
		}
	}

	// TODO convert these into proper handlers
	contentType := block.Type
	switch payload := block.Payload.(type) {
	case *ForwardRequest:
		w.logger.Info("got forwarded message", zap.String("recipient", payload.Recipient.Thumbprint()))
		w.outgoingPayloads <- outBlock{
			recipient: payload.Recipient,
			bytes:     payload.Block,
		}
		return nil

	case *BlockRequest:
		if err := w.handleRequestBlock(payload); err != nil {
			w.logger.Warn("could not handle request block", zap.Error(err))
		}

	case *BlockResponse:
		if err := w.handleBlockResponse(payload); err != nil {
			w.logger.Warn("could not handle request block", zap.Error(err))
		}

	case *HandshakePayload:
		if err := w.addressBook.PutPeerInfo(payload.PeerInfo); err != nil {
			return err
		}

		w.streams.Store(payload.PeerInfo.Thumbprint(), conn)
		return nil
	}

	for _, handler := range w.handlers {
		if strings.HasPrefix(contentType, handler.contentType) {
			if err := handler.handler(block.Payload); err != nil {
				w.logger.Info(
					"Could not handle event",
					zap.String("contentType", contentType),
					zap.Error(err),
				)
				return err
			}
		}
	}

	return nil
}

func (w *exchange) handleBlockResponse(payload *BlockResponse) error {
	// Check if nonce exists in local addressBook
	value, ok := w.getRequests.Load(payload.RequestID)
	if !ok {
		return nil
	}

	req, ok := value.(*BlockRequest)
	if !ok {
		return ErrInvalidRequest
	}

	block, err := blocks.Unmarshal(payload.Block, blocks.ReturnBlock())
	if err != nil {
		panic(err)
		return err
	}

	if blocks.ShouldPersist(block.(*blocks.Block).Type) {
		blockID, _ := blocks.SumSha3(payload.Block)
		w.store.Store(blockID, payload.Block)
	}

	req.response <- block.(*blocks.Block).Payload

	return nil
}

func (w *exchange) handleRequestBlock(payload *BlockRequest) error {
	blockBytes, err := w.store.Get(payload.ID)
	if err != nil {
		return err
	}

	// TODO check if policy allows requested to retrieve the block

	resp := &BlockResponse{
		RequestID: payload.RequestID,
		Block:     blockBytes,
	}

	signer := w.addressBook.GetLocalPeerInfo().Key
	if err := w.Send(context.Background(), resp, payload.Signature.Key, blocks.SignWith(signer)); err != nil {
		w.logger.Warn("blx.handleRequestBlock could not send block", zap.Error(err))
		return err
	}

	return nil
}

func (w *exchange) Get(ctx context.Context, id string) (interface{}, error) {
	// Check local storage for block
	if blockBytes, err := w.store.Get(id); err == nil {
		return blocks.Unmarshal(blockBytes)
	}

	req := &BlockRequest{
		RequestID: RandStringBytesMaskImprSrc(8),
		ID:        id,
		response:  make(chan interface{}),
	}

	defer func() {
		w.getRequests.Delete(req.RequestID)
		close(req.response)
	}()

	w.getRequests.Store(req.RequestID, req)
	signer := w.addressBook.GetLocalPeerInfo().Key

	go func() {
		providers, err := w.discovery.GetProviders(ctx, id)
		if err != nil {
			// TODO log err
			return
		}

		for provider := range providers {
			if err := w.Send(ctx, req, provider, blocks.SignWith(signer)); err != nil {
				w.logger.Warn("blx.Get could not send req block", zap.Error(err))
			}
		}
	}()

	for {
		select {
		case payload := <-req.response:
			return payload, nil

		case <-ctx.Done():
			return nil, storage.ErrNotFound
		}
	}
}

func (w *exchange) Send(ctx context.Context, o interface{}, recipient *blocks.Key, opts ...blocks.MarshalOption) error {
	// b := blocks.Encode(o)
	// if err := blocks.Sign(b, w.addressBook.GetLocalPeerInfo().Key); err != nil {
	// 	return err
	// }

	bytes, err := blocks.Marshal(o, opts...)
	if err != nil {
		return err
	}

	if os.Getenv("DEBUG_BLOCKS") != "" {
		fmt.Print("> ---------- out block / start ")
		b, _ := json.MarshalIndent(o, "> ", "  ")
		fmt.Print(string(b))
		fmt.Println(" ---------- out block / end")
	}

	if blocks.ShouldPersist(blocks.GetFromType(reflect.TypeOf(o))) {
		blockID, _ := blocks.SumSha3(bytes)
		w.store.Store(blockID, bytes)
	}

	// TODO right now there is no way to error on this, do we have to?
	w.outgoingPayloads <- outBlock{
		recipient: recipient,
		bytes:     bytes,
	}

	return nil
}

func (w *exchange) GetLocalBlocks() ([]string, error) {
	return w.store.List()
}

func (w *exchange) writeBlock(ctx context.Context, bytes []byte, rw io.ReadWriter) error {
	if _, err := rw.Write(bytes); err != nil {
		return err
	}

	w.logger.Debug("writing block", zap.String("bytes", blocks.Base58Encode(bytes)))

	return nil
}

func (w *exchange) getOrDialRelay(ctx context.Context, peerID string) (net.Conn, error) {
	peer, err := w.addressBook.GetPeerInfo(peerID)
	if err != nil {
		return nil, err
	}

	for _, address := range peer.Addresses {
		// TODO better check
		if strings.HasPrefix(address, "relay:") {
			relayPeerID := strings.Replace(address, "relay:", "", 1)
			conn, err := w.GetOrDial(ctx, relayPeerID)
			if err != nil {
				continue
			}
			return conn, nil
		}
	}

	return nil, ErrAllAddressesFailed
}

func (w *exchange) GetOrDial(ctx context.Context, peerID string) (net.Conn, error) {
	w.logger.Debug("looking for existing connection", zap.String("peer_id", peerID))
	if peerID == "" {
		return nil, errors.New("missing peer id")
	}

	existingConn, ok := w.streams.Load(peerID)
	if ok {
		return existingConn.(net.Conn), nil
	}

	conn, err := w.network.Dial(ctx, peerID)
	if err != nil {
		w.Close(peerID, conn)
		return nil, err
	}

	// TODO move after handshake
	// handle outgoing connections
	w.outgoing <- conn

	// store conn for reuse
	w.streams.Store(peerID, conn)

	w.logger.Debug("writing handshake")

	// handshake so the other side knows who we are
	// TODO can't ths use Send()?
	handshake := HandshakePayload{
		PeerInfo: w.addressBook.GetLocalPeerInfo().GetPeerInfo(),
	}

	signer := w.addressBook.GetLocalPeerInfo().Key
	handshakeBytes, err := blocks.Marshal(handshake, blocks.SignWith(signer))
	if err != nil {
		panic(err)
		return nil, err
	}

	if err := w.writeBlock(ctx, handshakeBytes, conn); err != nil {
		w.Close(peerID, conn)
		panic(err)
		return nil, err
	}

	return conn, nil
}

// Listen on an address
// TODO do we need to return a listener?
func (w *exchange) Listen(ctx context.Context, addr string) (net.Listener, error) {
	listener, err := w.network.Listen(ctx, addr)
	if err != nil {
		return nil, err
	}

	closed := false

	go func() {
		for {
			select {
			case conn := <-w.incoming:
				go func() {
					if err := w.HandleConnection(conn); err != nil {
						w.logger.Warn("failed to handle block", zap.Error(err))
					}
				}()
			case conn := <-w.outgoing:
				go func() {
					if err := w.HandleConnection(conn); err != nil {
						w.logger.Warn("failed to handle block", zap.Error(err))
					}
				}()
			case <-w.close:
				closed = true
				w.logger.Debug("connection closed")
				listener.Close()
			}
		}
	}()

	go func() {
		w.logger.Debug("accepting connections", zap.String("address", listener.Addr().String()))
		for {
			conn, err := listener.Accept()
			w.logger.Debug("connection accepted")
			if err != nil {
				if closed {
					return
				}
				w.logger.Error("could not accept", zap.Error(err))
				// TODO check conn is still alive and return
				return
			}
			w.incoming <- conn
		}
	}()

	return listener, nil
}
