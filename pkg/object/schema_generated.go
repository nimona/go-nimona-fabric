// Code generated by nimona.io/tools/codegen. DO NOT EDIT.

package object

import "nimona.io/pkg/errors"

type (
	SchemaProperty struct {
		raw        Object
		Metadata   Metadata
		Name       string
		Type       string
		Hint       string
		IsRepeated bool
		IsOptional bool
		Properties []*SchemaProperty
	}
	SchemaObject struct {
		raw        Object
		Metadata   Metadata
		Properties []*SchemaProperty
	}
)

func (e SchemaProperty) GetType() string {
	return "nimona.io/SchemaProperty"
}

func (e SchemaProperty) IsStreamRoot() bool {
	return false
}

func (e SchemaProperty) ToObject() Object {
	o := Object{}
	o = o.SetType("nimona.io/SchemaProperty")
	if len(e.Metadata.Stream) > 0 {
		o = o.SetStream(e.Metadata.Stream)
	}
	if len(e.Metadata.Parents) > 0 {
		o = o.SetParents(e.Metadata.Parents)
	}
	if len(e.Metadata.Owners) > 0 {
		o = o.SetOwners(e.Metadata.Owners)
	}
	o = o.AddSignature(e.Metadata.Signatures...)
	o = o.SetPolicy(e.Metadata.Policy)
	if e.Name != "" {
		o = o.Set("name:s", e.Name)
	}
	if e.Type != "" {
		o = o.Set("type:s", e.Type)
	}
	if e.Hint != "" {
		o = o.Set("hint:s", e.Hint)
	}
	o = o.Set("isRepeated:b", e.IsRepeated)
	o = o.Set("isOptional:b", e.IsOptional)
	if len(e.Properties) > 0 {
		v := List{}
		for _, iv := range e.Properties {
			v = v.Append(iv.ToObject().Raw())
		}
		o = o.Set("properties:am", v)
	}
	return o
}

func (e *SchemaProperty) FromObject(o Object) error {
	data, ok := o.Raw().Value("data:m").(Map)
	if !ok {
		return errors.New("missing data")
	}
	e.raw = Object{}
	e.raw = e.raw.SetType(o.GetType())
	e.Metadata.Stream = o.GetStream()
	e.Metadata.Parents = o.GetParents()
	e.Metadata.Owners = o.GetOwners()
	e.Metadata.Signatures = o.GetSignatures()
	e.Metadata.Policy = o.GetPolicy()
	if v := data.Value("name:s"); v != nil {
		e.Name = string(v.PrimitiveHinted().(string))
	}
	if v := data.Value("type:s"); v != nil {
		e.Type = string(v.PrimitiveHinted().(string))
	}
	if v := data.Value("hint:s"); v != nil {
		e.Hint = string(v.PrimitiveHinted().(string))
	}
	if v := data.Value("isRepeated:b"); v != nil {
		e.IsRepeated = bool(v.PrimitiveHinted().(bool))
	}
	if v := data.Value("isOptional:b"); v != nil {
		e.IsOptional = bool(v.PrimitiveHinted().(bool))
	}
	if v := data.Value("properties:am"); v != nil && v.IsList() {
		m := v.PrimitiveHinted().([]interface{})
		e.Properties = make([]*SchemaProperty, len(m))
		for i, iv := range m {
			es := &SchemaProperty{}
			eo := FromMap(iv.(map[string]interface{}))
			es.FromObject(eo)
			e.Properties[i] = es
		}
	}
	return nil
}

func (e SchemaObject) GetType() string {
	return "nimona.io/SchemaObject"
}

func (e SchemaObject) IsStreamRoot() bool {
	return false
}

func (e SchemaObject) ToObject() Object {
	o := Object{}
	o = o.SetType("nimona.io/SchemaObject")
	if len(e.Metadata.Stream) > 0 {
		o = o.SetStream(e.Metadata.Stream)
	}
	if len(e.Metadata.Parents) > 0 {
		o = o.SetParents(e.Metadata.Parents)
	}
	if len(e.Metadata.Owners) > 0 {
		o = o.SetOwners(e.Metadata.Owners)
	}
	o = o.AddSignature(e.Metadata.Signatures...)
	o = o.SetPolicy(e.Metadata.Policy)
	if len(e.Properties) > 0 {
		v := List{}
		for _, iv := range e.Properties {
			v = v.Append(iv.ToObject().Raw())
		}
		o = o.Set("properties:am", v)
	}
	return o
}

func (e *SchemaObject) FromObject(o Object) error {
	data, ok := o.Raw().Value("data:m").(Map)
	if !ok {
		return errors.New("missing data")
	}
	e.raw = Object{}
	e.raw = e.raw.SetType(o.GetType())
	e.Metadata.Stream = o.GetStream()
	e.Metadata.Parents = o.GetParents()
	e.Metadata.Owners = o.GetOwners()
	e.Metadata.Signatures = o.GetSignatures()
	e.Metadata.Policy = o.GetPolicy()
	if v := data.Value("properties:am"); v != nil && v.IsList() {
		m := v.PrimitiveHinted().([]interface{})
		e.Properties = make([]*SchemaProperty, len(m))
		for i, iv := range m {
			es := &SchemaProperty{}
			eo := FromMap(iv.(map[string]interface{}))
			es.FromObject(eo)
			e.Properties[i] = es
		}
	}
	return nil
}
