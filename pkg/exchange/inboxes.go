// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package exchange

import "sync"

type (
	inboxes string // nolint
	// InboxesMap -
	InboxesMap struct {
		m sync.Map
	}
)

// NewInboxesMap constructs a new SyncMap
func NewInboxesMap() *InboxesMap {
	return &InboxesMap{}
}

// GetOrPut -
func (m *InboxesMap) GetOrPut(k string, v *inbox) (*inbox, bool) {
	nv, ok := m.m.LoadOrStore(k, v)
	return nv.(*inbox), ok
}

// Put -
func (m *InboxesMap) Put(k string, v *inbox) {
	m.m.Store(k, v)
}

// Get -
func (m *InboxesMap) Get(k string) (*inbox, bool) {
	i, ok := m.m.Load(k)
	if !ok {
		return nil, false
	}

	v, ok := i.(*inbox)
	if !ok {
		return nil, false
	}

	return v, true
}

// Delete -
func (m *InboxesMap) Delete(k string) {
	m.m.Delete(k)
}

// Range -
func (m *InboxesMap) Range(i func(k string, v *inbox) bool) {
	m.m.Range(func(k, v interface{}) bool {
		return i(k.(string), v.(*inbox))
	})
}
