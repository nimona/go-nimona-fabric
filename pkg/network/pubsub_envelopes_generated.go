// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/geoah/genny

package network

import (
	"nimona.io/internal/pubsub"
)

type (
	// EnvelopePubSub -
	EnvelopePubSub interface {
		Publish(*Envelope)
		Subscribe(...EnvelopeFilter) EnvelopeSubscription
	}
	EnvelopeFilter func(*Envelope) bool
	// EnvelopeSubscription is returned for every subscription
	EnvelopeSubscription interface {
		Channel() <-chan *Envelope
		Next() (*Envelope, error)
		Cancel()
	}
	envelopeSubscription struct {
		subscription pubsub.Subscription
	}
	envelopePubSub struct {
		pubsub pubsub.PubSub
	}
)

// NewEnvelope constructs and returns a new Envelope
func NewEnvelopePubSub() EnvelopePubSub {
	return &envelopePubSub{
		pubsub: pubsub.New(),
	}
}

// Cancel the subscription
func (s *envelopeSubscription) Cancel() {
	s.subscription.Cancel()
}

// Channel returns a channel that will be returning the items from the queue
func (s *envelopeSubscription) Channel() <-chan *Envelope {
	c := s.subscription.Channel()
	r := make(chan *Envelope)
	go func() {
		for {
			v, ok := <-c
			if !ok {
				close(r)
				return
			}
			r <- v.(*Envelope)
		}
	}()
	return r
}

// Next returns the next item from the queue
func (s *envelopeSubscription) Next() (r *Envelope, err error) {
	next, err := s.subscription.Next()
	if err != nil {
		return
	}
	return next.(*Envelope), nil
}

// Subscribe to published events with optional filters
func (ps *envelopePubSub) Subscribe(filters ...EnvelopeFilter) EnvelopeSubscription {
	// cast filters
	iFilters := make([]pubsub.Filter, len(filters))
	for i, filter := range filters {
		filter := filter
		iFilters[i] = func(v interface{}) bool {
			return filter(v.(*Envelope))
		}
	}
	// create a new subscription
	sub := &envelopeSubscription{
		subscription: ps.pubsub.Subscribe(iFilters...),
	}

	return sub
}

// Publish to all subscribers
func (ps *envelopePubSub) Publish(v *Envelope) {
	ps.pubsub.Publish(v)
}
