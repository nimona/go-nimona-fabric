// Code generated by nimona.io/tools/codegen. DO NOT EDIT.

package main

import (
	"errors"

	crypto "nimona.io/pkg/crypto"
	immutable "nimona.io/pkg/immutable"
	object "nimona.io/pkg/object"
)

type (
	Msg struct {
		raw        object.Object
		Stream     object.Hash
		Parents    []object.Hash
		Owners     []crypto.PublicKey
		Policy     object.Policy
		Signatures []object.Signature
		Datetime   int64
		Body       string
	}
)

func (e Msg) GetType() string {
	return "nimona.io/msg"
}

func (e Msg) GetSchema() *object.SchemaObject {
	return &object.SchemaObject{
		Properties: []*object.SchemaProperty{
			&object.SchemaProperty{
				Name:       "datetime",
				Type:       "int",
				Hint:       "i",
				IsRepeated: false,
				IsOptional: false,
			},
			&object.SchemaProperty{
				Name:       "body",
				Type:       "string",
				Hint:       "s",
				IsRepeated: false,
				IsOptional: false,
			},
		},
	}
}

func (e Msg) ToObject() object.Object {
	o := object.Object{}
	o = o.SetType("nimona.io/msg")
	if len(e.Stream) > 0 {
		o = o.SetStream(e.Stream)
	}
	if len(e.Parents) > 0 {
		o = o.SetParents(e.Parents)
	}
	if len(e.Owners) > 0 {
		o = o.SetOwners(e.Owners)
	}
	o = o.AddSignature(e.Signatures...)
	o = o.SetPolicy(e.Policy)
	o = o.Set("datetime:i", e.Datetime)
	if e.Body != "" {
		o = o.Set("body:s", e.Body)
	}
	// if schema := e.GetSchema(); schema != nil {
	// 	m["_schema:o"] = schema.ToObject().ToMap()
	// }
	return o
}

func (e *Msg) FromObject(o object.Object) error {
	data, ok := o.Raw().Value("data:o").(immutable.Map)
	if !ok {
		return errors.New("missing data")
	}
	e.raw = object.Object{}
	e.raw = e.raw.SetType(o.GetType())
	e.Stream = o.GetStream()
	e.Parents = o.GetParents()
	e.Owners = o.GetOwners()
	e.Signatures = o.GetSignatures()
	e.Policy = o.GetPolicy()
	if v := data.Value("datetime:i"); v != nil {
		e.Datetime = int64(v.PrimitiveHinted().(int64))
	}
	if v := data.Value("body:s"); v != nil {
		e.Body = string(v.PrimitiveHinted().(string))
	}
	return nil
}