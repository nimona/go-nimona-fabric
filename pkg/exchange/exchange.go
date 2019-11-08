package exchange

import (
	"strings"
	"sync"

	"github.com/gobwas/glob"
	"github.com/sheerun/queue"

	"nimona.io/internal/rand"
	"nimona.io/internal/store/graph"
	"nimona.io/pkg/context"
	"nimona.io/pkg/crypto"
	"nimona.io/pkg/discovery"
	"nimona.io/pkg/errors"
	"nimona.io/pkg/log"
	"nimona.io/pkg/net"
	"nimona.io/pkg/object"
	"nimona.io/pkg/peer"
)

var objectRequestType = new(ObjectRequest).GetType()

const (
	// ErrInvalidRequest when received an invalid request object
	ErrInvalidRequest = errors.Error("invalid request")
	// ErrSendingTimedOut when sending times out
	ErrSendingTimedOut = errors.Error("sending timed out")
	// errOutboxForwarded when an object is forewarded to a different outbox
	// this usually happens when an existing connection already existed
	errOutboxForwarded = errors.Error("request has been moved to another outbox")
)

// nolint: lll
//go:generate $GOBIN/mockery -case underscore -inpkg -name Exchange
//go:generate $GOBIN/genny -in=$GENERATORS/syncmap_named/syncmap.go -out=addresses.go -pkg exchange gen "KeyType=string ValueType=addressState SyncmapName=addresses"
//go:generate $GOBIN/genny -in=$GENERATORS/syncmap_named/syncmap.go -out=inboxes.go -pkg exchange gen "KeyType=string ValueType=inbox SyncmapName=inboxes"
//go:generate $GOBIN/genny -in=$GENERATORS/syncmap_named/syncmap.go -out=outboxes.go -pkg exchange gen "KeyType=crypto.Fingerprint ValueType=outbox SyncmapName=outboxes"

type (
	// Exchange interface for mocking exchange
	Exchange interface {
		Request(
			ctx context.Context,
			object *object.Hash,
			address string,
			options ...Option,
		) error
		Handle(
			contentTypeGlob string,
			handler func(object *Envelope) error,
		) (
			cancelationFunc func(),
			err error,
		)
		Send(
			ctx context.Context,
			object object.Object,
			address string,
			options ...Option,
		) error
	}
	// echange implements an Exchange
	exchange struct {
		key *crypto.PrivateKey
		net net.Network

		discover discovery.Discoverer
		local    *peer.LocalPeer

		outboxes *OutboxesMap
		inboxes  *InboxesMap

		store graph.Store // TODO remove
	}
	// Options (mostly) for Send()
	Options struct {
		LocalDiscovery bool
		Async          bool
	}
	Option func(*Options)
	// inbox holds a single handler, and the messages for it.
	// every registered handler will have one inbox.
	// the queue should only hold `*Envelope`s.
	inbox struct {
		contentType glob.Glob
		handler     func(*Envelope) error
		queue       *queue.Queue
	}
	// addressState defines the states of a peer's address
	// current options are:
	// * -1 unconnectable
	// * 0 unknown
	// * 1 connectable
	// * 2 blacklisted
	addressState int
	// outbox holds information about a single peer, its open connection,
	// and the messages for it.
	// the queue should only hold `*outgoingObject`s.
	outbox struct {
		peer      crypto.Fingerprint
		addresses *AddressesMap
		conn      *net.Connection
		connLock  sync.RWMutex
		queue     *queue.Queue
	}
	// outgoingObject holds an object that is about to be sent
	outgoingObject struct {
		context   context.Context
		recipient string
		object    object.Object
		options   *Options
		err       chan error
	}
)

// New creates a exchange on a given network
func New(
	ctx context.Context,
	key *crypto.PrivateKey,
	n net.Network,
	store graph.Store,
	discover discovery.Discoverer,
	localInfo *peer.LocalPeer,
) (
	Exchange,
	error,
) {
	w := &exchange{
		key: key,
		net: n,

		discover: discover,
		local:    localInfo,

		outboxes: NewOutboxesMap(),
		inboxes:  NewInboxesMap(),

		store: store,
	}

	// TODO(superdecimal) we should probably remove .Listen() from here, net
	// should have a function that accepts a connection handler or something.
	incomingConnections, err := w.net.Listen(ctx)
	if err != nil {
		return nil, err
	}

	logger := log.
		FromContext(ctx).
		Named("exchange").
		With(
			log.String("method", "exchange.New"),
			log.String("local.fingerprint", localInfo.
				GetFingerprint().
				String(),
			),
		)

	// add request object handler
	w.Handle(objectRequestType, w.handleObjectRequest) // nolint: errcheck

	// handle new incoming connections
	go func() {
		for {
			conn := <-incomingConnections
			go func(conn *net.Connection) {
				if err := w.handleConnection(conn); err != nil {
					logger.Warn("failed to handle connection", log.Error(err))
				}
			}(conn)
		}
	}()

	return w, nil
}

