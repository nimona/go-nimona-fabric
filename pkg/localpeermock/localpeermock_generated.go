// Code generated by MockGen. DO NOT EDIT.
// Source: localpeer.go

// Package localpeermock is a generated GoMock package.
package localpeermock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	crypto "nimona.io/pkg/crypto"
	localpeer "nimona.io/pkg/localpeer"
	object "nimona.io/pkg/object"
	peer "nimona.io/pkg/peer"
)

// MockLocalPeer is a mock of LocalPeer interface.
type MockLocalPeer struct {
	ctrl     *gomock.Controller
	recorder *MockLocalPeerMockRecorder
}

// MockLocalPeerMockRecorder is the mock recorder for MockLocalPeer.
type MockLocalPeerMockRecorder struct {
	mock *MockLocalPeer
}

// NewMockLocalPeer creates a new mock instance.
func NewMockLocalPeer(ctrl *gomock.Controller) *MockLocalPeer {
	mock := &MockLocalPeer{ctrl: ctrl}
	mock.recorder = &MockLocalPeerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLocalPeer) EXPECT() *MockLocalPeerMockRecorder {
	return m.recorder
}

// ConnectionInfo mocks base method.
func (m *MockLocalPeer) ConnectionInfo() *peer.ConnectionInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConnectionInfo")
	ret0, _ := ret[0].(*peer.ConnectionInfo)
	return ret0
}

// ConnectionInfo indicates an expected call of ConnectionInfo.
func (mr *MockLocalPeerMockRecorder) ConnectionInfo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectionInfo", reflect.TypeOf((*MockLocalPeer)(nil).ConnectionInfo))
}

// GetAddresses mocks base method.
func (m *MockLocalPeer) GetAddresses() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAddresses")
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetAddresses indicates an expected call of GetAddresses.
func (mr *MockLocalPeerMockRecorder) GetAddresses() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAddresses", reflect.TypeOf((*MockLocalPeer)(nil).GetAddresses))
}

// GetCIDs mocks base method.
func (m *MockLocalPeer) GetCIDs() []object.CID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCIDs")
	ret0, _ := ret[0].([]object.CID)
	return ret0
}

// GetCIDs indicates an expected call of GetCIDs.
func (mr *MockLocalPeerMockRecorder) GetCIDs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCIDs", reflect.TypeOf((*MockLocalPeer)(nil).GetCIDs))
}

// GetCertificates mocks base method.
func (m *MockLocalPeer) GetCertificates() []*object.Certificate {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCertificates")
	ret0, _ := ret[0].([]*object.Certificate)
	return ret0
}

// GetCertificates indicates an expected call of GetCertificates.
func (mr *MockLocalPeerMockRecorder) GetCertificates() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCertificates", reflect.TypeOf((*MockLocalPeer)(nil).GetCertificates))
}

// GetContentTypes mocks base method.
func (m *MockLocalPeer) GetContentTypes() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContentTypes")
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetContentTypes indicates an expected call of GetContentTypes.
func (mr *MockLocalPeerMockRecorder) GetContentTypes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContentTypes", reflect.TypeOf((*MockLocalPeer)(nil).GetContentTypes))
}

// GetPrimaryIdentityKey mocks base method.
func (m *MockLocalPeer) GetPrimaryIdentityKey() crypto.PrivateKey {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPrimaryIdentityKey")
	ret0, _ := ret[0].(crypto.PrivateKey)
	return ret0
}

// GetPrimaryIdentityKey indicates an expected call of GetPrimaryIdentityKey.
func (mr *MockLocalPeerMockRecorder) GetPrimaryIdentityKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPrimaryIdentityKey", reflect.TypeOf((*MockLocalPeer)(nil).GetPrimaryIdentityKey))
}

// GetPrimaryPeerKey mocks base method.
func (m *MockLocalPeer) GetPrimaryPeerKey() crypto.PrivateKey {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPrimaryPeerKey")
	ret0, _ := ret[0].(crypto.PrivateKey)
	return ret0
}

// GetPrimaryPeerKey indicates an expected call of GetPrimaryPeerKey.
func (mr *MockLocalPeerMockRecorder) GetPrimaryPeerKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPrimaryPeerKey", reflect.TypeOf((*MockLocalPeer)(nil).GetPrimaryPeerKey))
}

