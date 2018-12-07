// Code generated by nimona.io/go/cmd/objectify. DO NOT EDIT.

// +build !generate

package net

import (
	"nimona.io/go/encoding"
	"nimona.io/go/peers"
)

// ToMap returns a map compatible with f12n
func (s HandshakeSyn) ToMap() map[string]interface{} {
	m := map[string]interface{}{
		"@ctx:s":  "/handshake.syn",
		"nonce:s": s.Nonce,
	}
	if s.PeerInfo != nil {
		m["peerInfo:o"] = s.PeerInfo.ToMap()
	}
	return m
}

// ToObject returns a f12n object
func (s HandshakeSyn) ToObject() *encoding.Object {
	return encoding.NewObjectFromMap(s.ToMap())
}

// FromMap populates the struct from a f12n compatible map
func (s *HandshakeSyn) FromMap(m map[string]interface{}) error {
	s.RawObject = encoding.NewObjectFromMap(m)
	if v, ok := m["@:o"].(*encoding.Object); ok {
		s.RawObject = v
	}
	if v, ok := m["nonce:s"].(string); ok {
		s.Nonce = v
	}
	if v, ok := m["peerInfo:o"].(map[string]interface{}); ok {
		s.PeerInfo = &peers.PeerInfo{}
		if err := s.PeerInfo.FromMap(v); err != nil {
			return err
		}
	} else if v, ok := m["peerInfo:o"].(*peers.PeerInfo); ok {
		s.PeerInfo = v
	}
	if v, ok := m["peerInfo:o"].(*peers.PeerInfo); ok {
		s.PeerInfo = v
	}
	return nil
}

// FromObject populates the struct from a f12n object
func (s *HandshakeSyn) FromObject(o *encoding.Object) error {
	return s.FromMap(o.ToMap())
}

// GetType returns the object's type
func (s HandshakeSyn) GetType() string {
	return "/handshake.syn"
}
