package hyperspace

import (
	"crypto/rand"
	"database/sql"
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
	"nimona.io/pkg/eventbus"
	"nimona.io/pkg/exchange"
	"nimona.io/pkg/keychain"
	"nimona.io/pkg/net"
	"nimona.io/pkg/object"
	"nimona.io/pkg/peer"
	"nimona.io/pkg/sqlobjectstore"
)

func TestDiscoverer_TwoPeersCanFindEachOther(t *testing.T) {
	_, k0, kc0, eb0, n0, x0, disc0, ctx0 := newPeer(t, "peer0")

	d0, err := New(ctx0, disc0, kc0, eb0, x0, nil)
	assert.NoError(t, err)

	err = disc0.AddDiscoverer(d0)
	assert.NoError(t, err)

	ba := []*peer.Peer{
		{
			Addresses: n0.Addresses(),
			Owners:    kc0.ListPublicKeys(keychain.PeerKey),
		},
	}

	time.Sleep(time.Second)

	_, k1, kc1, eb1, n1, x1, disc1, ctx1 := newPeer(t, "peer1")

	d1, err := New(ctx1, disc1, kc1, eb1, x1, ba)
	assert.NoError(t, err)
	err = disc1.AddDiscoverer(d1)
	assert.NoError(t, err)

	time.Sleep(time.Second)

	ctx := context.New(
		context.WithCorrelationID("req1"),
		context.WithTimeout(time.Second),
	)

	peersChan, err := d1.Lookup(ctx, peer.LookupByOwner(k0.PublicKey()))

	peers := gatherPeers(peersChan)
	require.NoError(t, err)
	require.Len(t, peers, 1)
	require.Equal(t, n0.Addresses(), peers[0].Addresses)

	ctxR2 := context.New(
		context.WithCorrelationID("req2"),
		context.WithTimeout(time.Second),
	)
	peersChan, err = d0.Lookup(ctxR2, peer.LookupByOwner(k1.PublicKey()))
	peers = gatherPeers(peersChan)
	require.NoError(t, err)
	require.Len(t, peers, 1)
	require.Equal(t, n1.Addresses(), peers[0].Addresses)
}

func TestDiscoverer_TwoPeersAndOneBootstrapCanFindEachOther(t *testing.T) {
	_, k0, kc0, eb0, n0, x0, disc0, ctx0 := newPeer(t, "peer0")

	// bootstrap node
	d0, err := New(ctx0, disc0, kc0, eb0, x0, nil)
	assert.NoError(t, err)
	err = disc0.AddDiscoverer(d0)
	assert.NoError(t, err)

	time.Sleep(time.Second)

	_, k1, kc1, eb1, n1, x1, disc1, ctx1 := newPeer(t, "peer1")
	_, k2, kc2, eb2, n2, x2, disc2, ctx2 := newPeer(t, "peer2")

	// bootstrap address
	ba := []*peer.Peer{
		{
			Addresses: n0.Addresses(),
			Owners:    kc0.ListPublicKeys(keychain.PeerKey),
		},
	}

	// node 1
	d1, err := New(ctx1, disc1, kc1, eb1, x1, ba)
	assert.NoError(t, err)
	err = disc1.AddDiscoverer(d1)
	assert.NoError(t, err)

	time.Sleep(time.Second)

	// node 2
	d2, err := New(ctx2, disc2, kc2, eb2, x2, ba)
	assert.NoError(t, err)
	err = disc2.AddDiscoverer(d2)
	assert.NoError(t, err)

	time.Sleep(time.Second)

	// find bootstrap from node1
	ctx := context.New(
		context.WithCorrelationID("req1"),
		context.WithTimeout(time.Second*2),
	)
	peersChan, err := d1.Lookup(ctx, peer.LookupByOwner(k0.PublicKey()))
	peers := gatherPeers(peersChan)
	require.NoError(t, err)
	require.Len(t, peers, 1)
	require.Equal(t, n0.Addresses(), peers[0].Addresses)

	// find node 1 from node 2
	ctx = context.New(
		context.WithCorrelationID("req2"),
		context.WithTimeout(time.Second*2),
	)
	peersChan, err = d2.Lookup(ctx, peer.LookupByOwner(k1.PublicKey()))
	peers = gatherPeers(peersChan)
	require.NoError(t, err)
	require.Len(t, peers, 1)
	require.Equal(t, n1.Addresses(), peers[0].Addresses)

	// find node 2 from node 1
	ctx = context.New(
		context.WithCorrelationID("req3"),
		context.WithTimeout(time.Second*2),
	)

	peersChan, err = d1.Lookup(ctx, peer.LookupByOwner(k2.PublicKey()))
	peers = gatherPeers(peersChan)
	require.NoError(t, err)
	require.Len(t, peers, 1)
	require.Equal(t, n2.Addresses(), peers[0].Addresses)

	// add extra peer
	_, k3, kc3, eb3, n3, x3, disc3, ctx3 := newPeer(t, "peer3")

	// setup node 3
	d3, err := New(ctx3, disc3, kc3, eb3, x3, ba)
	assert.NoError(t, err)

	err = disc3.AddDiscoverer(d3)
	assert.NoError(t, err)
	assert.NotNil(t, d3)

	time.Sleep(time.Second)

	fmt.Println("peer0", k0)
	fmt.Println("peer1", k1)
	fmt.Println("peer2", k2)
	fmt.Println("peer3", k3)

	fmt.Println("-------------------")
	fmt.Println("-------------------")
	fmt.Println("-------------------")
	fmt.Println("-------------------")

	// allow bootstraping to settle
	time.Sleep(time.Second)

	// find node 3 from node 1 from it's identity
	ctx = context.New(
		context.WithCorrelationID("req4"),
		context.WithTimeout(time.Second*2),
	)
	peersChan, err = d1.Lookup(
		ctx,
		peer.LookupByCertificateSigner(
			kc3.ListPublicKeys(keychain.IdentityKey)[0],
		),
	)
	peers = gatherPeers(peersChan)
	require.NoError(t, err)
	require.Len(t, peers, 1)
	require.Equal(t, n3.Addresses(), peers[0].Addresses)

	// find node 3 from node 2 from it's identity
	ctx = context.New(
		context.WithCorrelationID("req5"),
		context.WithTimeout(time.Second*2),
	)
	peersChan, err = d2.Lookup(
		ctx,
		peer.LookupByCertificateSigner(
			kc3.ListPublicKeys(keychain.IdentityKey)[0],
		),
	)
	peers = gatherPeers(peersChan)
	require.NoError(t, err)
	require.Len(t, peers, 1)
	require.Equal(t, n3.Addresses(), peers[0].Addresses)
}

