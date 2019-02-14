// Code generated by nimona.io/tools/objectify. DO NOT EDIT.

// +build !generate

package telemetry

import (
	"nimona.io/pkg/object"
)

const (
	ConnectionEventType = "nimona.io/telemetry/connection"
)

// ToObject returns a f12n object
func (s ConnectionEvent) ToObject() *object.Object {
	o := object.New()
	o.SetType(ConnectionEventType)
	if s.Direction != "" {
		o.SetRaw("direction", s.Direction)
	}
	return o
}

// FromObject populates the struct from a f12n object
func (s *ConnectionEvent) FromObject(o *object.Object) error {
	if v, ok := o.GetRaw("direction").(string); ok {
		s.Direction = v
	}
	return nil
}

// GetType returns the object's type
func (s ConnectionEvent) GetType() string {
	return ConnectionEventType
}
