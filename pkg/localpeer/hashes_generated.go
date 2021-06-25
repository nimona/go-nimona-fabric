// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/geoah/genny

package localpeer

import (
	"sync"

	"nimona.io/pkg/chore"
)

type (
	// ChoreHashSyncList -
	ChoreHashSyncList struct {
		m sync.Map
	}
)

// Put -
func (m *ChoreHashSyncList) Put(k chore.Hash) {
	m.m.Store(k, true)
}

// Exists -
func (m *ChoreHashSyncList) Exists(k chore.Hash) bool {
	_, ok := m.m.Load(k)
	return ok
}

// Delete -
func (m *ChoreHashSyncList) Delete(k chore.Hash) {
	m.m.Delete(k)
}

// Range -
func (m *ChoreHashSyncList) Range(i func(k chore.Hash) bool) {
	m.m.Range(func(k, v interface{}) bool {
		return i(k.(chore.Hash))
	})
}

// List -
func (m *ChoreHashSyncList) List() []chore.Hash {
	r := []chore.Hash{}
	m.m.Range(func(k, v interface{}) bool {
		r = append(r, k.(chore.Hash))
		return true
	})
	return r
}