func (w *exchange) createInbox(
	handlerID string,
	contentType glob.Glob,
	handler func(o *Envelope) error,
) (
	func(),
	error,
) {
	inbox := &inbox{
		contentType: contentType,
		handler:     handler,
		queue:       queue.New(),
	}
	close := func() {
		w.inboxes.Delete(handlerID)
	}
	go w.processInbox(inbox)
	w.inboxes.Put(handlerID, inbox)
	return close, nil
}

func (w *exchange) processInbox(inbox *inbox) {
	logger := log.DefaultLogger.
		With(
			log.String("method", "exchange.processInbox"),
		)
	for {
		v := inbox.queue.Pop()
		e := v.(*Envelope)
		// TODO(geoah) validate payload and sender
		ct := e.Payload.GetType()
		if !inbox.contentType.Match(ct) {
			continue
		}
		go func(handler func(*Envelope) error, e *Envelope) {
			defer func() {
				if r := recover(); r != nil {
					logger.
						With(
							log.Stack(),
						).
						Error("Recovered while handling", log.Any("r", r))
				}
			}()
			if err := handler(e); err != nil {
				logger.Info(
					"Could not handle event",
					log.String("contentType", ct),
					log.Error(err),
				)
			}
		}(inbox.handler, e)
	}
}

func (w *exchange) getInbox(handlerID string) *inbox {
	inbox, _ := w.inboxes.Get(handlerID)
	return inbox
}

func (w *exchange) writeToInboxes(e *Envelope) {
	w.inboxes.Range(func(handlerID string, inbox *inbox) bool {
		inbox.queue.Append(e)
		return true
	})
}

func (w *exchange) getOutbox(peer crypto.Fingerprint) *outbox {
	outbox := &outbox{
		peer:      peer,
		addresses: NewAddressesMap(),
		queue:     queue.New(),
	}
	outbox, loaded := w.outboxes.GetOrPut(peer, outbox)
	if !loaded {
		go w.processOutbox(outbox)
	}
	return outbox
}

func (w *exchange) updateOutboxConn(outbox *outbox, conn *net.Connection) {
	outbox.connLock.Lock()
	if outbox.conn != nil {
		outbox.conn.Close() // nolint: errcheck
	}
	outbox.conn = conn
	outbox.connLock.Unlock()
}

func (w *exchange) updateOutboxConnIfEmpty(
	outbox *outbox,
	conn *net.Connection,
) bool {
	outbox.connLock.Lock()
	if outbox.conn == nil {
		outbox.conn = conn
		return true
	}
	outbox.connLock.Unlock()
	return false
}

func (w *exchange) processOutbox(outbox *outbox) {
	getConnection := func(req *outgoingObject) (*net.Connection, error) {
		outbox.connLock.RLock()
		if outbox.conn != nil {
			outbox.connLock.RUnlock()
			return outbox.conn, nil
		}
		outbox.connLock.RUnlock()
		netOpts := []net.Option{}
		// if req.options.LocalDiscovery {
		netOpts = append(netOpts, net.WithLocalDiscoveryOnly())
		// }
		conn, err := w.net.Dial(req.context, req.recipient, netOpts...)
		if err != nil {
			return nil, err
		}
		// if the remote peer doesn't match the one on our outbox
		// this check is mainly done for when the outbox.peer is actually an
		// address, ie tcps:0.0.0.0:0
		// once we have the real remote peer, we should be replacing the outbox
		if conn.RemotePeerKey.Fingerprint().String() != outbox.peer.String() {
			// check if we already have an outbox with the new peer
			existingOutbox, outboxExisted := w.outboxes.GetOrPut(
				conn.RemotePeerKey.Fingerprint(),
				outbox,
			)
			// and if so
			if outboxExisted {
				// enqueue the object on that outbox
				existingOutbox.queue.Append(req)
				// try to update the connection if its gone
				updated := w.updateOutboxConnIfEmpty(existingOutbox, conn)
				// close the connection if we are not using it
				if !updated {
					conn.Close() // nolint: errcheck
					w.updateOutboxConn(outbox, existingOutbox.conn)
				}
				// and finally return errOutboxForwarded so caller knows to exit
				return nil, errOutboxForwarded
			}
		}
		if err := w.handleConnection(conn); err != nil {
			w.updateOutboxConn(outbox, nil)
			log.DefaultLogger.Warn(
				"failed to handle outbox connection",
				log.Error(err),
			)
			return nil, err
		}
		return conn, nil
	}
	for {
		// dequeue the next item to send
		// TODO figure out what can go wrong here
		v := outbox.queue.Pop()
		req := v.(*outgoingObject)
		// check if the context for this is done
		if err := req.context.Err(); err != nil {
			req.err <- err
			continue
		}
		// make a logger from our req context
		logger := log.FromContext(req.context).With(
			log.String("recipient", req.recipient),
			log.String("object.@type", req.object.GetType()),
		)
		// validate req
		if req.recipient == "" {
			logger.Info("missing recipient")
			req.err <- errors.New("missing recipient")
			continue
		}
		// try to send the object
		var lastErr error
		for i := 0; i < 3; i++ {
			logger.Debug("trying to get connection", log.Int("attempt", i+1))
			conn, err := getConnection(req)
			if err != nil {
				// the object has been forwarded to another outbox
				if err == errOutboxForwarded {
					return
				}
				lastErr = err
				w.updateOutboxConn(outbox, nil)
				continue
			}
			logger.Debug("trying write object", log.Int("attempt", i+1))
			if err := net.Write(req.object, conn); err != nil {
				lastErr = err
				w.updateOutboxConn(outbox, nil)
				continue
			}
			lastErr = nil
			break
		}
		if lastErr == nil {
			logger.Debug("wrote object")
		}
		// errOutboxForwarded are not considered errors here
		// else, we should report back with something
		if lastErr != errOutboxForwarded {
			req.err <- lastErr
		}
		continue
	}
}

