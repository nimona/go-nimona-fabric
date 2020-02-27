// Code generated by nimona.io/tools/codegen. DO NOT EDIT.

package crypto

import (
	"errors"

	crypto "nimona.io/pkg/crypto"
	immutable "nimona.io/pkg/immutable"
	object "nimona.io/pkg/object"
)

type (
	Hash struct {
		raw        object.Object
		Stream     object.Hash
		Parents    []object.Hash
		Owners     []crypto.PublicKey
		Policy     object.Policy
		Signatures []object.Signature
		HashType   string
		Digest     []byte
	}
	HeaderSignature struct {
		raw        object.Object
		Stream     object.Hash
		Parents    []object.Hash
		Owners     []crypto.PublicKey
		Policy     object.Policy
		Signatures []object.Signature
		PublicKey  *PublicKey
		Algorithm  string
		R          []byte
		S          []byte
	}
	PrivateKey struct {
		raw        object.Object
		Stream     object.Hash
		Parents    []object.Hash
		Owners     []crypto.PublicKey
		Policy     object.Policy
		Signatures []object.Signature
		PublicKey  *PublicKey
		KeyType    string
		Algorithm  string
		Curve      string
		X          []byte
		Y          []byte
		D          []byte
	}
	PublicKey struct {
		raw        object.Object
		Stream     object.Hash
		Parents    []object.Hash
		Owners     []crypto.PublicKey
		Policy     object.Policy
		Signatures []object.Signature
		KeyType    string
		Algorithm  string
		Curve      string
		X          []byte
		Y          []byte
	}
)

func (e Hash) GetType() string {
	return "example/crypto.Hash"
}

func (e Hash) GetSchema() *object.SchemaObject {
	return &object.SchemaObject{
		Properties: []*object.SchemaProperty{
			&object.SchemaProperty{
				Name:       "hashType",
				Type:       "string",
				Hint:       "s",
				IsRepeated: false,
				IsOptional: false,
			},
			&object.SchemaProperty{
				Name:       "digest",
				Type:       "data",
				Hint:       "d",
				IsRepeated: false,
				IsOptional: false,
			},
		},
	}
}

func (e Hash) ToObject() object.Object {
	o := object.Object{}
	o = o.SetType("example/crypto.Hash")
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
	if e.HashType != "" {
		o = o.Set("hashType:s", e.HashType)
	}
	if len(e.Digest) != 0 {
		o = o.Set("digest:d", e.Digest)
	}
	// if schema := e.GetSchema(); schema != nil {
	// 	m["_schema:o"] = schema.ToObject().ToMap()
	// }
	return o
}

func (e *Hash) FromObject(o object.Object) error {
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
	if v := data.Value("hashType:s"); v != nil {
		e.HashType = string(v.PrimitiveHinted().(string))
	}
	if v := data.Value("digest:d"); v != nil {
		e.Digest = []byte(v.PrimitiveHinted().([]byte))
	}
	return nil
}

func (e HeaderSignature) GetType() string {
	return "example/object.Header.Signature"
}

func (e HeaderSignature) GetSchema() *object.SchemaObject {
	return &object.SchemaObject{
		Properties: []*object.SchemaProperty{
			&object.SchemaProperty{
				Name:       "publicKey",
				Type:       "PublicKey",
				Hint:       "o",
				IsRepeated: false,
				IsOptional: false,
			},
			&object.SchemaProperty{
				Name:       "algorithm",
				Type:       "string",
				Hint:       "s",
				IsRepeated: false,
				IsOptional: false,
			},
			&object.SchemaProperty{
				Name:       "r",
				Type:       "data",
				Hint:       "d",
				IsRepeated: false,
				IsOptional: false,
			},
			&object.SchemaProperty{
				Name:       "s",
				Type:       "data",
				Hint:       "d",
				IsRepeated: false,
				IsOptional: false,
			},
		},
	}
}

func (e HeaderSignature) ToObject() object.Object {
	o := object.Object{}
	o = o.SetType("example/object.Header.Signature")
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
	if e.PublicKey != nil {
		o = o.Set("publicKey:o", e.PublicKey.ToObject().Raw())
	}
	if e.Algorithm != "" {
		o = o.Set("algorithm:s", e.Algorithm)
	}
	if len(e.R) != 0 {
		o = o.Set("r:d", e.R)
	}
	if len(e.S) != 0 {
		o = o.Set("s:d", e.S)
	}
	// if schema := e.GetSchema(); schema != nil {
	// 	m["_schema:o"] = schema.ToObject().ToMap()
	// }
	return o
}

func (e *HeaderSignature) FromObject(o object.Object) error {
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
	if v := data.Value("publicKey:o"); v != nil {
		es := &PublicKey{}
		eo := object.FromMap(v.PrimitiveHinted().(map[string]interface{}))
		es.FromObject(eo)
		e.PublicKey = es
	}
	if v := data.Value("algorithm:s"); v != nil {
		e.Algorithm = string(v.PrimitiveHinted().(string))
	}
	if v := data.Value("r:d"); v != nil {
		e.R = []byte(v.PrimitiveHinted().([]byte))
	}
	if v := data.Value("s:d"); v != nil {
		e.S = []byte(v.PrimitiveHinted().([]byte))
	}
	return nil
}

func (e PrivateKey) GetType() string {
	return "example/crypto.PrivateKey"
}

