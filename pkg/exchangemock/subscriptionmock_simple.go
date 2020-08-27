package exchangemock

import (
	"sync"

	"nimona.io/pkg/errors"
	"nimona.io/pkg/exchange"
)

type (
	MockSubscriptionSimple struct {
		mutex   sync.Mutex
		index   int
		Objects []*exchange.Envelope
		done    chan struct{}
	}
)

// Cancel the subscription
func (s *MockSubscriptionSimple) Cancel() {
	select {
	case s.done <- struct{}{}:
	default:
	}
}

// Next returns the next object
func (s *MockSubscriptionSimple) Next() (*exchange.Envelope, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.index >= len(s.Objects) {
		s.done <- struct{}{}
		return nil, errors.New("done")
	}
	r := s.Objects[s.index]
	s.index++
	return r, nil
}
