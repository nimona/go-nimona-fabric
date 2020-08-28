// Code generated by MockGen. DO NOT EDIT.
// Source: localpeer.go

// Package localpeermock is a generated GoMock package.
package localpeermock

import (
	gomock "github.com/golang/mock/gomock"
	crypto "nimona.io/pkg/crypto"
	object "nimona.io/pkg/object"
	peer "nimona.io/pkg/peer"
	reflect "reflect"
)

// MockLocalPeer is a mock of LocalPeer interface
type MockLocalPeer struct {
	ctrl     *gomock.Controller
	recorder *MockLocalPeerMockRecorder
}

// MockLocalPeerMockRecorder is the mock recorder for MockLocalPeer
type MockLocalPeerMockRecorder struct {
	mock *MockLocalPeer
}

// NewMockLocalPeer creates a new mock instance
func NewMockLocalPeer(ctrl *gomock.Controller) *MockLocalPeer {
	mock := &MockLocalPeer{ctrl: ctrl}
	mock.recorder = &MockLocalPeerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLocalPeer) EXPECT() *MockLocalPeerMockRecorder {
	return m.recorder
}

// GetPrimaryPeerKey mocks base method
func (m *MockLocalPeer) GetPrimaryPeerKey() crypto.PrivateKey {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPrimaryPeerKey")
	ret0, _ := ret[0].(crypto.PrivateKey)
	return ret0
}

// GetPrimaryPeerKey indicates an expected call of GetPrimaryPeerKey
func (mr *MockLocalPeerMockRecorder) GetPrimaryPeerKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPrimaryPeerKey", reflect.TypeOf((*MockLocalPeer)(nil).GetPrimaryPeerKey))
}

// PutPrimaryPeerKey mocks base method
func (m *MockLocalPeer) PutPrimaryPeerKey(arg0 crypto.PrivateKey) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PutPrimaryPeerKey", arg0)
}

// PutPrimaryPeerKey indicates an expected call of PutPrimaryPeerKey
func (mr *MockLocalPeerMockRecorder) PutPrimaryPeerKey(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutPrimaryPeerKey", reflect.TypeOf((*MockLocalPeer)(nil).PutPrimaryPeerKey), arg0)
}

// GetPrimaryIdentityKey mocks base method
func (m *MockLocalPeer) GetPrimaryIdentityKey() crypto.PrivateKey {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPrimaryIdentityKey")
	ret0, _ := ret[0].(crypto.PrivateKey)
	return ret0
}

// GetPrimaryIdentityKey indicates an expected call of GetPrimaryIdentityKey
func (mr *MockLocalPeerMockRecorder) GetPrimaryIdentityKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPrimaryIdentityKey", reflect.TypeOf((*MockLocalPeer)(nil).GetPrimaryIdentityKey))
}

// PutPrimaryIdentityKey mocks base method
func (m *MockLocalPeer) PutPrimaryIdentityKey(arg0 crypto.PrivateKey) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PutPrimaryIdentityKey", arg0)
}

// PutPrimaryIdentityKey indicates an expected call of PutPrimaryIdentityKey
func (mr *MockLocalPeerMockRecorder) PutPrimaryIdentityKey(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutPrimaryIdentityKey", reflect.TypeOf((*MockLocalPeer)(nil).PutPrimaryIdentityKey), arg0)
}

// GetCertificates mocks base method
func (m *MockLocalPeer) GetCertificates(arg0 crypto.PublicKey) []*peer.Certificate {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCertificates", arg0)
	ret0, _ := ret[0].([]*peer.Certificate)
	return ret0
}

// GetCertificates indicates an expected call of GetCertificates
func (mr *MockLocalPeerMockRecorder) GetCertificates(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCertificates", reflect.TypeOf((*MockLocalPeer)(nil).GetCertificates), arg0)
}

// PutCertificate mocks base method
func (m *MockLocalPeer) PutCertificate(arg0 *peer.Certificate) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PutCertificate", arg0)
}

// PutCertificate indicates an expected call of PutCertificate
func (mr *MockLocalPeerMockRecorder) PutCertificate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutCertificate", reflect.TypeOf((*MockLocalPeer)(nil).PutCertificate), arg0)
}

// GetContentHashes mocks base method
func (m *MockLocalPeer) GetContentHashes() []object.Hash {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContentHashes")
	ret0, _ := ret[0].([]object.Hash)
	return ret0
}

// GetContentHashes indicates an expected call of GetContentHashes
func (mr *MockLocalPeerMockRecorder) GetContentHashes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContentHashes", reflect.TypeOf((*MockLocalPeer)(nil).GetContentHashes))
}

// PutContentHashes mocks base method
func (m *MockLocalPeer) PutContentHashes(arg0 ...object.Hash) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "PutContentHashes", varargs...)
}

// PutContentHashes indicates an expected call of PutContentHashes
func (mr *MockLocalPeerMockRecorder) PutContentHashes(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutContentHashes", reflect.TypeOf((*MockLocalPeer)(nil).PutContentHashes), arg0...)
}

// GetRelays mocks base method
func (m *MockLocalPeer) GetRelays() []*peer.Peer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRelays")
	ret0, _ := ret[0].([]*peer.Peer)
	return ret0
}

// GetRelays indicates an expected call of GetRelays
func (mr *MockLocalPeerMockRecorder) GetRelays() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRelays", reflect.TypeOf((*MockLocalPeer)(nil).GetRelays))
}

// PutRelays mocks base method
func (m *MockLocalPeer) PutRelays(arg0 ...*peer.Peer) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "PutRelays", varargs...)
}

// PutRelays indicates an expected call of PutRelays
func (mr *MockLocalPeerMockRecorder) PutRelays(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutRelays", reflect.TypeOf((*MockLocalPeer)(nil).PutRelays), arg0...)
}
