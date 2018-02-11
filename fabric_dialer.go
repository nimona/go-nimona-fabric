package fabric

import (
	"context"
	"errors"

	"go.uber.org/zap"
)

var (
	// ErrCouldNotDial when no transports are available or internal error occured
	ErrCouldNotDial = errors.New("Could not dial")
)

// RequestIDKey for context
type RequestIDKey struct{}

// DialContext will attempt to connect to the given address and go through the
// various middlware that it needs until the connection is fully established
func (f *Fabric) DialContext(ctx context.Context, as string) (context.Context, Conn, error) {
	ctx = context.WithValue(ctx, RequestIDKey{}, generateReqID())
	lgr := Logger(ctx)
	lgr.Info("Dialing", zap.String("address", as))

	// TODO validate the address
	addr := NewAddress(as)

	// figure out if the addr can be dialed and connect to the target
	c, err := f.dialTransport(ctx, addr)
	if err != nil {
		return nil, nil, err
	}

	// pop first item which should be the transport
	c.GetAddress().Pop()

	return f.Negotiate(ctx, c)
}

func (f *Fabric) dialTransport(ctx context.Context, addr Address) (Conn, error) {
	// get transport
	tr, err := f.getTransport(addr)
	if err != nil {
		return nil, ErrNoTransport
	}

	// dial
	tcon, err := tr.DialContext(ctx, addr)
	if err != nil {
		return nil, ErrCouldNotDial
	}

	// create a new Conn that will be used to hold underlaying connections
	// from transports, protocol, as well as information about the
	// two parties.
	c := newConnWrapper(tcon, &addr)

	return c, nil
}

func (f *Fabric) getTransport(addr Address) (Transport, error) {
	// find transport we can dial
	// TODO figure out priorities, eg yamux should be more important than tcp
	for _, tr := range f.transports {
		cd, err := tr.CanDial(addr)
		if err != nil {
			// TODO should we error here?
			return nil, ErrCouldNotDial
		}
		if cd {
			return tr, nil
		}
	}

	return nil, ErrNoTransport
}
