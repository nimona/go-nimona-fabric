package discovery

import (
	"fmt"
	"sync"

	"go.uber.org/zap"

	"nimona.io/internal/context"
	"nimona.io/internal/errors"
	"nimona.io/internal/log"
	"nimona.io/pkg/net/peer"
)

// nolint: lll
//go:generate go run github.com/vektra/mockery/cmd/mockery -name Discoverer -case underscore

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

// Discoverer interface
type Discoverer interface {
	AddProvider(provider Provider) error
	Add(v *peer.PeerInfo)
	// AddPersistent(v *peer.PeerInfo)
	FindByFingerprint(
		ctx context.Context,
		fingerprint string,
		opts ...Option,
	) ([]*peer.PeerInfo, error)
	FindByContent(
		ctx context.Context,
		contentHash string,
		opts ...Option,
	) ([]*peer.PeerInfo, error)
}

// NewDiscoverer creates a new empty discoverer with no providers
func NewDiscoverer() Discoverer {
	return &discoverer{
		providersLock:   sync.RWMutex{},
		providers:       []Provider{},
		cacheLock:       sync.RWMutex{},
		cacheTemp:       map[string]*peer.PeerInfo{},
		cachePersistent: map[string]*peer.PeerInfo{},
	}
}

// discoverer wraps multiple providers to allow resolving peer keys to peer infos
// TODO consider allowing the discoverer to accept an interface, and select
// the provider based on the input's type. This would require registering
// providers with the inputs they accept.
type discoverer struct {
	providersLock   sync.RWMutex
	providers       []Provider
	cacheLock       sync.RWMutex
	cacheTemp       map[string]*peer.PeerInfo
	cachePersistent map[string]*peer.PeerInfo
}

// FindByFingerprint goes through the given providers until one returns something
func (r *discoverer) FindByFingerprint(
	ctx context.Context,
	fingerprint string,
	opts ...Option,
) ([]*peer.PeerInfo, error) {
	opt := ParseOptions(opts...)

	logger := log.Logger(ctx).With(
		zap.String("method", "discovery/discoverer.FindByFingerprint"),
		zap.String("fingerprint", fingerprint),
		zap.String("opts", fmt.Sprintf("%#v", opt)),
	)

	logger.Debug("trying to find peers")

	// r.providersLock.RLock()
	// defer r.providersLock.RUnlock()

	for _, p := range r.providers {
		res, err := p.FindByFingerprint(ctx, fingerprint, opts...)
		if err == nil && res != nil {

			return res, nil
		}
	}

	// r.cacheLock.RLock()
	// defer r.cacheLock.RUnlock()

	// TODO move persistence into its own provider

	if res, ok := r.cacheTemp[fingerprint]; ok && res != nil {
		return []*peer.PeerInfo{res}, nil
	}

	if res, ok := r.cachePersistent[fingerprint]; ok && res != nil {
		return []*peer.PeerInfo{res}, nil
	}

	return nil, errors.New("could not resolve")
}

// FindByContent goes through the given providers until one returns something
func (r *discoverer) FindByContent(
	ctx context.Context,
	contentHash string,
	opts ...Option,
) ([]*peer.PeerInfo, error) {
	opt := ParseOptions(opts...)

	logger := log.Logger(ctx).With(
		zap.String("method", "discovery/discoverer.FindByContent"),
		zap.String("contentHash", contentHash),
		zap.String("opts", fmt.Sprintf("%#v", opt)),
	)

	logger.Debug("trying to find peers")

	r.providersLock.RLock()
	for _, p := range r.providers {
		res, err := p.FindByContent(ctx, contentHash, opts...)
		if err == nil && res != nil {
			r.providersLock.RUnlock()
			return res, nil
		}
	}
	r.providersLock.RUnlock()

	return nil, errors.New("could not resolve")
}

// AddProvider to the discoverer
func (r *discoverer) AddProvider(provider Provider) error {
	r.providersLock.Lock()
	r.providers = append(r.providers, provider)
	r.providersLock.Unlock()
	return nil
}

// Add allows manually adding peer infos to be resolved.
// These peers will eventually be gc-ed.
func (r *discoverer) Add(v *peer.PeerInfo) {
	r.cacheLock.Lock()
	r.cacheTemp[v.Fingerprint()] = v
	r.cacheLock.Unlock()
}

// AddPersistent allows adding permanent peer infos to be resolved.
// These peers can be overshadowed by other discoverers, but will never be gc-ed
// Mainly used for adding bootstrap nodes.
func (r *discoverer) AddPersistent(v *peer.PeerInfo) {
	r.cacheLock.Lock()
	r.cachePersistent[v.Fingerprint()] = v
	r.cacheLock.Unlock()
}
