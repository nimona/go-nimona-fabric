// Code generated by nimona.io/tools/codegen. DO NOT EDIT.

package peer

import (
	crypto "nimona.io/pkg/crypto"
	object "nimona.io/pkg/object"
)

type (
	ConnectionInfo struct {
		Metadata      object.Metadata
		Version       int64             `nimona:"version:i"`
		PublicKey     crypto.PublicKey  `nimona:"publicKey:s"`
		Addresses     []string          `nimona:"addresses:as"`
		Relays        []*ConnectionInfo `nimona:"relays:ao"`
		ObjectFormats []string          `nimona:"objectFormats:as"`
	}
)

func (e *ConnectionInfo) Type() string {
	return "nimona.io/peer.ConnectionInfo"
}

func (e *ConnectionInfo) MarshalMap() (object.Map, error) {
	return e.ToObject().Map(), nil
}

func (e *ConnectionInfo) MarshalObject() (*object.Object, error) {
	return e.ToObject(), nil
}

func (e ConnectionInfo) ToObject() *object.Object {
	r := &object.Object{
		Type:     "nimona.io/peer.ConnectionInfo",
		Metadata: e.Metadata,
		Data:     object.Map{},
	}
	r.Data["version"] = object.Int(e.Version)
	if v, err := e.PublicKey.MarshalString(); err == nil {
		r.Data["publicKey"] = object.String(v)
	}
	if len(e.Addresses) > 0 {
		rv := make(object.StringArray, len(e.Addresses))
		for i, iv := range e.Addresses {
			rv[i] = object.String(iv)
		}
		r.Data["addresses"] = rv
	}
	if len(e.Relays) > 0 {
		rv := make(object.ObjectArray, len(e.Relays))
		for i, v := range e.Relays {
			if iv, err := v.MarshalObject(); err == nil {
				rv[i] = (iv)
			}
		}
		r.Data["relays"] = rv
	}
	if len(e.ObjectFormats) > 0 {
		rv := make(object.StringArray, len(e.ObjectFormats))
		for i, iv := range e.ObjectFormats {
			rv[i] = object.String(iv)
		}
		r.Data["objectFormats"] = rv
	}
	return r
}

func (e *ConnectionInfo) UnmarshalMap(m object.Map) error {
	return e.FromObject(object.FromMap(m))
}

func (e *ConnectionInfo) UnmarshalObject(o *object.Object) error {
	return e.FromObject(o)
}

func (e *ConnectionInfo) FromObject(o *object.Object) error {
	e.Metadata = o.Metadata
	if v, ok := o.Data["version"]; ok {
		if t, ok := v.(object.Int); ok {
			e.Version = int64(t)
		}
	}
	if v, ok := o.Data["publicKey"]; ok {
		if ev, ok := v.(object.String); ok {
			es := crypto.PublicKey{}
			if err := es.UnmarshalString(string(ev)); err == nil {
				e.PublicKey = es
			}
		}
	}
	if v, ok := o.Data["addresses"]; ok {
		if t, ok := v.(object.StringArray); ok {
			rv := make([]string, len(t))
			for i, iv := range t {
				rv[i] = string(iv)
			}
			e.Addresses = rv
		}
	}
	if v, ok := o.Data["relays"]; ok {
		if ev, ok := v.(object.ObjectArray); ok {
			e.Relays = make([]*ConnectionInfo, len(ev))
			for i, iv := range ev {
				es := &ConnectionInfo{}
				if err := es.UnmarshalObject((iv)); err == nil {
					e.Relays[i] = es
				}
			}
		}
	}
	if v, ok := o.Data["objectFormats"]; ok {
		if t, ok := v.(object.StringArray); ok {
			rv := make([]string, len(t))
			for i, iv := range t {
				rv[i] = string(iv)
			}
			e.ObjectFormats = rv
		}
	}
	return nil
}
