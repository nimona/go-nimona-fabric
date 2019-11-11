// Code generated by nimona.io/tools/codegen. DO NOT EDIT.

package handshake

import (
	json "encoding/json"

	crypto "nimona.io/pkg/crypto"
	object "nimona.io/pkg/object"
)

type (
	Syn struct {
		Nonce     string            `json:"nonce:s,omitempty"`
		Signature *crypto.Signature `json:"@signature:o,omitempty"`
		Identity  crypto.PublicKey  `json:"@identity:s,omitempty"`
	}
	SynAck struct {
		Nonce     string            `json:"nonce:s,omitempty"`
		Signature *crypto.Signature `json:"@signature:o,omitempty"`
		Identity  crypto.PublicKey  `json:"@identity:s,omitempty"`
	}
	Ack struct {
		Nonce     string            `json:"nonce:s,omitempty"`
		Signature *crypto.Signature `json:"@signature:o,omitempty"`
		Identity  crypto.PublicKey  `json:"@identity:s,omitempty"`
	}
)

func (e *Syn) GetType() string {
	return "nimona.io/net/handshake.Syn"
}

func (e *Syn) ToObject() object.Object {
	m := map[string]interface{}{}
	m["@type:s"] = "nimona.io/net/handshake.Syn"
	m["nonce:s"] = e.Nonce
	if e.Signature != nil {
		m["@signature:o"] = e.Signature.ToObject().ToMap()
	}
	m["@identity:s"] = e.Identity
	return object.Object(m)
}

func (e *Syn) FromObject(o object.Object) error {
	b, _ := json.Marshal(map[string]interface{}(o))
	return json.Unmarshal(b, e)
}

func (e *SynAck) GetType() string {
	return "nimona.io/net/handshake.SynAck"
}

func (e *SynAck) ToObject() object.Object {
	m := map[string]interface{}{}
	m["@type:s"] = "nimona.io/net/handshake.SynAck"
	m["nonce:s"] = e.Nonce
	if e.Signature != nil {
		m["@signature:o"] = e.Signature.ToObject().ToMap()
	}
	m["@identity:s"] = e.Identity
	return object.Object(m)
}

func (e *SynAck) FromObject(o object.Object) error {
	b, _ := json.Marshal(map[string]interface{}(o))
	return json.Unmarshal(b, e)
}

func (e *Ack) GetType() string {
	return "nimona.io/net/handshake.Ack"
}

func (e *Ack) ToObject() object.Object {
	m := map[string]interface{}{}
	m["@type:s"] = "nimona.io/net/handshake.Ack"
	m["nonce:s"] = e.Nonce
	if e.Signature != nil {
		m["@signature:o"] = e.Signature.ToObject().ToMap()
	}
	m["@identity:s"] = e.Identity
	return object.Object(m)
}

func (e *Ack) FromObject(o object.Object) error {
	b, _ := json.Marshal(map[string]interface{}(o))
	return json.Unmarshal(b, e)
}
