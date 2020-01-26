package sqlobjectstore

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"path"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	_ "github.com/mattn/go-sqlite3"

	"nimona.io/internal/fixtures"
	"nimona.io/pkg/crypto"
	"nimona.io/pkg/errors"
	"nimona.io/pkg/hash"
	"nimona.io/pkg/object"
	"nimona.io/pkg/stream"
)

func tempSqlite3(t *testing.T) *sql.DB {
	dirPath, err := ioutil.TempDir("", "nimona-store-sql")
	require.NoError(t, err)
	db, err := sql.Open("sqlite3", path.Join(dirPath, "sqlite3.db"))
	require.NoError(t, err)
	return db
}

func TestNewDatabase(t *testing.T) {
	dblite := tempSqlite3(t)
	store, err := New(dblite)
	require.NoError(t, err)
	require.NotNil(t, store)

	err = store.Close()
	require.NoError(t, err)
}

func TestStoreRetrieveUpdate(t *testing.T) {
	dblite := tempSqlite3(t)
	store, err := New(dblite)
	require.NoError(t, err)
	require.NotNil(t, store)

	p := fixtures.TestStream{
		Nonce: "asdf",
	}
	c := fixtures.TestSubscribed{
		Stream: hash.New(p.ToObject()),
	}
	obj := c.ToObject()
	obj.SetType("foo")
	obj.Set("key:s", "value")

	err = store.Put(
		obj,
		WithTTL(0),
	)
	require.NoError(t, err)

	err = store.Put(
		obj,
		WithTTL(10),
	)

	require.NoError(t, err)
	retrievedObj, err := store.Get(hash.New(obj))
	require.NoError(t, err)

	val := retrievedObj.Get("key:s")
	require.NotNil(t, val)
	assert.Equal(t, "value", val.(string))

	stHash := stream.GetStream(obj)
	require.NotEmpty(t, stHash)

	err = store.UpdateTTL(hash.New(obj), 10)
	require.NoError(t, err)

	hashList, err := store.GetRelations(hash.New(p.ToObject()))
	require.NoError(t, err)
	assert.NotEmpty(t, hashList)

	err = store.Remove(hash.New(p.ToObject()))
	require.NoError(t, err)

	retrievedObj2, err := store.Get(hash.New(p.ToObject()))
	require.True(t, errors.CausedBy(err, ErrNotFound))
	require.Nil(t, retrievedObj2)

	err = store.Close()
	require.NoError(t, err)
}

func TestSubscribe(t *testing.T) {
	// create db
	dblite := tempSqlite3(t)
	store, err := New(dblite)
	require.NoError(t, err)
	require.NotNil(t, store)

	// setup data
	p := fixtures.TestStream{
		Nonce: "asdf",
	}
	streamHash := hash.New(p.ToObject())
	c := fixtures.TestSubscribed{
		Stream: streamHash,
	}
	obj := c.ToObject()
	obj.SetType("foo")
	obj.Set("key:s", "value")

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		// subscribe
		subscription := store.Subscribe(
			FilterByStreamHash(streamHash),
		)

		go func() {
			hs, err := subscription.Next()
			require.NoError(t, err)
			require.NotEmpty(t, hs)
			wg.Done()
		}()
	}

	// store data
	err = store.Put(
		obj,
		WithTTL(10),
	)
	require.NoError(t, err)

	done := make(chan bool)
	go func() {
		wg.Wait()
		done <- true
	}()

	select {
	case <-done:
		break
	case <-time.After(1 * time.Second):
		t.Fatalf("failed to get update")
	}
}

func TestFilter(t *testing.T) {
	dblite := tempSqlite3(t)
	store, err := New(dblite)
	require.NoError(t, err)
	require.NotNil(t, store)

	k, err := crypto.GenerateEd25519PrivateKey()
	require.NoError(t, err)

	p := fixtures.TestStream{
		Nonce: "asdf",
	}

	s, err := crypto.NewSignature(k, p.ToObject())
	require.NoError(t, err)

	p.Signature = s

	err = store.Put(p.ToObject(), WithTTL(0))
	require.NoError(t, err)

	ph := hash.New(p.ToObject())

	c := fixtures.TestSubscribed{
		Stream: ph,
	}

	hashes := []object.Hash{}
	for i := 0; i < 5; i++ {
		obj := c.ToObject()
		obj.Set("key:s", fmt.Sprintf("value_%d", i))
		if i%2 == 0 {
			obj.Set("@identity:s", s.Signer.String())
		}
		err = store.Put(obj, WithTTL(0))
		require.NoError(t, err)
		hashes = append(hashes, hash.New(obj))
	}

	objects, err := store.Filter(
		FilterByHash(hashes[0]),
		FilterByHash(hashes[1]),
		FilterByHash(hashes[2]),
		FilterByHash(hashes[3]),
		FilterByHash(hashes[4]),
	)
	require.NoError(t, err)
	require.Len(t, objects, len(hashes))

	objects, err = store.Filter(
		FilterBySigner(k.PublicKey()),
	)
	require.NoError(t, err)
	require.Len(t, objects, 1)

	objects, err = store.Filter(
		FilterByIdentity(k.PublicKey()),
	)
	require.NoError(t, err)
	require.Len(t, objects, 3)

	objects, err = store.Filter(
		FilterBySigner(crypto.PublicKey("foo")),
	)
	require.NoError(t, err)
	require.Len(t, objects, 0)

	objects, err = store.Filter(
		FilterByObjectType(c.GetType()),
	)
	require.NoError(t, err)
	require.Len(t, objects, len(hashes))

	objects, err = store.Filter(
		FilterByStreamHash(ph),
	)
	require.NoError(t, err)
	require.Len(t, objects, len(hashes)+1)

	objects, err = store.Filter(
		FilterByHash(hashes[0]),
		FilterByObjectType(c.GetType()),
		FilterByStreamHash(ph),
	)
	require.NoError(t, err)
	require.Len(t, objects, 1)

	err = store.Close()
	require.NoError(t, err)
}