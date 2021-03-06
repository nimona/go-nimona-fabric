// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/geoah/genny

package network

import (
	"sync"
)

type (
	// StringSyncList -
	StringSyncList struct {
		m sync.Map
	}
)

// Put -
func (m *StringSyncList) Put(k string) {
	m.m.Store(k, true)
}

// Exists -
func (m *StringSyncList) Exists(k string) bool {
	_, ok := m.m.Load(k)
	return ok
}

// Delete -
func (m *StringSyncList) Delete(k string) {
	m.m.Delete(k)
}

// Range -
func (m *StringSyncList) Range(i func(k string) bool) {
	m.m.Range(func(k, v interface{}) bool {
		return i(k.(string))
	})
}

// List -
func (m *StringSyncList) List() []string {
	r := []string{}
	m.m.Range(func(k, v interface{}) bool {
		r = append(r, k.(string))
		return true
	})
	return r
}
