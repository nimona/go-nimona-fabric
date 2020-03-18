package net

import (
	"database/sql"
	"errors"
	"fmt"
	"io/ioutil"
	"path"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	_ "github.com/mattn/go-sqlite3"

	"nimona.io/pkg/context"
	"nimona.io/pkg/crypto"
	"nimona.io/pkg/discovery"
	"nimona.io/pkg/object"
	"nimona.io/pkg/peer"
	"nimona.io/pkg/sqlobjectstore"
)

func TestNetDiscoverer(t *testing.T) {
	dblite1 := tempSqlite3(t)
	store1, err := sqlobjectstore.New(dblite1)
	assert.NoError(t, err)

	dblite2 := tempSqlite3(t)
	store2, err := sqlobjectstore.New(dblite2)
	assert.NoError(t, err)

	disc1 := discovery.NewPeerStorer(store1)
	disc2 := discovery.NewPeerStorer(store2)

	_, _, l1 := newPeer(t, "", disc1)
	_, _, l2 := newPeer(t, "", disc2)

	ctx := context.New()

	disc1.Add(l2.GetSignedPeer(), true)
	disc2.Add(l1.GetSignedPeer(), true)

	ps2, err := disc1.Lookup(ctx, peer.LookupByOwner(l2.GetPeerPublicKey()))
	p2 := gatherPeers(ps2)[0]
	assert.NoError(t, err)
	// assert.Equal(t, n2.key.PublicKey, p2.SignerKey)
	assert.Equal(t,
		l2.GetPeerPublicKey(),
		p2.PublicKey(),
	)

	ps1, err := disc2.Lookup(ctx, peer.LookupByOwner(l1.GetPeerPublicKey()))
	p1 := gatherPeers(ps1)[0]
	assert.NoError(t, err)
	// assert.Equal(t, n1.key.PublicKey, p1.SignerKey)
	assert.Equal(t,
		l1.GetPeerPublicKey(),
		p1.PublicKey(),
	)
}

func TestNetConnectionSuccess(t *testing.T) {
	dblite1 := tempSqlite3(t)
	store1, err := sqlobjectstore.New(dblite1)
	assert.NoError(t, err)

	dblite2 := tempSqlite3(t)
	store2, err := sqlobjectstore.New(dblite2)
	assert.NoError(t, err)

	disc1 := discovery.NewPeerStorer(store1)
	disc2 := discovery.NewPeerStorer(store2)

	ctx := context.New()

	BindLocal = true
	_, n1, l1 := newPeer(t, "", disc1)
	_, n2, l2 := newPeer(t, "", disc2)

	// we need to start listening before we add the peer
	// otherwise the addresses are not populated
	sconn, err := n1.Listen(ctx)
	assert.NoError(t, err)

	disc1.Add(l2.GetSignedPeer(), true)
	disc2.Add(l1.GetSignedPeer(), true)

	done := make(chan bool)

	go func() {
		cconn, err := n2.Dial(ctx, l1.GetSignedPeer())
		assert.NoError(t, err)
		o := object.FromMap(map[string]interface{}{ // nolint: errcheck
			"foo:s": "bar",
		})
		err = Write(o, cconn)
		assert.NoError(t, err)
		done <- true
	}()

	sc := <-sconn

	o := object.FromMap(map[string]interface{}{ // nolint: errcheck
		"data:o": map[string]interface{}{
			"foo:s": "bar",
		},
	})
	err = Write(o, sc)
	assert.NoError(t, err)

	<-done
}

func TestNetDialBackoff(t *testing.T) {
	dblite1 := tempSqlite3(t)
	store1, err := sqlobjectstore.New(dblite1)
	assert.NoError(t, err)

	disc1 := discovery.NewPeerStorer(store1)

	ctx := context.New()

	p := &peer.Peer{
		Owners:    []crypto.PublicKey{"foo"},
		Addresses: []string{"tcps:240.0.0.1:1000"},
	}

	// attempt 1, failed
	_, n1, _ := newPeer(t, "", disc1)
	_, err = n1.Dial(ctx, p)
	assert.Equal(t, ErrAllAddressesFailed, err)

	// attempt 2, blacklisted
	_, err = n1.Dial(ctx, p)
	assert.Equal(t, ErrAllAddressesBlacklisted, err)

	// wait for backoff to expire
	time.Sleep(time.Second * 2)

	// attempt 3, failed
	_, err = n1.Dial(ctx, p)
	assert.Equal(t, ErrAllAddressesFailed, err)

	// attempt 4, blacklisted
	_, err = n1.Dial(ctx, p)
	assert.Equal(t, ErrAllAddressesBlacklisted, err)

	// wait, but backoff should not have expired
	time.Sleep(time.Second * 2)

	// attempt 5, blacklisted
	_, err = n1.Dial(ctx, p)
	assert.Equal(t, ErrAllAddressesBlacklisted, err)

	// wait for backoff to expire
	time.Sleep(time.Second * 2)

	// attempt 6, failed
	_, err = n1.Dial(ctx, p)
	assert.Equal(t, ErrAllAddressesFailed, err)
}

func TestNetConnectionFailureMiddleware(t *testing.T) {
	dblite1 := tempSqlite3(t)
	store1, err := sqlobjectstore.New(dblite1)
	assert.NoError(t, err)

	dblite2 := tempSqlite3(t)
	store2, err := sqlobjectstore.New(dblite2)
	assert.NoError(t, err)

	disc1 := discovery.NewPeerStorer(store1)
	disc2 := discovery.NewPeerStorer(store2)

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

	disc1.Add(l2.GetSignedPeer(), true)
	disc2.Add(l1.GetSignedPeer(), true)

	done := make(chan bool)

	go func() {
		cconn, err := n2.Dial(ctx, l1.GetSignedPeer())
		assert.Error(t, err)
		assert.Nil(t, cconn)
		done <- true
	}()

	<-done
	assert.Len(t, sconn, 0)
}

func newPeer(
	t *testing.T,
	relayAddress string,
	discover discovery.Discoverer,
) (
	crypto.PrivateKey,
	*network,
	*peer.LocalPeer,
) {
	pk, err := crypto.GenerateEd25519PrivateKey()
	assert.NoError(t, err)

	localInfo, _ := peer.NewLocalPeer("", pk) // nolint: ineffassign, errcheck
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

func tempSqlite3(t *testing.T) *sql.DB {
	dirPath, err := ioutil.TempDir("", "nimona-new")
	require.NoError(t, err)
	fmt.Println(path.Join(dirPath, "sqlite3.db"))
	db, err := sql.Open("sqlite3", path.Join(dirPath, "sqlite3.db"))
	require.NoError(t, err)
	return db
}

func gatherPeers(p <-chan *peer.Peer) []*peer.Peer {
	ps := []*peer.Peer{}
	for p := range p {
		p := p
		ps = append(ps, p)
	}
	return peer.Unique(ps)
}
