// Code generated by nimona.io/tools/objectify. DO NOT EDIT.

// +build !generate

package dht

import (
	"nimona.io/pkg/crypto"
	"nimona.io/pkg/net/peer"
	"nimona.io/pkg/object"
)

// ToMap returns a map compatible with f12n
func (s PeerInfoResponse) ToMap() map[string]interface{} {
	m := map[string]interface{}{
		"@ctx:s":      "nimona.io/dht/peerinfo.response",
		"requestID:s": s.RequestID,
	}
	if s.PeerInfo != nil {
		m["peerInfo:o"] = s.PeerInfo.ToMap()
	}
	if s.ClosestPeers != nil {
		sClosestPeers := []map[string]interface{}{}
		for _, v := range s.ClosestPeers {
			sClosestPeers = append(sClosestPeers, v.ToMap())
		}
		m["closestPeers:a<o>"] = sClosestPeers
	}
	if s.Signer != nil {
		m["@signer:o"] = s.Signer.ToMap()
	}
	if s.Authority != nil {
		m["@authority:o"] = s.Authority.ToMap()
	}
	if s.Signature != nil {
		m["@signature:o"] = s.Signature.ToMap()
	}
	return m
}

// ToObject returns a f12n object
func (s PeerInfoResponse) ToObject() *object.Object {
	return object.FromMap(s.ToMap())
}

// FromMap populates the struct from a f12n compatible map
func (s *PeerInfoResponse) FromMap(m map[string]interface{}) error {
	if v, ok := m["requestID:s"].(string); ok {
		s.RequestID = v
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
	s.ClosestPeers = []*peer.PeerInfo{}
	if ss, ok := m["closestPeers:a<o>"].([]interface{}); ok {
		for _, si := range ss {
			if v, ok := si.(*peer.PeerInfo); ok {
				s.ClosestPeers = append(s.ClosestPeers, v)
			} else if v, ok := si.(map[string]interface{}); ok {
				sClosestPeers := &peer.PeerInfo{}
				if err := sClosestPeers.FromMap(v); err != nil {
					return err
				}
				s.ClosestPeers = append(s.ClosestPeers, sClosestPeers)
			}
		}
	}
	if v, ok := m["closestPeers:a<o>"].([]*peer.PeerInfo); ok {
		s.ClosestPeers = v
	}
	s.RawObject = object.FromMap(m)
	if v, ok := m["@:o"].(*object.Object); ok {
		s.RawObject = v
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
	if v, ok := m["@authority:o"].(map[string]interface{}); ok {
		s.Authority = &crypto.Key{}
		if err := s.Authority.FromMap(v); err != nil {
			return err
		}
	} else if v, ok := m["@authority:o"].(*crypto.Key); ok {
		s.Authority = v
	}
	if v, ok := m["@authority:o"].(*crypto.Key); ok {
		s.Authority = v
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
func (s *PeerInfoResponse) FromObject(o *object.Object) error {
	return s.FromMap(o.ToMap())
}

// GetType returns the object's type
func (s PeerInfoResponse) GetType() string {
	return "nimona.io/dht/peerinfo.response"
}
