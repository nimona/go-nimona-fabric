package networkmock

import (
	"sync"

	"nimona.io/pkg/errors"
	"nimona.io/pkg/network"
)

type (
	MockSubscriptionSimple struct {
		mutex   sync.Mutex
		index   int
		Objects []*network.Envelope
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
func (s *MockSubscriptionSimple) Next() (*network.Envelope, error) {
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