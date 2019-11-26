// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package exchange

import (
	"nimona.io/internal/pubsub"
)

type (
	envelope string // nolint
	// EnvelopePubSub -
	EnvelopePubSub interface {
		Publish(*Envelope)
		Subscribe(...EnvelopeFilter) EnvelopeSubscription
	}
	EnvelopeFilter func(*Envelope) bool
	// EnvelopeSubscription is returned for every subscription
	EnvelopeSubscription interface {
		Next() (*Envelope, error)
		Cancel()
	}
	psEnvelopeSubscription struct {
		subscription pubsub.Subscription
	}
	psEnvelope struct {
		pubsub pubsub.PubSub
	}
)

// NewEnvelope constructs and returns a new EnvelopePubSub
func NewEnvelopePubSub() EnvelopePubSub {
	return &psEnvelope{
		pubsub: pubsub.New(),
	}
}

// Cancel the subscription
func (s *psEnvelopeSubscription) Cancel() {
	s.subscription.Cancel()
}

// Next returns the an item from the queue
func (s *psEnvelopeSubscription) Next() (*Envelope, error) {
	next, err := s.subscription.Next()
	if err != nil {
		return nil, err
	}
	return next.(*Envelope), nil
}

// Subscribe to published events with optional filters
func (ps *psEnvelope) Subscribe(filters ...EnvelopeFilter) EnvelopeSubscription {
	// cast filters
	iFilters := make([]pubsub.Filter, len(filters))
	for i, filter := range filters {
		filter := filter
		iFilters[i] = func(v interface{}) bool {
			return filter(v.(*Envelope))
		}
	}
	// create a new subscription
	sub := &psEnvelopeSubscription{
		subscription: ps.pubsub.Subscribe(iFilters...),
	}

	return sub
}

// Publish to all subscribers
func (ps *psEnvelope) Publish(v *Envelope) {
	ps.pubsub.Publish(v)
}