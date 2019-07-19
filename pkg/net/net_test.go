package net

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"nimona.io/internal/context"
	"nimona.io/pkg/crypto"
	"nimona.io/pkg/discovery"
	"nimona.io/pkg/peer"
	npeer "nimona.io/pkg/peer"
)

func TestNetDiscoverer(t *testing.T) {
	disc1 := discovery.NewDiscoverer()
	disc2 := discovery.NewDiscoverer()

	_, _, l1 := newPeer(t, "", disc1)
	_, _, l2 := newPeer(t, "", disc2)

	ctx := context.New()

	disc1.Add(l2.GetSignedPeer())
	disc2.Add(l1.GetSignedPeer())

	ps2, err := disc1.FindByFingerprint(ctx, l2.GetPeerKey().Fingerprint())
	p2 := ps2[0]
	assert.NoError(t, err)
	// assert.Equal(t, n2.key.PublicKey, p2.SignerKey)
	assert.Equal(t,
		l2.GetPeerKey().Fingerprint(),
		p2.Fingerprint(),
	)

	ps1, err := disc2.FindByFingerprint(ctx, l1.GetPeerKey().Fingerprint())
	p1 := ps1[0]
	assert.NoError(t, err)
	// assert.Equal(t, n1.key.PublicKey, p1.SignerKey)
	assert.Equal(t,
		l1.GetPeerKey().Fingerprint(),
		p1.Fingerprint(),
	)
}

func TestNetConnectionSuccess(t *testing.T) {

	disc1 := discovery.NewDiscoverer()
	disc2 := discovery.NewDiscoverer()

	ctx := context.New()

	BindLocal = true
	_, n1, l1 := newPeer(t, "", disc1)
	_, n2, l2 := newPeer(t, "", disc2)

	// we need to start listening before we add the peer
	// otherwise the addresses are not populated
	sconn, err := n1.Listen(ctx)
	assert.NoError(t, err)

	disc1.Add(l2.GetSignedPeer())
	disc2.Add(l1.GetSignedPeer())

	peer1Addr := l1.GetAddresses()[0]

	done := make(chan bool)

	go func() {
		cconn, err := n2.Dial(ctx, peer1Addr)
		assert.NoError(t, err)
		err = Write((npeer.PeerRequest{}).ToObject(), cconn)
		assert.NoError(t, err)
		done <- true

	}()

	sc := <-sconn

	err = Write((npeer.PeerRequest{}).ToObject(), sc)
	assert.NoError(t, err)

	<-done
}

func TestNetConnectionFailureMiddleware(t *testing.T) {

	disc1 := discovery.NewDiscoverer()
	disc2 := discovery.NewDiscoverer()

	ctx := context.New()

	BindLocal = true
	_, n1, l1 := newPeer(t, "", disc1)
	_, n2, l2 := newPeer(t, "", disc2)

	// we need to start listening before we add the peer
	// otherwise the addresses are not populated
	fm := fakeMid{}

	sconn, err := n1.Listen(ctx)
	n1.AddMiddleware(fm.Handle())
	assert.NoError(t, err)

	disc1.Add(l2.GetSignedPeer())
	disc2.Add(l1.GetSignedPeer())

	peer1Addr := l1.GetAddresses()[0]

	done := make(chan bool)

	go func() {
		cconn, err := n2.Dial(ctx, peer1Addr)
		assert.Error(t, err)
		assert.Nil(t, cconn)
		done <- true
	}()

	<-done
	assert.Len(t, sconn, 0)
}

func newPeer(t *testing.T, relayAddress string, discover discovery.Discoverer) (
	*crypto.PrivateKey, *network, *peer.LocalPeer) {
	pk, err := crypto.GenerateKey()
	assert.NoError(t, err)

	localInfo, err := peer.NewLocalPeer("", pk) // nolint: ineffassign
	n, err := New(discover, localInfo)
	assert.NoError(t, err)

	tcpTr := NewTCPTransport(
		localInfo,
		"0.0.0.0:0",
	)
	n.AddTransport("tcps", tcpTr)

	return pk, n.(*network), localInfo
}

type fakeMid struct {
}

func (fm *fakeMid) Handle() MiddlewareHandler {
	return func(ctx context.Context, conn *Connection) (*Connection, error) {
		return conn, errors.New("what?")
	}
}
