package main

import (
	"context"
	"crypto/rand"
	"fmt"

	"github.com/tylertreat/bench"

	"github.com/nimona/go-nimona/mesh"
	"github.com/nimona/go-nimona/wire"
)

// WireRequesterFactory implements RequesterFactory for our wire
type WireRequesterFactory struct {
	wire      wire.Wire
	recipient mesh.PeerInfo
	bytes     int
}

// GetRequester returns a new Requester, called for each Benchmark connection.
func (w *WireRequesterFactory) GetRequester(uint64) bench.Requester {
	return &wireRequester{
		wire:      w.wire,
		recipient: w.recipient,
		bytes:     w.bytes,
	}
}

// wireRequester implements Requester by making sending a message to a peer
type wireRequester struct {
	wire      wire.Wire
	recipient mesh.PeerInfo
	bytes     int
	payload   []byte
}

// Setup prepares the Requester for benchmarking.
func (w *wireRequester) Setup() error {
	w.payload = make([]byte, w.bytes)
	if _, err := rand.Read(w.payload); err != nil {
		return err
	}
	return nil
}

// Request performs a synchronous request to the system under test.
func (w *wireRequester) Request() error {
	ctx := context.Background()
	recipient := w.recipient.ID
	err := w.wire.Send(ctx, "foo", "bar", w.payload, []string{recipient})
	if err != nil {
		fmt.Println("could not send message, error:", err)
	}
	return err
}

// Teardown is called upon benchmark completion.
func (w *wireRequester) Teardown() error {
	return nil
}
