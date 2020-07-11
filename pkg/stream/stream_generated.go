// Code generated by nimona.io/tools/codegen. DO NOT EDIT.

package stream

import (
	"errors"

	crypto "nimona.io/pkg/crypto"
	object "nimona.io/pkg/object"
)

type (
	Policy struct {
		raw        object.Object
		Stream     object.Hash
		Parents    []object.Hash
		Owners     []crypto.PublicKey
		Policy     object.Policy
		Signatures []object.Signature
		Subjects   []string
		Resources  []string
		Conditions []string
		Action     string
	}
	Request struct {
		raw        object.Object
		Stream     object.Hash
		Parents    []object.Hash
		Owners     []crypto.PublicKey
		Policy     object.Policy
		Signatures []object.Signature
		Nonce      string
		Leaves     []object.Hash
	}
	Response struct {
		raw        object.Object
		Stream     object.Hash
		Parents    []object.Hash
		Owners     []crypto.PublicKey
		Policy     object.Policy
		Signatures []object.Signature
		Nonce      string
		Children   []object.Hash
	}
	Announcement struct {
		raw        object.Object
		Stream     object.Hash
		Parents    []object.Hash
		Owners     []crypto.PublicKey
		Policy     object.Policy
		Signatures []object.Signature
		Nonce      string
		Objects    []*object.Object
	}
)

func (e Policy) GetType() string {
	return "nimona.io/stream.Policy"
}

func (e Policy) GetSchema() *object.SchemaObject {
	return &object.SchemaObject{
		Properties: []*object.SchemaProperty{
			&object.SchemaProperty{
				Name:       "subjects",
				Type:       "string",
				Hint:       "s",
				IsRepeated: true,
				IsOptional: false,
			},
			&object.SchemaProperty{
				Name:       "resources",
				Type:       "string",
				Hint:       "s",
				IsRepeated: true,
				IsOptional: false,
			},
			&object.SchemaProperty{
				Name:       "conditions",
				Type:       "string",
				Hint:       "s",
				IsRepeated: true,
				IsOptional: false,
			},
			&object.SchemaProperty{
				Name:       "action",
				Type:       "string",
				Hint:       "s",
				IsRepeated: false,
				IsOptional: false,
			},
		},
	}
}

func (e Policy) ToObject() object.Object {
	o := object.Object{}
	o = o.SetType("nimona.io/stream.Policy")
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
	if len(e.Subjects) > 0 {
		v := object.List{}
		for _, iv := range e.Subjects {
			v = v.Append(object.String(iv))
		}
		o = o.Set("subjects:as", v)
	}
	if len(e.Resources) > 0 {
		v := object.List{}
		for _, iv := range e.Resources {
			v = v.Append(object.String(iv))
		}
		o = o.Set("resources:as", v)
	}
	if len(e.Conditions) > 0 {
		v := object.List{}
		for _, iv := range e.Conditions {
			v = v.Append(object.String(iv))
		}
		o = o.Set("conditions:as", v)
	}
	if e.Action != "" {
		o = o.Set("action:s", e.Action)
	}
	// if schema := e.GetSchema(); schema != nil {
	// 	m["_schema:m"] = schema.ToObject().ToMap()
	// }
	return o
}

func (e *Policy) FromObject(o object.Object) error {
	data, ok := o.Raw().Value("data:m").(object.Map)
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
	if v := data.Value("subjects:as"); v != nil && v.IsList() {
		m := v.PrimitiveHinted().([]string)
		e.Subjects = make([]string, len(m))
		for i, iv := range m {
			e.Subjects[i] = string(iv)
		}
	}
	if v := data.Value("resources:as"); v != nil && v.IsList() {
		m := v.PrimitiveHinted().([]string)
		e.Resources = make([]string, len(m))
		for i, iv := range m {
			e.Resources[i] = string(iv)
		}
	}
	if v := data.Value("conditions:as"); v != nil && v.IsList() {
		m := v.PrimitiveHinted().([]string)
		e.Conditions = make([]string, len(m))
		for i, iv := range m {
			e.Conditions[i] = string(iv)
		}
	}
	if v := data.Value("action:s"); v != nil {
		e.Action = string(v.PrimitiveHinted().(string))
	}
	return nil
}

func (e Request) GetType() string {
	return "nimona.io/stream.Request"
}

func (e Request) GetSchema() *object.SchemaObject {
	return &object.SchemaObject{
		Properties: []*object.SchemaProperty{
			&object.SchemaProperty{
				Name:       "nonce",
				Type:       "string",
				Hint:       "s",
				IsRepeated: false,
				IsOptional: false,
			},
			&object.SchemaProperty{
				Name:       "leaves",
				Type:       "nimona.io/object.Hash",
				Hint:       "s",
				IsRepeated: true,
				IsOptional: false,
			},
		},
	}
}

