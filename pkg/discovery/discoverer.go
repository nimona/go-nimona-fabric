package discovery

import (
	"fmt"
	"sync"

	"nimona.io/internal/rand"
	"nimona.io/pkg/context"
	"nimona.io/pkg/crypto"
	"nimona.io/pkg/errors"
	"nimona.io/pkg/log"
	"nimona.io/pkg/object"
	"nimona.io/pkg/peer"
)

//go:generate $GOBIN/mockery -name Discoverer -case underscore
//go:generate $GOBIN/mockery -name Provider -case underscore

type (
	// Provider defines the interface for a discoverer provider, eg our DHT
	Provider interface {
		FindByPublicKey(
			ctx context.Context,
			key crypto.PublicKey,
			opts ...Option,
		) ([]*peer.Peer, error)
		FindByContent(
			ctx context.Context,
			contentHash *object.Hash,
			opts ...Option,
		) ([]crypto.PublicKey, error)
	}
	// Discoverer interface
	Discoverer interface {
		AddProvider(provider Provider) error
		Add(peer *peer.Peer)
		FindByPublicKey(
			ctx context.Context,
			key crypto.PublicKey,
			opts ...Option,
		) ([]*peer.Peer, error)
		FindByContent(
			ctx context.Context,
			contentHash *object.Hash,
			opts ...Option,
		) ([]crypto.PublicKey, error)
	}
)

// Options is the complete options structure for the discoverer
type Options struct {
	Local bool
}

// Option is the type for our functional options
type Option func(*Options)

// Local forces the discoverer to only look at its cache
func Local() Option {
	return func(opts *Options) {
		opts.Local = true
	}
}

func ParseOptions(opts ...Option) *Options {
	options := &Options{}
	for _, o := range opts {
		o(options)
	}
	return options
}

// NewDiscoverer creates a new empty discoverer with no providers
func NewDiscoverer() Discoverer {
	return &discoverer{
		providers:       sync.Map{},
		cacheTemp:       &StringPeerPeerSyncMap{},
		cachePersistent: &StringPeerPeerSyncMap{},
	}
}

// discoverer wraps multiple providers to allow resolving peer keys to peer infos
// TODO consider allowing the discoverer to accept an interface, and select
// the provider based on the input's type. This would require registering
// providers with the inputs they accept.
type discoverer struct {
	providers       sync.Map
	cacheTemp       *StringPeerPeerSyncMap
	cachePersistent *StringPeerPeerSyncMap
}

// FindByPublicKey goes through the given providers until one returns something
func (r *discoverer) FindByPublicKey(
	ctx context.Context,
	key crypto.PublicKey,
	opts ...Option,
) ([]*peer.Peer, error) {
	opt := ParseOptions(opts...)

	logger := log.FromContext(ctx).With(
		log.String("method", "discovery/discoverer.FindByPublicKey"),
		log.String("key", key.String()),
		log.String("opts", fmt.Sprintf("%#v", opt)),
	)

	logger.Debug("trying to find peers")

	ps := []*peer.Peer{}
	r.providers.Range(func(_, v interface{}) bool {
		p, ok := v.(Provider)
		if !ok {
			return true
		}
		eps, err := p.FindByPublicKey(ctx, key, opts...)
		if err != nil {
			logger.With(
				log.Error(err),
			).Debug("provider failed")
			return true
		}
		ps = append(ps, eps...)
		logger.With(
			log.Int("n", len(eps)),
			log.Int("total.n", len(ps)),
			log.Any("peers", ps),
		).Debug("found n peers")
		return true
	})

	// TODO move persistence into its own provider

	if res, ok := r.cacheTemp.Get(key.String()); ok && res != nil {
		ps = append(ps, res)
	}

	if res, ok := r.cachePersistent.Get(key.String()); ok && res != nil {
		ps = append(ps, res)
	}

	if len(ps) == 0 {
		return nil, errors.New("could not resolve")
	}

	return ps, nil
}

// FindByContent goes through the given providers until one returns something
func (r *discoverer) FindByContent(
	ctx context.Context,
	contentHash *object.Hash,
	opts ...Option,
) ([]crypto.PublicKey, error) {
	opt := ParseOptions(opts...)

	logger := log.FromContext(ctx).With(
		log.String("method", "discovery/discoverer.FindByContent"),
		log.String("contentHash", contentHash.Compact()),
		log.String("opts", fmt.Sprintf("%#v", opt)),
	)

	logger.Debug("trying to find peers")

	ps := []crypto.PublicKey{}
	r.providers.Range(func(_, v interface{}) bool {
		p, ok := v.(Provider)
		if !ok {
			return true
		}
		eps, err := p.FindByContent(ctx, contentHash, opts...)
		if err != nil {
			logger.With(
				log.Error(err),
			).Debug("provider failed")
			return true
		}
		ps = append(ps, eps...)
		logger.With(
			log.Int("n", len(eps)),
			log.Any("peers", ps),
		).Debug("found n peers")
		return true
	})

	if len(ps) == 0 {
		return nil, errors.New("could not resolve")
	}

	return ps, nil
}

// AddProvider to the discoverer
func (r *discoverer) AddProvider(provider Provider) error {
	r.providers.Store(rand.String(5), provider)
	return nil
}

// Add allows manually adding peer infos to be resolved.
// These peers will eventually be gc-ed.
func (r *discoverer) Add(peer *peer.Peer) {
	r.cacheTemp.Put(peer.Signature.Signer.String(), peer)
}

// AddPersistent allows adding permanent peer infos to be resolved.
// These peers can be overshadowed by other discoverers, but will never be gc-ed
// Mainly used for adding bootstrap nodes.
func (r *discoverer) AddPersistent(peer *peer.Peer) {
	r.cachePersistent.Put(peer.Signature.Signer.String(), peer)
}
