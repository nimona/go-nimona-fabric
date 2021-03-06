package feedmanager_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"nimona.io/internal/fixtures"
	"nimona.io/internal/rand"
	"nimona.io/pkg/config"
	"nimona.io/pkg/context"
	"nimona.io/pkg/crypto"
	"nimona.io/pkg/daemon"
	"nimona.io/pkg/hyperspace/resolver"
	"nimona.io/pkg/object"
	"nimona.io/pkg/peer"
	"nimona.io/pkg/testutils"
)

func TestManager_Integration(t *testing.T) {
	// construct local bootstrap peer
	bootstrapConnectionInfo := testutils.NewTestBootstrapPeer(t)

	time.Sleep(time.Second * 5)

	// construct new identity key
	id, err := crypto.NewEd25519PrivateKey()
	require.NoError(t, err)

	// construct peer 0
	p0 := newDaemon(t, "p0", id, bootstrapConnectionInfo)
	defer p0.Close()

	time.Sleep(time.Second)

	// construct peer 1
	p1 := newDaemon(t, "p1", id, bootstrapConnectionInfo)
	defer p1.Close()

	fmt.Println("p0", p0.LocalPeer().GetPeerKey().PublicKey())
	fmt.Println("p1", p1.LocalPeer().GetPeerKey().PublicKey())

	// put a new stream on p0
	o0 := object.MustMarshal(
		&fixtures.TestStream{
			Metadata: object.Metadata{
				Owner: id.PublicKey(),
			},
			Nonce: "foo",
		},
	)
	_, err = p0.ObjectManager().Put(context.TODO(), o0)
	require.NoError(t, err)

	// wait until resolver sees provider
	found := false
	for i := 0; i < 10; i++ {
		r, err := p1.Resolver().Lookup(
			context.New(
				context.WithTimeout(time.Second),
			),
			resolver.LookupByHash(o0.Hash()),
		)
		if err != nil {
			continue
		}
		if len(r) > 0 {
			found = true
			break
		}
		time.Sleep(time.Second * 2)
	}
	require.True(t, found)

	time.Sleep(time.Second * 2)

	// wait a bit, and check stream on p1
	g0, err := p1.ObjectStore().Get(o0.Hash())
	require.NoError(t, err)
	require.NotNil(t, g0)
}

func newDaemon(
	t *testing.T,
	name string,
	id crypto.PrivateKey,
	bootstrapConnectionInfo *peer.ConnectionInfo,
) daemon.Daemon {
	d, err := daemon.New(
		context.Background(),
		daemon.WithConfigOptions(
			config.WithDefaultPath(
				t.TempDir(),
			),
			config.WithDefaultListenOnLocalIPs(),
			config.WithDefaultListenOnPrivateIPs(),
			config.WithDefaultBootstraps([]peer.Shorthand{
				peer.Shorthand(
					fmt.Sprintf(
						"%s@%s",
						bootstrapConnectionInfo.PublicKey.String(),
						bootstrapConnectionInfo.Addresses[0],
					),
				),
			},
			),
		),
	)
	require.NoError(t, err)
	time.Sleep(time.Second)

	peerKey := d.LocalPeer().GetPeerKey()
	csr := &object.CertificateRequest{
		Metadata: object.Metadata{
			Owner: peerKey.PublicKey(),
		},
		Nonce:      rand.String(8),
		VendorName: "foo",
	}
	csrObj := object.MustMarshal(csr)
	csr.Metadata.Signature, err = object.NewSignature(peerKey, csrObj)
	require.NoError(t, err)

	csrRes, err := object.NewCertificate(id, *csr, true, "bar")
	require.NoError(t, err)

	err = d.FeedManager().RegisterFeed(fixtures.TestStreamType)
	require.NoError(t, err)

	d.LocalPeer().SetPeerCertificate(csrRes)

	return d
}
