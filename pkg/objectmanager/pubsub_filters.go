package objectmanager

import (
	"fmt"

	"github.com/gobwas/glob"

	"nimona.io/pkg/chore"
	"nimona.io/pkg/crypto"
	"nimona.io/pkg/object"
)

// LookupOptions
type (
	LookupOption  func(*LookupOptions)
	LookupOptions struct {
		// Lookups are used to perform db queries for these filters
		// TODO find a better name for this
		Lookups struct {
			ObjectHashes []chore.Hash
			StreamHashes []chore.Hash
			ContentTypes []string
			Owners       []crypto.PublicKey
		}
		// filters are the lookups equivalents for matching objects for pubsub
		Filters []ObjectFilter
	}
)

func newLookupOptions(lookupOptions ...LookupOption) LookupOptions {
	options := &LookupOptions{
		Lookups: struct {
			ObjectHashes []chore.Hash
			StreamHashes []chore.Hash
			ContentTypes []string
			Owners       []crypto.PublicKey
		}{
			ObjectHashes: []chore.Hash{},
			StreamHashes: []chore.Hash{},
			ContentTypes: []string{},
			Owners:       []crypto.PublicKey{},
		},
		Filters: []ObjectFilter{},
	}
	for _, lookupOption := range lookupOptions {
		lookupOption(options)
	}
	return *options
}

func FilterByHash(hs ...chore.Hash) LookupOption {
	return func(opts *LookupOptions) {
		opts.Lookups.ObjectHashes = append(opts.Lookups.ObjectHashes, hs...)
		opts.Filters = append(opts.Filters, func(o *object.Object) bool {
			for _, h := range hs {
				if !h.IsEmpty() && o != nil && o.Hash().Equal(h) {
					return true
				}
			}
			return false
		})
	}
}

func FilterByOwner(hs ...crypto.PublicKey) LookupOption {
	return func(opts *LookupOptions) {
		opts.Lookups.Owners = append(opts.Lookups.Owners, hs...)
		opts.Filters = append(opts.Filters, func(o *object.Object) bool {
			for _, h := range hs {
				owner := o.Metadata.Owner
				if !owner.IsEmpty() && !h.IsEmpty() && owner.Equals(h) {
					return true
				}
			}
			return false
		})
	}
}

func FilterByStreamHash(hs ...chore.Hash) LookupOption {
	return func(opts *LookupOptions) {
		opts.Lookups.StreamHashes = append(opts.Lookups.StreamHashes, hs...)
		opts.Filters = append(opts.Filters, func(o *object.Object) bool {
			for _, h := range hs {
				if !h.IsEmpty() && o != nil && o.Metadata.Stream.Equal(h) {
					return true
				}
			}
			return false
		})
	}
}

func FilterByObjectType(typePatterns ...string) LookupOption {
	patterns := make([]glob.Glob, len(typePatterns))
	for i, typePattern := range typePatterns {
		g, err := glob.Compile(typePattern, '.', '/', '#')
		if err != nil {
			panic(fmt.Errorf("invalid pattern: %w", err))
		}
		patterns[i] = g
	}
	return func(opts *LookupOptions) {
		opts.Lookups.ContentTypes = append(opts.Lookups.ContentTypes, typePatterns...)
		opts.Filters = append(opts.Filters, func(o *object.Object) bool {
			for _, pattern := range patterns {
				if pattern.Match(o.Type) {
					return true
				}
			}
			return false
		})
	}
}
