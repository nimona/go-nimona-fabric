package net

import (
	"expvar"
	"io"
	"math"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"nimona.io/pkg/eventbus"

	"github.com/patrickmn/go-cache"
	"github.com/zserge/metric"

	"nimona.io/pkg/context"
	"nimona.io/pkg/errors"
	"nimona.io/pkg/keychain"
	"nimona.io/pkg/log"
	"nimona.io/pkg/peer"
)

var (
	UseUPNP = false
)

// TODO remove UseUPNP and replace with option
// nolint: gochecknoinits
func init() {
	UseUPNP, _ = strconv.ParseBool(os.Getenv("UPNP"))

	connConnOutCounter := metric.NewCounter("2m1s", "15m30s", "1h1m")
	expvar.Publish("nm:net.conn.out", connConnOutCounter)

	connConnIncCounter := metric.NewCounter("2m1s", "15m30s", "1h1m")
	expvar.Publish("nm:net.conn.in", connConnIncCounter)

	connDialCounter := metric.NewCounter("2m1s", "15m30s", "1h1m")
	expvar.Publish("nm:net.dial", connDialCounter)

	connBlacklistCounter := metric.NewCounter("2m1s", "15m30s", "1h1m")
	expvar.Publish("nm:net.conn.dial.blacklist", connBlacklistCounter)
}

type (
	// Network interface
	Network interface {
		Dial(
			ctx context.Context,
			peer *peer.Peer,
		) (*Connection, error)
		Listen(
			ctx context.Context,
			bindAddress string,
		) (Listener, error)
		Accept() (*Connection, error)
		Addresses() []string
	}
	// Option for customizing a new network
	Option func(*network)
)

// New creates a new p2p network
func New(opts ...Option) Network {
	n := &network{
		keychain: keychain.DefaultKeychain,
		eventbus: eventbus.DefaultEventbus,
		transports: map[string]Transport{
			"tcps": &tcpTransport{},
		},
		middleware:  []MiddlewareHandler{},
		listeners:   []*listener{},
		connections: make(chan *Connection),
		blacklist:   cache.New(time.Second*5, time.Second*60),
	}
	for _, opt := range opts {
		opt(n)
	}
	n.middleware = append(
		n.middleware,
		(&Handshake{keychain: n.keychain}).Handle(),
	)
	return n
}

// network allows dialing and listening for p2p connections
type network struct {
	eventbus    eventbus.Eventbus
	keychain    keychain.Keychain
	transports  map[string]Transport
	middleware  []MiddlewareHandler
	listeners   []*listener
	connections chan *Connection
	attempts    attemptsMap
	blacklist   *cache.Cache
}

// Dial to a peer and return a net.Conn or error
func (n *network) Dial(
	ctx context.Context,
	p *peer.Peer,
) (*Connection, error) {
	logger := log.FromContext(ctx).With(
		log.String("peer", p.PublicKey().String()),
		log.Strings("addresses", p.Addresses),
	)

	logger.Debug("dialing")
	expvar.Get("nm:net.dial").(metric.Metric).Add(1)

	// keep a flag on whether all addresses where blacklisted so we can return
	// an ErrMissingSignature error
	allBlacklisted := true

	// go through all addresses and try to dial them
	for _, address := range p.Addresses {
		// check if address is currently blacklisted
		if _, blacklisted := n.blacklist.Get(address); blacklisted {
			logger.Debug("address is blacklisted, skipping")
			continue
		}
		// get protocol from address
		addressType := strings.Split(address, ":")[0]
		trsp, ok := n.transports[addressType]
		if !ok {
			logger.Debug("not sure how to dial",
				log.String("type", addressType),
			)
			continue
		}

		// reset blacklist flag
		allBlacklisted = false

		// dial address
		conn, err := trsp.Dial(ctx, address)
		if err != nil {
			// blacklist address
			expvar.Get("nm:net.conn.dial.blacklist").(metric.Metric).Add(1)
			attempts, backoff := n.exponentialyBlacklist(address)
			logger.Error("could not dial address, blacklisting",
				log.Int("failedAttempts", attempts),
				log.String("backoff", backoff.String()),
				log.String("type", addressType),
				log.Error(err),
			)
			continue
		}

		// pass connection to all middleware
		conn, err = n.handleMiddleware(ctx, conn)
		if err != nil {
			continue
		}

		// at this point we consider the connection successful, so we can
		// reset the failed attempts
		n.attempts.Put(address, 0)
		n.attempts.Put(p.PublicKey().String(), 0)

		expvar.Get("nm:net.conn.out").(metric.Metric).Add(1)

		return conn, nil
	}

	err := ErrAllAddressesFailed
	if allBlacklisted {
		err = ErrAllAddressesBlacklisted
	}

	logger.Error("could not dial peer", log.Error(err))
	return nil, err
}

