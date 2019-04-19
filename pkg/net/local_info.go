package net

import (
	"crypto/ecdsa"
	"sync"

	"nimona.io/pkg/crypto"
	"nimona.io/pkg/net/peer"
)

type LocalInfo struct {
	hostname      string
	key           *crypto.Key
	mandate       *crypto.Mandate
	addressesLock sync.RWMutex
	addresses     []string

	// TODO replace with generated string-bool syncmap
	contentHashesLock sync.RWMutex
	contentHashes     map[string]bool // map[hash]publishable
}

func NewLocalInfo(hostname string, key *crypto.Key) (
	*LocalInfo, error) {
	if key == nil {
		return nil, ErrMissingKey
	}

	if _, ok := key.Materialize().(*ecdsa.PrivateKey); !ok {
		return nil, ErrECDSAPrivateKeyRequired
	}

	return &LocalInfo{
		hostname:      hostname,
		key:           key,
		addresses:     []string{},
		contentHashes: map[string]bool{},
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

// AddContentHash that should be published with the peer info
func (l *LocalInfo) AddContentHash(hashes ...string) {
	l.contentHashesLock.Lock()
	for _, hash := range hashes {
		l.contentHashes[hash] = true
	}
	l.contentHashesLock.Unlock()
}

// RemoveContentHash from the peer info
func (l *LocalInfo) RemoveContentHash(hashes ...string) {
	l.contentHashesLock.Lock()
	for _, hash := range hashes {
		delete(l.contentHashes, hash)
	}
	l.contentHashesLock.Unlock()
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
	addresses := make([]string, len(l.addresses))
	for i, a := range l.addresses {
		addresses[i] = a
	}
	p.Addresses = addresses
	if l.mandate != nil {
		p.AuthorityKey = l.mandate.Signer
		p.Mandate = l.mandate
	}
	l.addressesLock.RUnlock()

	l.contentHashesLock.RLock()
	hashes := []string{}
	for hash, publishable := range l.contentHashes {
		if !publishable {
			continue
		}
		hashes = append(hashes, hash)
	}
	p.ContentIDs = hashes
	l.contentHashesLock.RUnlock()

	o := p.ToObject()
	if err := crypto.Sign(o, l.key); err != nil {
		panic(err)
	}
	p.FromObject(o)
	return p
}

func (l *LocalInfo) GetHostname() string {
	return l.hostname
}
