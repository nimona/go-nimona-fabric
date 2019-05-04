// Code generated by nimona.io/tools/objectify. DO NOT EDIT.

// +build !generate

package crypto

import (
	"github.com/mitchellh/mapstructure"
	"nimona.io/pkg/object"
)

const (
	PrivateKeyType = "/key.private"
)

// ToObject returns a f12n object
func (s PrivateKey) ToObject() *object.Object {
	o := object.New()
	o.SetType(PrivateKeyType)
	if s.Algorithm != "" {
		o.SetRaw("alg", s.Algorithm)
	}
	if s.KeyType != "" {
		o.SetRaw("kty", s.KeyType)
	}
	if s.Curve != "" {
		o.SetRaw("crv", s.Curve)
	}
	if len(s.X) > 0 {
		o.SetRaw("x", s.X)
	}
	if len(s.Y) > 0 {
		o.SetRaw("y", s.Y)
	}
	if len(s.D) > 0 {
		o.SetRaw("d", s.D)
	}
	if s.PublicKey != nil {
		o.SetRaw("pub", s.PublicKey)
	}
	return o
}

func anythingToAnythingForPrivateKey(
	from interface{},
	to interface{},
) error {
	config := &mapstructure.DecoderConfig{
		Result:  to,
		TagName: "json",
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}

	if err := decoder.Decode(from); err != nil {
		return err
	}

	return nil
}

// FromObject populates the struct from a f12n object
func (s *PrivateKey) FromObject(o *object.Object) error {
	atoa := anythingToAnythingForPrivateKey
	if err := atoa(o.GetRaw("alg"), &s.Algorithm); err != nil {
		return err
	}
	if err := atoa(o.GetRaw("kty"), &s.KeyType); err != nil {
		return err
	}
	if err := atoa(o.GetRaw("crv"), &s.Curve); err != nil {
		return err
	}
	if err := atoa(o.GetRaw("x"), &s.X); err != nil {
		return err
	}
	if err := atoa(o.GetRaw("y"), &s.Y); err != nil {
		return err
	}
	if err := atoa(o.GetRaw("d"), &s.D); err != nil {
		return err
	}
	if v, ok := o.GetRaw("pub").(*PublicKey); ok {
		s.PublicKey = v
	} else if v, ok := o.GetRaw("pub").(map[string]interface{}); ok {
		s.PublicKey = &PublicKey{}
		o := &object.Object{}
		if err := o.FromMap(v); err != nil {
			return err
		}
		s.PublicKey.FromObject(o)
	}

	if ao, ok := interface{}(s).(interface{ afterFromObject() }); ok {
		ao.afterFromObject()
	}

	return nil
}

// GetType returns the object's type
func (s PrivateKey) GetType() string {
	return PrivateKeyType
}
