// Code generated by nimona.io/tools/codegen. DO NOT EDIT.

package exchange

import (
	json "encoding/json"

	crypto "nimona.io/pkg/crypto"
	object "nimona.io/pkg/object"
	schema "nimona.io/pkg/schema"
)

type (
	ObjectRequest struct {
		ObjectHash object.Hash       `json:"objectHash:s,omitempty"`
		Signature  *crypto.Signature `json:"@signature:o,omitempty"`
		Identity   crypto.PublicKey  `json:"@identity:s,omitempty"`
	}
	ObjectForward struct {
		Recipient string            `json:"recipient:s,omitempty"`
		FwObject  *object.Object    `json:"fwObject:o,omitempty"`
		Signature *crypto.Signature `json:"@signature:o,omitempty"`
		Identity  crypto.PublicKey  `json:"@identity:s,omitempty"`
	}
)

func (e ObjectRequest) GetType() string {
	return "nimona.io/exchange.ObjectRequest"
}

func (e ObjectRequest) GetSchema() *schema.Object {
	return &schema.Object{
		Properties: []*schema.Property{
			&schema.Property{
				Name:       "objectHash",
				Type:       "nimona.io/object.Hash",
				Hint:       "s",
				IsRepeated: false,
				IsOptional: false,
			},
			&schema.Property{
				Name:       "@signature",
				Type:       "nimona.io/crypto.Signature",
				Hint:       "o",
				IsRepeated: false,
				IsOptional: false,
			},
			&schema.Property{
				Name:       "@identity",
				Type:       "nimona.io/crypto.PublicKey",
				Hint:       "s",
				IsRepeated: false,
				IsOptional: false,
			},
		},
		Links: []*schema.Link{},
	}
}

func (e ObjectRequest) ToObject() object.Object {
	m := map[string]interface{}{}
	m["@type:s"] = "nimona.io/exchange.ObjectRequest"
	if e.ObjectHash != "" {
		m["objectHash:s"] = e.ObjectHash
	}
	if e.Signature != nil {
		m["@signature:o"] = e.Signature.ToObject().ToMap()
	}
	if e.Identity != "" {
		m["@identity:s"] = e.Identity
	}

	if schema := e.GetSchema(); schema != nil {
		m["$schema:o"] = schema.ToObject().ToMap()
	}
	return object.Object(m)
}

func (e *ObjectRequest) FromObject(o object.Object) error {
	b, _ := json.Marshal(map[string]interface{}(o))
	return json.Unmarshal(b, e)
}

func (e ObjectForward) GetType() string {
	return "nimona.io/exchange.ObjectForward"
}

func (e ObjectForward) GetSchema() *schema.Object {
	return &schema.Object{
		Properties: []*schema.Property{
			&schema.Property{
				Name:       "recipient",
				Type:       "string",
				Hint:       "s",
				IsRepeated: false,
				IsOptional: false,
			},
			&schema.Property{
				Name:       "fwObject",
				Type:       "nimona.io/object.Object",
				Hint:       "o",
				IsRepeated: false,
				IsOptional: false,
			},
			&schema.Property{
				Name:       "@signature",
				Type:       "nimona.io/crypto.Signature",
				Hint:       "o",
				IsRepeated: false,
				IsOptional: false,
			},
			&schema.Property{
				Name:       "@identity",
				Type:       "nimona.io/crypto.PublicKey",
				Hint:       "s",
				IsRepeated: false,
				IsOptional: false,
			},
		},
		Links: []*schema.Link{},
	}
}

func (e ObjectForward) ToObject() object.Object {
	m := map[string]interface{}{}
	m["@type:s"] = "nimona.io/exchange.ObjectForward"
	if e.Recipient != "" {
		m["recipient:s"] = e.Recipient
	}
	if e.FwObject != nil {
		m["fwObject:o"] = e.FwObject.ToObject().ToMap()
	}
	if e.Signature != nil {
		m["@signature:o"] = e.Signature.ToObject().ToMap()
	}
	if e.Identity != "" {
		m["@identity:s"] = e.Identity
	}

	if schema := e.GetSchema(); schema != nil {
		m["$schema:o"] = schema.ToObject().ToMap()
	}
	return object.Object(m)
}

func (e *ObjectForward) FromObject(o object.Object) error {
	b, _ := json.Marshal(map[string]interface{}(o))
	return json.Unmarshal(b, e)
}