func (e PrivateKey) GetSchema() *object.SchemaObject {
	return &object.SchemaObject{
		Properties: []*object.SchemaProperty{
			&object.SchemaProperty{
				Name:       "publicKey",
				Type:       "PublicKey",
				Hint:       "o",
				IsRepeated: false,
				IsOptional: false,
			},
			&object.SchemaProperty{
				Name:       "keyType",
				Type:       "string",
				Hint:       "s",
				IsRepeated: false,
				IsOptional: false,
			},
			&object.SchemaProperty{
				Name:       "algorithm",
				Type:       "string",
				Hint:       "s",
				IsRepeated: false,
				IsOptional: false,
			},
			&object.SchemaProperty{
				Name:       "curve",
				Type:       "string",
				Hint:       "s",
				IsRepeated: false,
				IsOptional: false,
			},
			&object.SchemaProperty{
				Name:       "x",
				Type:       "data",
				Hint:       "d",
				IsRepeated: false,
				IsOptional: false,
			},
			&object.SchemaProperty{
				Name:       "y",
				Type:       "data",
				Hint:       "d",
				IsRepeated: false,
				IsOptional: false,
			},
			&object.SchemaProperty{
				Name:       "d",
				Type:       "data",
				Hint:       "d",
				IsRepeated: false,
				IsOptional: false,
			},
		},
	}
}

func (e PrivateKey) ToObject() object.Object {
	o := object.Object{}
	o = o.SetType("example/crypto.PrivateKey")
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
	if e.PublicKey != nil {
		o = o.Set("publicKey:o", e.PublicKey.ToObject().Raw())
	}
	if e.KeyType != "" {
		o = o.Set("keyType:s", e.KeyType)
	}
	if e.Algorithm != "" {
		o = o.Set("algorithm:s", e.Algorithm)
	}
	if e.Curve != "" {
		o = o.Set("curve:s", e.Curve)
	}
	if len(e.X) != 0 {
		o = o.Set("x:d", e.X)
	}
	if len(e.Y) != 0 {
		o = o.Set("y:d", e.Y)
	}
	if len(e.D) != 0 {
		o = o.Set("d:d", e.D)
	}
	// if schema := e.GetSchema(); schema != nil {
	// 	m["_schema:o"] = schema.ToObject().ToMap()
	// }
	return o
}

func (e *PrivateKey) FromObject(o object.Object) error {
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
	if v := data.Value("publicKey:o"); v != nil {
		es := &PublicKey{}
		eo := object.FromMap(v.PrimitiveHinted().(map[string]interface{}))
		es.FromObject(eo)
		e.PublicKey = es
	}
	if v := data.Value("keyType:s"); v != nil {
		e.KeyType = string(v.PrimitiveHinted().(string))
	}
	if v := data.Value("algorithm:s"); v != nil {
		e.Algorithm = string(v.PrimitiveHinted().(string))
	}
	if v := data.Value("curve:s"); v != nil {
		e.Curve = string(v.PrimitiveHinted().(string))
	}
	if v := data.Value("x:d"); v != nil {
		e.X = []byte(v.PrimitiveHinted().([]byte))
	}
	if v := data.Value("y:d"); v != nil {
		e.Y = []byte(v.PrimitiveHinted().([]byte))
	}
	if v := data.Value("d:d"); v != nil {
		e.D = []byte(v.PrimitiveHinted().([]byte))
	}
	return nil
}

func (e PublicKey) GetType() string {
	return "example/crypto.PublicKey"
}

func (e PublicKey) GetSchema() *object.SchemaObject {
	return &object.SchemaObject{
		Properties: []*object.SchemaProperty{
			&object.SchemaProperty{
				Name:       "keyType",
				Type:       "string",
				Hint:       "s",
				IsRepeated: false,
				IsOptional: false,
			},
			&object.SchemaProperty{
				Name:       "algorithm",
				Type:       "string",
				Hint:       "s",
				IsRepeated: false,
				IsOptional: false,
			},
			&object.SchemaProperty{
				Name:       "curve",
				Type:       "string",
				Hint:       "s",
				IsRepeated: false,
				IsOptional: false,
			},
			&object.SchemaProperty{
				Name:       "x",
				Type:       "data",
				Hint:       "d",
				IsRepeated: false,
				IsOptional: false,
			},
			&object.SchemaProperty{
				Name:       "y",
				Type:       "data",
				Hint:       "d",
				IsRepeated: false,
				IsOptional: false,
			},
		},
	}
}

func (e PublicKey) ToObject() object.Object {
	o := object.Object{}
	o = o.SetType("example/crypto.PublicKey")
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
	if e.KeyType != "" {
		o = o.Set("keyType:s", e.KeyType)
	}
	if e.Algorithm != "" {
		o = o.Set("algorithm:s", e.Algorithm)
	}
	if e.Curve != "" {
		o = o.Set("curve:s", e.Curve)
	}
	if len(e.X) != 0 {
		o = o.Set("x:d", e.X)
	}
	if len(e.Y) != 0 {
		o = o.Set("y:d", e.Y)
	}
	// if schema := e.GetSchema(); schema != nil {
	// 	m["_schema:o"] = schema.ToObject().ToMap()
	// }
	return o
}

func (e *PublicKey) FromObject(o object.Object) error {
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
	if v := data.Value("keyType:s"); v != nil {
		e.KeyType = string(v.PrimitiveHinted().(string))
	}
	if v := data.Value("algorithm:s"); v != nil {
		e.Algorithm = string(v.PrimitiveHinted().(string))
	}
	if v := data.Value("curve:s"); v != nil {
		e.Curve = string(v.PrimitiveHinted().(string))
	}
	if v := data.Value("x:d"); v != nil {
		e.X = []byte(v.PrimitiveHinted().([]byte))
	}
	if v := data.Value("y:d"); v != nil {
		e.Y = []byte(v.PrimitiveHinted().([]byte))
	}
	return nil
}