func TestDiscoverer_TwoPeersAndOneBootstrapCanProvide(t *testing.T) {
	_, k0, kc0, eb0, n0, x0, disc0, ctx0 := newPeer(t, "peer0")
	_, k1, kc1, eb1, _, x1, disc1, ctx1 := newPeer(t, "peer1")
	_, k2, kc2, eb2, _, x2, disc2, ctx2 := newPeer(t, "peer2")

	// make peer 1 a provider
	token := make([]byte, 32)
	rand.Read(token) // nolint: errcheck
	ch := object.HashFromBytes(token)

	// print peer info
	fmt.Println("k0", k0)
	fmt.Println("k1", k1)
	fmt.Println("k2", k2)

	// bootstrap peer
	d0, err := New(ctx0, disc0, kc0, eb0, x0, nil)
	assert.NoError(t, err)
	err = disc0.AddDiscoverer(d0)
	assert.NoError(t, err)

	time.Sleep(time.Second)

	// bootstrap address
	ba := []*peer.Peer{
		{
			Addresses: n0.Addresses(),
			Owners:    kc0.ListPublicKeys(keychain.PeerKey),
		},
	}

	// peer 1
	d1, err := New(ctx1, disc1, kc1, eb1, x1, ba)
	assert.NoError(t, err)
	err = disc1.AddDiscoverer(d1)
	assert.NoError(t, err)

	eb1.Publish(
		eventbus.ObjectPinned{
			Hash: ch,
		},
	)

	time.Sleep(time.Second)

	// peer 2
	d2, err := New(ctx2, disc2, kc2, eb2, x2, ba)
	assert.NoError(t, err)
	err = disc2.AddDiscoverer(d2)
	assert.NoError(t, err)

	time.Sleep(time.Second)

	// find peer 1 from peer 2
	ctx := context.New(
		context.WithCorrelationID("req1"),
		context.WithTimeout(time.Second),
	)
	providersChan, err := d2.Lookup(ctx, peer.LookupByContentHash(ch))
	providers := gatherPeers(providersChan)
	require.NoError(t, err)
	require.Len(t, providers, 1)
	require.Equal(t, k1.PublicKey(), providers[0].PublicKey())

	// find peer 1 from bootstrap
	ctx = context.New(
		context.WithCorrelationID("req2"),
		context.WithTimeout(time.Second*2),
	)
	providersChan, err = d0.Lookup(ctx, peer.LookupByContentHash(ch))
	providers = gatherPeers(providersChan)
	require.NoError(t, err)
	require.Len(t, providers, 1)
	require.Equal(t, k1.PublicKey(), providers[0].PublicKey())
}

// nolint: gocritic
func newPeer(
	t *testing.T,
	name string,
) (
	crypto.PrivateKey,
	crypto.PrivateKey,
	keychain.Keychain,
	eventbus.Eventbus,
	net.Network,
	exchange.Exchange,
	discovery.PeerStorer,
	context.Context,
) {
	ctx := context.New(context.WithCorrelationID(name))

	eb := eventbus.New()

	// identity key
	opk, err := crypto.GenerateEd25519PrivateKey()
	assert.NoError(t, err)

	// peer key
	pk, err := crypto.GenerateEd25519PrivateKey()
	assert.NoError(t, err)

	// peer certificate
	c := peer.NewCertificate(
		pk.PublicKey(),
		opk,
	)

	kc := keychain.New()
	kc.Put(keychain.PrimaryPeerKey, pk)
	kc.Put(keychain.IdentityKey, opk)
	kc.PutCertificate(&c)

	dblite := tempSqlite3(t)
	store, err := sqlobjectstore.New(dblite)
	assert.NoError(t, err)

	disc := discovery.NewPeerStorer(store)

	n := net.New(
		net.WithEventBus(eb),
		net.WithKeychain(kc),
	)

	_, err = n.Listen(context.Background(), "127.0.0.1:0")
	require.NoError(t, err)

	x, err := exchange.New(
		ctx,
		eb,
		kc,
		n,
		store,
		disc,
	)
	assert.NoError(t, err)

	return opk, pk, kc, eb, n, x, disc, ctx
}

func tempSqlite3(t *testing.T) *sql.DB {
	dirPath, err := ioutil.TempDir("", "nimona-store-sql")
	require.NoError(t, err)
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
