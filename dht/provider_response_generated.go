// Code generated by nimona.io/go/cmd/objectify. DO NOT EDIT.

// +build !generate

package dht

import (
	"nimona.io/go/crypto"
	"nimona.io/go/encoding"
	"nimona.io/go/peers"
)

// ToMap returns a map compatible with f12n
func (s ProviderResponse) ToMap() map[string]interface{} {
	m := map[string]interface{}{
		"@ctx:s":      "nimona.io/dht/provider.response",
		"requestID:s": s.RequestID,
	}
	if s.Providers != nil {
		sProviders := []map[string]interface{}{}
		for _, v := range s.Providers {
			sProviders = append(sProviders, v.ToMap())
		}
		m["providers:a<o>"] = sProviders
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
func (s ProviderResponse) ToObject() *encoding.Object {
	return encoding.NewObjectFromMap(s.ToMap())
}

// FromMap populates the struct from a f12n compatible map
func (s *ProviderResponse) FromMap(m map[string]interface{}) error {
	if v, ok := m["requestID:s"].(string); ok {
		s.RequestID = v
	}
	s.Providers = []*Provider{}
	if ss, ok := m["providers:a<o>"].([]interface{}); ok {
		for _, si := range ss {
			if v, ok := si.(map[string]interface{}); ok {
				sProviders := &Provider{}
				if err := sProviders.FromMap(v); err != nil {
					return err
				}
				s.Providers = append(s.Providers, sProviders)
			} else if v, ok := m["providers:a<o>"].(*Provider); ok {
				s.Providers = append(s.Providers, v)
			}
		}
	}
	s.ClosestPeers = []*peers.PeerInfo{}
	if ss, ok := m["closestPeers:a<o>"].([]interface{}); ok {
		for _, si := range ss {
			if v, ok := si.(map[string]interface{}); ok {
				sClosestPeers := &peers.PeerInfo{}
				if err := sClosestPeers.FromMap(v); err != nil {
					return err
				}
				s.ClosestPeers = append(s.ClosestPeers, sClosestPeers)
			} else if v, ok := m["closestPeers:a<o>"].(*peers.PeerInfo); ok {
				s.ClosestPeers = append(s.ClosestPeers, v)
			}
		}
	}
	s.RawObject = encoding.NewObjectFromMap(m)
	if v, ok := m["@signer:o"].(map[string]interface{}); ok {
		s.Signer = &crypto.Key{}
		if err := s.Signer.FromMap(v); err != nil {
			return err
		}
	} else if v, ok := m["@signer:o"].(*crypto.Key); ok {
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
	if v, ok := m["@signature:o"].(map[string]interface{}); ok {
		s.Signature = &crypto.Signature{}
		if err := s.Signature.FromMap(v); err != nil {
			return err
		}
	} else if v, ok := m["@signature:o"].(*crypto.Signature); ok {
		s.Signature = v
	}
	return nil
}

// FromObject populates the struct from a f12n object
func (s *ProviderResponse) FromObject(o *encoding.Object) error {
	return s.FromMap(o.ToMap())
}

// GetType returns the object's type
func (s ProviderResponse) GetType() string {
	return "nimona.io/dht/provider.response"
}
