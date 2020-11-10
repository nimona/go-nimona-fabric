// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/geoah/genny

package localpeer

import (
	"sync"

	"nimona.io/pkg/peer"
)

type (
	// PeerPeerSyncList -
	PeerPeerSyncList struct {
		m sync.Map
	}
)

// Put -
func (m *PeerPeerSyncList) Put(k *peer.ConnectionInfo) {
	m.m.Store(k, true)
}

// Exists -
func (m *PeerPeerSyncList) Exists(k *peer.ConnectionInfo) bool {
	_, ok := m.m.Load(k)
	return ok
}

// Delete -
func (m *PeerPeerSyncList) Delete(k *peer.ConnectionInfo) {
	m.m.Delete(k)
}

// Range -
func (m *PeerPeerSyncList) Range(i func(k *peer.ConnectionInfo) bool) {
	m.m.Range(func(k, v interface{}) bool {
		return i(k.(*peer.ConnectionInfo))
	})
}

// List -
func (m *PeerPeerSyncList) List() []*peer.ConnectionInfo {
	r := []*peer.ConnectionInfo{}
	m.m.Range(func(k, v interface{}) bool {
		r = append(r, k.(*peer.ConnectionInfo))
		return true
	})
	return r
}
