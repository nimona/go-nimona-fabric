// Code generated by nimona.io/tools/codegen. DO NOT EDIT.

package peer

import (
	crypto "nimona.io/pkg/crypto"
	object "nimona.io/pkg/object"
)

type (
	ConnectionInfo struct {
		Metadata  object.Metadata   `nimona:"metadata:m,omitempty"`
		PublicKey crypto.PublicKey  `nimona:"publicKey:s,omitempty"`
		Addresses []string          `nimona:"addresses:as,omitempty"`
		Relays    []*ConnectionInfo `nimona:"relays:ao,omitempty"`
	}
)

func (e *ConnectionInfo) Type() string {
	return "nimona.io/peer.ConnectionInfo"
}

func (e ConnectionInfo) ToObject() *object.Object {
	r := &object.Object{
		Type:     "nimona.io/peer.ConnectionInfo",
		Metadata: e.Metadata,
		Data:     map[string]interface{}{},
	}
	r.Data["publicKey:s"] = e.PublicKey
	if len(e.Addresses) > 0 {
		r.Data["addresses:as"] = e.Addresses
	}
	if len(e.Relays) > 0 {
		rv := make([]*object.Object, len(e.Relays))
		for i, v := range e.Relays {
			rv[i] = v.ToObject()
		}
		r.Data["relays:ao"] = rv
	}
	return r
}

func (e ConnectionInfo) ToObjectMap() map[string]interface{} {
	d := map[string]interface{}{}
	d["publicKey:s"] = e.PublicKey
	if len(e.Addresses) > 0 {
		d["addresses:as"] = e.Addresses
	}
	if len(e.Relays) > 0 {
		rv := make([]*object.Object, len(e.Relays))
		for i, v := range e.Relays {
			rv[i] = v.ToObject()
		}
		d["relays:ao"] = rv
	}
	r := map[string]interface{}{
		"type:s":     "nimona.io/peer.ConnectionInfo",
		"metadata:m": object.MetadataToMap(&e.Metadata),
		"data:m":     d,
	}
	return r
}

func (e *ConnectionInfo) FromObject(o *object.Object) error {
	return object.Decode(o, e)
}
