// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package peer

import (
	crypto "nimona.io/pkg/crypto"
	"sync"
)

type (
	// CryptoPublicKeySyncList -
	CryptoPublicKeySyncList struct {
		m sync.Map
	}
)

// NewCryptoPublicKeyValueTypeSyncMap constructs a new SyncMap
func NewCryptoPublicKeyValueTypeSyncMap() *CryptoPublicKeySyncList {
	return &CryptoPublicKeySyncList{}
}

// Put -
func (m *CryptoPublicKeySyncList) Put(k crypto.PublicKey) {
	m.m.Store(k, true)
}

// Exists -
func (m *CryptoPublicKeySyncList) Exists(k crypto.PublicKey) bool {
	_, ok := m.m.Load(k)
	if !ok {
		return false
	}

	return true
}

// Delete -
func (m *CryptoPublicKeySyncList) Delete(k crypto.PublicKey) {
	m.m.Delete(k)
}

// Range -
func (m *CryptoPublicKeySyncList) Range(i func(k crypto.PublicKey) bool) {
	m.m.Range(func(k, v interface{}) bool {
		return i(k.(crypto.PublicKey))
	})
}