// GetRelays mocks base method.
func (m *MockLocalPeer) GetRelays() []peer.ConnectionInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRelays")
	ret0, _ := ret[0].([]peer.ConnectionInfo)
	return ret0
}

// GetRelays indicates an expected call of GetRelays.
func (mr *MockLocalPeerMockRecorder) GetRelays() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRelays", reflect.TypeOf((*MockLocalPeer)(nil).GetRelays))
}

// ListenForUpdates mocks base method.
func (m *MockLocalPeer) ListenForUpdates() (<-chan localpeer.UpdateEvent, func()) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListenForUpdates")
	ret0, _ := ret[0].(<-chan localpeer.UpdateEvent)
	ret1, _ := ret[1].(func())
	return ret0, ret1
}

// ListenForUpdates indicates an expected call of ListenForUpdates.
func (mr *MockLocalPeerMockRecorder) ListenForUpdates() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListenForUpdates", reflect.TypeOf((*MockLocalPeer)(nil).ListenForUpdates))
}

// PutAddresses mocks base method.
func (m *MockLocalPeer) PutAddresses(arg0 ...string) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "PutAddresses", varargs...)
}

// PutAddresses indicates an expected call of PutAddresses.
func (mr *MockLocalPeerMockRecorder) PutAddresses(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutAddresses", reflect.TypeOf((*MockLocalPeer)(nil).PutAddresses), arg0...)
}

// PutCIDs mocks base method.
func (m *MockLocalPeer) PutCIDs(arg0 ...object.CID) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "PutCIDs", varargs...)
}

// PutCIDs indicates an expected call of PutCIDs.
func (mr *MockLocalPeerMockRecorder) PutCIDs(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutCIDs", reflect.TypeOf((*MockLocalPeer)(nil).PutCIDs), arg0...)
}

// PutCertificate mocks base method.
func (m *MockLocalPeer) PutCertificate(arg0 *object.Certificate) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PutCertificate", arg0)
}

// PutCertificate indicates an expected call of PutCertificate.
func (mr *MockLocalPeerMockRecorder) PutCertificate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutCertificate", reflect.TypeOf((*MockLocalPeer)(nil).PutCertificate), arg0)
}

// PutContentTypes mocks base method.
func (m *MockLocalPeer) PutContentTypes(arg0 ...string) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "PutContentTypes", varargs...)
}

// PutContentTypes indicates an expected call of PutContentTypes.
func (mr *MockLocalPeerMockRecorder) PutContentTypes(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutContentTypes", reflect.TypeOf((*MockLocalPeer)(nil).PutContentTypes), arg0...)
}

// PutPrimaryIdentityKey mocks base method.
func (m *MockLocalPeer) PutPrimaryIdentityKey(arg0 crypto.PrivateKey) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PutPrimaryIdentityKey", arg0)
}

// PutPrimaryIdentityKey indicates an expected call of PutPrimaryIdentityKey.
func (mr *MockLocalPeerMockRecorder) PutPrimaryIdentityKey(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutPrimaryIdentityKey", reflect.TypeOf((*MockLocalPeer)(nil).PutPrimaryIdentityKey), arg0)
}

// PutPrimaryPeerKey mocks base method.
func (m *MockLocalPeer) PutPrimaryPeerKey(arg0 crypto.PrivateKey) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PutPrimaryPeerKey", arg0)
}

// PutPrimaryPeerKey indicates an expected call of PutPrimaryPeerKey.
func (mr *MockLocalPeerMockRecorder) PutPrimaryPeerKey(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutPrimaryPeerKey", reflect.TypeOf((*MockLocalPeer)(nil).PutPrimaryPeerKey), arg0)
}

// PutRelays mocks base method.
func (m *MockLocalPeer) PutRelays(arg0 ...peer.ConnectionInfo) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "PutRelays", varargs...)
}

// PutRelays indicates an expected call of PutRelays.
func (mr *MockLocalPeerMockRecorder) PutRelays(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutRelays", reflect.TypeOf((*MockLocalPeer)(nil).PutRelays), arg0...)
}