func (n *network) exponentialyBlacklist(k string) (int, time.Duration) {
	baseBackoff := float64(time.Second * 1)
	maxBackoff := float64(time.Minute * 10)
	attempts, _ := n.attempts.Get(k)
	attempts++
	backoff := baseBackoff * math.Pow(1.5, float64(attempts))
	if backoff > maxBackoff {
		backoff = maxBackoff
	}
	n.attempts.Put(k, attempts)
	n.blacklist.Set(k, attempts, time.Duration(backoff))
	return attempts, time.Duration(backoff)
}

func (n *network) Accept() (*Connection, error) {
	conn := <-n.connections
	return conn, nil
}

// Listen
// TODO do we need to return a listener?
func (n *network) Listen(
	ctx context.Context,
	bindAddress string,
) (Listener, error) {
	mlst := &listener{
		addresses: []string{},
		listeners: []net.Listener{},
	}
	k := n.keychain.GetPrimaryPeerKey()
	for pt, tsp := range n.transports {
		lst, err := tsp.Listen(ctx, bindAddress, k)
		if err != nil {
			return nil, err
		}

		mlst.listeners = append(mlst.listeners, lst)
		mlst.addresses = append(mlst.addresses, GetAddresses(pt, lst)...)

		n.listeners = append(n.listeners, mlst)

		for _, addr := range mlst.addresses {
			n.eventbus.Publish(eventbus.NetworkAddressAdded{
				Address: addr,
			})
		}

		// TODO goroutine never ends
		go func() {
			for {
				rawConn, err := lst.Accept()
				if err != nil {
					// we need to check whether the error is temporary,
					// a non-temporary error would be for example a closed
					// listener
					errIsTemp := true
					if opErr, ok := err.(*net.OpError); ok {
						errIsTemp = opErr.Temporary()
					}
					// if the error is not temporary stop trying to accept
					// connections from this lsitener
					if !errIsTemp {
						return
					}
					// else, just move on
					continue
				}

				conn := newConnection(rawConn, true)
				conn.remoteAddress = rawConn.RemoteAddr().String()
				conn.localAddress = rawConn.LocalAddr().String()

				conn, err = n.handleMiddleware(ctx, conn)
				if err != nil {
					continue
				}

				n.eventbus.Publish(eventbus.PeerConnectionEstablished{
					PublicKey: conn.RemotePeerKey,
				})

				expvar.Get("nm:net.conn.in").(metric.Metric).Add(1)
				n.connections <- conn
			}
		}()
	}
	return mlst, nil
}

func (n *network) Addresses() []string {
	addrs := []string{}
	for _, l := range n.listeners {
		addrs = append(addrs, l.Addresses()...)
	}
	return addrs
}

func (n *network) handleMiddleware(
	ctx context.Context,
	conn *Connection,
) (*Connection, error) {
	var err error
	for _, mh := range n.middleware {
		conn, err = mh(ctx, conn)
		if err != nil {
			if errors.CausedBy(err, io.EOF) {
				break
			}
			if conn != nil {
				conn.conn.Close() // nolint: errcheck
			}
			return nil, err
		}
	}
	return conn, nil
}