// Code generated by nimona.io/tools/objectify. DO NOT EDIT.

// +build !generate

package dht

import (
	"nimona.io/pkg/crypto"
	"nimona.io/pkg/object"
)

// ToMap returns a map compatible with f12n
func (s ProviderRequest) ToMap() map[string]interface{} {
	m := map[string]interface{}{
		"@ctx:s":      "nimona.io/dht/provider.request",
		"requestID:s": s.RequestID,
		"key:s":       s.Key,
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
func (s ProviderRequest) ToObject() *object.Object {
	return object.NewObjectFromMap(s.ToMap())
}

// FromMap populates the struct from a f12n compatible map
func (s *ProviderRequest) FromMap(m map[string]interface{}) error {
	if v, ok := m["requestID:s"].(string); ok {
		s.RequestID = v
	}
	if v, ok := m["key:s"].(string); ok {
		s.Key = v
	}
	s.RawObject = object.NewObjectFromMap(m)
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
func (s *ProviderRequest) FromObject(o *object.Object) error {
	return s.FromMap(o.ToMap())
}

// GetType returns the object's type
func (s ProviderRequest) GetType() string {
	return "nimona.io/dht/provider.request"
}