// Request an object given its hash from an address
func (w *exchange) Request(
	ctx context.Context,
	hash *object.Hash,
	address string,
	options ...Option,
) error {
	req := &ObjectRequest{
		ObjectHash: hash,
	}
	return w.Send(ctx, req.ToObject(), address, options...)
}

// Handle allows registering a callback function to handle incoming objects
func (w *exchange) Handle(
	typePatern string,
	h func(o *Envelope) error,
) (
	func(),
	error,
) {
	g, err := glob.Compile(typePatern, '.', '/', '#')
	if err != nil {
		return nil, err
	}
	hID := rand.String(8)
	return w.createInbox(hID, g, h)
}

func (w *exchange) handleConnection(
	conn *net.Connection,
) error {
	if conn == nil {
		// TODO should this be nil?
		panic(errors.New("missing connection"))
	}

	// get outbox and update the connection to the peer
	outbox := w.getOutbox(conn.RemotePeerKey.Fingerprint())
	w.updateOutboxConn(outbox, conn)

	// TODO(geoah) this looks like a hack
	if err := net.Write(
		w.local.GetSignedPeer().ToObject(),
		conn,
	); err != nil {
		return err
	}

	go func() {
		for {
			payload, err := net.Read(conn)
			// TODO split errors into connection or payload
			// ie a payload that cannot be unmarshalled or verified
			// should not kill the connection
			if err != nil {
				log.DefaultLogger.Warn(
					"failed to read from connection",
					log.Error(err),
				)
				w.updateOutboxConn(outbox, nil)
				return
			}

			w.writeToInboxes(&Envelope{
				Sender:  conn.RemotePeerKey,
				Payload: payload,
			})
		}
	}()

	return nil
}

// handleObjectRequest -
func (w *exchange) handleObjectRequest(
	e *Envelope,
) error {
	// TODO verify signature
	switch e.Payload.GetType() {
	case objectRequestType:
		req := &ObjectRequest{}
		if err := req.FromObject(e.Payload); err != nil {
			return err
		}
		res, err := w.store.Get(req.ObjectHash.Compact())
		if err != nil {
			return errors.Wrap(
				errors.Error("could not retrieve object"),
				err,
			)
		}
		go w.Send( // nolint: errcheck
			context.New(),
			res,
			"peer:"+e.Sender.Fingerprint().String(),
			WithLocalDiscoveryOnly(),
		)
		return nil
	}

	return nil
}

// WithLocalDiscoveryOnly will only use local discovery to resolve addresses.
func WithLocalDiscoveryOnly() Option {
	return func(opt *Options) {
		opt.LocalDiscovery = true
	}
}

// WithAsync will not wait to actually send the object
func WithAsync() Option {
	return func(opt *Options) {
		opt.Async = true
	}
}

// Send an object to an address
func (w *exchange) Send(
	ctx context.Context,
	oo object.Object,
	address string,
	options ...Option,
) error {
	ctx = context.FromContext(ctx)
	opts := &Options{}
	for _, option := range options {
		option(opts)
	}
	o := object.Copy(oo) // TODO do we really need to copy?
	// if !strings.HasPrefix(address, "peer:") {
	// 	panic("NO PEER PREFIX")
	// }
	fingerprint := strings.Replace(address, "peer:", "", 1)
	outbox := w.getOutbox(crypto.Fingerprint(fingerprint))
	errRecv := make(chan error, 1)
	req := &outgoingObject{
		context:   ctx,
		recipient: address,
		object:    o,
		options:   opts,
		err:       errRecv,
	}
	outbox.queue.Append(req)
	if opts.Async {
		return nil
	}
	select {
	case <-ctx.Done():
		return ErrSendingTimedOut
	case err := <-errRecv:
		return err
	}
}
