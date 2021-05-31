package peer

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"nimona.io/pkg/crypto"
	"nimona.io/pkg/object"
)

func TestEncoding(t *testing.T) {
	k, err := crypto.NewEd25519PrivateKey(crypto.PeerKey)
	require.NoError(t, err)

	c := &ConnectionInfo{
		Metadata: object.Metadata{
			Owner:    k.PublicKey(),
			Datetime: time.Now().Format(time.RFC3339),
		},
		Version:       1,
		PublicKey:     k.PublicKey(),
		Addresses:     []string{"foo", "bar"},
		ObjectFormats: []string{"foobar"},
	}
	b, err := json.Marshal(object.MustMarshal(c))
	require.NoError(t, err)

	fmt.Println(string(b))

	g := &object.Object{}
	err = json.Unmarshal(b, g)
	require.NoError(t, err)

	fmt.Println(g.Data["addresses"])

	r := &ConnectionInfo{}
	err = r.UnmarshalObject(g)
	require.NoError(t, err)

	require.Equal(t, c, r)
}