func (e Request) ToObject() object.Object {
	o := object.Object{}
	o = o.SetType("nimona.io/stream.Request")
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
	if e.Nonce != "" {
		o = o.Set("nonce:s", e.Nonce)
	}
	if len(e.Leaves) > 0 {
		v := object.List{}
		for _, iv := range e.Leaves {
			v = v.Append(object.String(iv))
		}
		o = o.Set("leaves:as", v)
	}
	// if schema := e.GetSchema(); schema != nil {
	// 	m["_schema:m"] = schema.ToObject().ToMap()
	// }
	return o
}

func (e *Request) FromObject(o object.Object) error {
	data, ok := o.Raw().Value("data:m").(object.Map)
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
	if v := data.Value("nonce:s"); v != nil {
		e.Nonce = string(v.PrimitiveHinted().(string))
	}
	if v := data.Value("leaves:as"); v != nil && v.IsList() {
		m := v.PrimitiveHinted().([]string)
		e.Leaves = make([]object.Hash, len(m))
		for i, iv := range m {
			e.Leaves[i] = object.Hash(iv)
		}
	}
	return nil
}

func (e Response) GetType() string {
	return "nimona.io/stream.Response"
}

func (e Response) GetSchema() *object.SchemaObject {
	return &object.SchemaObject{
		Properties: []*object.SchemaProperty{
			&object.SchemaProperty{
				Name:       "nonce",
				Type:       "string",
				Hint:       "s",
				IsRepeated: false,
				IsOptional: false,
			},
			&object.SchemaProperty{
				Name:       "children",
				Type:       "nimona.io/object.Hash",
				Hint:       "s",
				IsRepeated: true,
				IsOptional: false,
			},
		},
	}
}

func (e Response) ToObject() object.Object {
	o := object.Object{}
	o = o.SetType("nimona.io/stream.Response")
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
	if e.Nonce != "" {
		o = o.Set("nonce:s", e.Nonce)
	}
	if len(e.Children) > 0 {
		v := object.List{}
		for _, iv := range e.Children {
			v = v.Append(object.String(iv))
		}
		o = o.Set("children:as", v)
	}
	// if schema := e.GetSchema(); schema != nil {
	// 	m["_schema:m"] = schema.ToObject().ToMap()
	// }
	return o
}

func (e *Response) FromObject(o object.Object) error {
	data, ok := o.Raw().Value("data:m").(object.Map)
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
	if v := data.Value("nonce:s"); v != nil {
		e.Nonce = string(v.PrimitiveHinted().(string))
	}
	if v := data.Value("children:as"); v != nil && v.IsList() {
		m := v.PrimitiveHinted().([]string)
		e.Children = make([]object.Hash, len(m))
		for i, iv := range m {
			e.Children[i] = object.Hash(iv)
		}
	}
	return nil
}

func (e Announcement) GetType() string {
	return "nimona.io/stream.Announcement"
}

func (e Announcement) GetSchema() *object.SchemaObject {
	return &object.SchemaObject{
		Properties: []*object.SchemaProperty{
			&object.SchemaProperty{
				Name:       "nonce",
				Type:       "string",
				Hint:       "s",
				IsRepeated: false,
				IsOptional: false,
			},
			&object.SchemaProperty{
				Name:       "objects",
				Type:       "nimona.io/object.Object",
				Hint:       "m",
				IsRepeated: true,
				IsOptional: false,
			},
		},
	}
}

func (e Announcement) ToObject() object.Object {
	o := object.Object{}
	o = o.SetType("nimona.io/stream.Announcement")
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
	if e.Nonce != "" {
		o = o.Set("nonce:s", e.Nonce)
	}
	if len(e.Objects) > 0 {
		v := object.List{}
		for _, iv := range e.Objects {
			v = v.Append(iv.ToObject().Raw())
		}
		o = o.Set("objects:am", v)
	}
	// if schema := e.GetSchema(); schema != nil {
	// 	m["_schema:m"] = schema.ToObject().ToMap()
	// }
	return o
}

func (e *Announcement) FromObject(o object.Object) error {
	data, ok := o.Raw().Value("data:m").(object.Map)
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
	if v := data.Value("nonce:s"); v != nil {
		e.Nonce = string(v.PrimitiveHinted().(string))
	}
	if v := data.Value("objects:am"); v != nil && v.IsList() {
		m := v.PrimitiveHinted().([]interface{})
		e.Objects = make([]*object.Object, len(m))
		for i, iv := range m {
			eo := object.FromMap(iv.(map[string]interface{}))
			e.Objects[i] = &eo
		}
	}
	return nil
}
