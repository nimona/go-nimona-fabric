// Code generated by MockGen. DO NOT EDIT.
// Source: keychain.go

// Package keychainmock is a generated GoMock package.
package keychainmock

import (
	gomock "github.com/golang/mock/gomock"
	crypto "nimona.io/pkg/crypto"
	peer "nimona.io/pkg/peer"
	reflect "reflect"
)

// MockKeychain is a mock of Keychain interface
type MockKeychain struct {
	ctrl     *gomock.Controller
	recorder *MockKeychainMockRecorder
}

// MockKeychainMockRecorder is the mock recorder for MockKeychain
type MockKeychainMockRecorder struct {
	mock *MockKeychain
}

// NewMockKeychain creates a new mock instance
func NewMockKeychain(ctrl *gomock.Controller) *MockKeychain {
	mock := &MockKeychain{ctrl: ctrl}
	mock.recorder = &MockKeychainMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockKeychain) EXPECT() *MockKeychainMockRecorder {
	return m.recorder
}

// GetPrimaryPeerKey mocks base method
func (m *MockKeychain) GetPrimaryPeerKey() crypto.PrivateKey {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPrimaryPeerKey")
	ret0, _ := ret[0].(crypto.PrivateKey)
	return ret0
}

// GetPrimaryPeerKey indicates an expected call of GetPrimaryPeerKey
func (mr *MockKeychainMockRecorder) GetPrimaryPeerKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPrimaryPeerKey", reflect.TypeOf((*MockKeychain)(nil).GetPrimaryPeerKey))
}

// PutPrimaryPeerKey mocks base method
func (m *MockKeychain) PutPrimaryPeerKey(arg0 crypto.PrivateKey) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PutPrimaryPeerKey", arg0)
}

// PutPrimaryPeerKey indicates an expected call of PutPrimaryPeerKey
func (mr *MockKeychainMockRecorder) PutPrimaryPeerKey(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutPrimaryPeerKey", reflect.TypeOf((*MockKeychain)(nil).PutPrimaryPeerKey), arg0)
}

// GetPrimaryIdentityKey mocks base method
func (m *MockKeychain) GetPrimaryIdentityKey() crypto.PrivateKey {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPrimaryIdentityKey")
	ret0, _ := ret[0].(crypto.PrivateKey)
	return ret0
}

// GetPrimaryIdentityKey indicates an expected call of GetPrimaryIdentityKey
func (mr *MockKeychainMockRecorder) GetPrimaryIdentityKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPrimaryIdentityKey", reflect.TypeOf((*MockKeychain)(nil).GetPrimaryIdentityKey))
}

// PutPrimaryIdentityKey mocks base method
func (m *MockKeychain) PutPrimaryIdentityKey(arg0 crypto.PrivateKey) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PutPrimaryIdentityKey", arg0)
}

// PutPrimaryIdentityKey indicates an expected call of PutPrimaryIdentityKey
func (mr *MockKeychainMockRecorder) PutPrimaryIdentityKey(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutPrimaryIdentityKey", reflect.TypeOf((*MockKeychain)(nil).PutPrimaryIdentityKey), arg0)
}

// GetCertificates mocks base method
func (m *MockKeychain) GetCertificates(arg0 crypto.PublicKey) []*peer.Certificate {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCertificates", arg0)
	ret0, _ := ret[0].([]*peer.Certificate)
	return ret0
}

// GetCertificates indicates an expected call of GetCertificates
func (mr *MockKeychainMockRecorder) GetCertificates(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCertificates", reflect.TypeOf((*MockKeychain)(nil).GetCertificates), arg0)
}

// PutCertificate mocks base method
func (m *MockKeychain) PutCertificate(arg0 *peer.Certificate) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PutCertificate", arg0)
}

// PutCertificate indicates an expected call of PutCertificate
func (mr *MockKeychainMockRecorder) PutCertificate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutCertificate", reflect.TypeOf((*MockKeychain)(nil).PutCertificate), arg0)
}
