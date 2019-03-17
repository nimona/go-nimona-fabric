package net

import (
	"crypto/ecdsa"
	"sync"

	"nimona.io/pkg/crypto"
	"nimona.io/pkg/net/peer"
)

type NetState interface {
	AttachMandate(m *crypto.Mandate) error
	AddAddress(addrs ...string)
	GetPeerInfo() *peer.PeerInfo
	GetPeerKey() *crypto.Key
}

type LocalInfo struct {
	key            *crypto.Key
	mandate        *crypto.Mandate
	addressesLock  sync.RWMutex
	addresses      []string
	relayAddresses []string
}

func NewLocalInfo(key *crypto.Key, relayAddresses []string) (
	NetState, error) {
	if key == nil {
		return nil, ErrMissingKey
	}

	if _, ok := key.Materialize().(*ecdsa.PrivateKey); !ok {
		return nil, ErrECDSAPrivateKeyRequired
	}

	return &LocalInfo{
		key:            key,
		relayAddresses: relayAddresses,
	}, nil
}

func (l *LocalInfo) AttachMandate(m *crypto.Mandate) error {
	// TODO(geoah): Check if our peer key is the mandate's subject
	l.addressesLock.Lock()
	l.mandate = m
	l.addressesLock.Unlock()
	return nil
}

func (l *LocalInfo) AddAddress(addrs ...string) {
	l.addressesLock.Lock()
	if l.addresses == nil {
		l.addresses = []string{}
	}
	l.addresses = append(l.addresses, addrs...)
	l.addressesLock.Unlock()
}

func (l *LocalInfo) GetPeerKey() *crypto.Key {
	return l.key
}

// GetPeerInfo returns the local peer info
func (l *LocalInfo) GetPeerInfo() *peer.PeerInfo {
	// TODO cache peer info and reuse
	p := &peer.PeerInfo{
		SignerKey: l.key.GetPublicKey(),
	}

	l.addressesLock.RLock()
	// TODO Check all the transports for addresses
	p.Addresses = l.addresses
	if l.mandate != nil {
		p.AuthorityKey = l.mandate.Signer
		p.Mandate = l.mandate
	}
	l.addressesLock.RUnlock()

	o := p.ToObject()
	if err := crypto.Sign(o, l.key); err != nil {
		panic(err)
	}
	p.FromObject(o)
	return p
}
