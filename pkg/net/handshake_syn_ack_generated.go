// Code generated by nimona.io/tools/objectify. DO NOT EDIT.

// +build !generate

package net

import (
	"nimona.io/pkg/crypto"
	"nimona.io/pkg/object"
	"nimona.io/pkg/net/peer"
)

// ToMap returns a map compatible with f12n
func (s HandshakeSynAck) ToMap() map[string]interface{} {
	m := map[string]interface{}{
		"@ctx:s":  "/handshake.syn-ack",
		"nonce:s": s.Nonce,
	}
	if s.PeerInfo != nil {
		m["peerInfo:o"] = s.PeerInfo.ToMap()
	}
	if s.Signer != nil {
		m["@signer:o"] = s.Signer.ToMap()
	}
	if s.Signature != nil {
		m["@signature:o"] = s.Signature.ToMap()
	}
	return m
}

// ToObject returns a f12n object
func (s HandshakeSynAck) ToObject() *object.Object {
	return object.FromMap(s.ToMap())
}

// FromMap populates the struct from a f12n compatible map
func (s *HandshakeSynAck) FromMap(m map[string]interface{}) error {
	if v, ok := m["nonce:s"].(string); ok {
		s.Nonce = v
	}
	if v, ok := m["peerInfo:o"].(map[string]interface{}); ok {
		s.PeerInfo = &peer.PeerInfo{}
		if err := s.PeerInfo.FromMap(v); err != nil {
			return err
		}
	} else if v, ok := m["peerInfo:o"].(*peer.PeerInfo); ok {
		s.PeerInfo = v
	}
	if v, ok := m["peerInfo:o"].(*peer.PeerInfo); ok {
		s.PeerInfo = v
	}
	if v, ok := m["@signer:o"].(map[string]interface{}); ok {
		s.Signer = &crypto.Key{}
		if err := s.Signer.FromMap(v); err != nil {
			return err
		}
	} else if v, ok := m["@signer:o"].(*crypto.Key); ok {
		s.Signer = v
	}
	if v, ok := m["@signer:o"].(*crypto.Key); ok {
		s.Signer = v
	}
	if v, ok := m["@signature:o"].(map[string]interface{}); ok {
		s.Signature = &crypto.Signature{}
		if err := s.Signature.FromMap(v); err != nil {
			return err
		}
	} else if v, ok := m["@signature:o"].(*crypto.Signature); ok {
		s.Signature = v
	}
	if v, ok := m["@signature:o"].(*crypto.Signature); ok {
		s.Signature = v
	}
	return nil
}

// FromObject populates the struct from a f12n object
func (s *HandshakeSynAck) FromObject(o *object.Object) error {
	return s.FromMap(o.ToMap())
}

// GetType returns the object's type
func (s HandshakeSynAck) GetType() string {
	return "/handshake.syn-ack"
}
