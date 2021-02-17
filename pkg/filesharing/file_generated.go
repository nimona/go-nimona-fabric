// Code generated by nimona.io/tools/codegen. DO NOT EDIT.

package filesharing

import (
	object "nimona.io/pkg/object"
)

type (
	File struct {
		Metadata object.Metadata `nimona:"metadata:m,omitempty"`
		Name     string
		Chunks   []object.Hash
	}
	TransferRequest struct {
		Metadata object.Metadata `nimona:"metadata:m,omitempty"`
		File     *File
		Nonce    string
	}
	TransferResponse struct {
		Metadata object.Metadata `nimona:"metadata:m,omitempty"`
		Nonce    string
		Accepted bool
	}
)

func (e *File) Type() string {
	return "nimona.io/File"
}

func (e File) ToObject() *object.Object {
	r := &object.Object{
		Type:     "nimona.io/File",
		Metadata: e.Metadata,
		Data:     object.Map{},
	}
	// else
	// r.Data["name"] = object.String(e.Name)
	r.Data["name"] = object.String(e.Name)
	// if $member.IsRepeated
	if len(e.Chunks) > 0 {
		// else
		// r.Data["chunks"] = object.ToStringArray(e.Chunks)
		rv := make(object.StringArray, len(e.Chunks))
		for i, iv := range e.Chunks {
			rv[i] = object.String(iv)
		}
		r.Data["chunks"] = rv
	}
	return r
}

func (e *File) FromObject(o *object.Object) error {
	e.Metadata = o.Metadata
	if v, ok := o.Data["name"]; ok {
		if t, ok := v.(object.String); ok {
			e.Name = string(t)
		}
	}
	if v, ok := o.Data["chunks"]; ok {
		if t, ok := v.(object.StringArray); ok {
			rv := make([]object.Hash, len(t))
			for i, iv := range t {
				rv[i] = object.Hash(iv)
			}
			e.Chunks = rv
		}
	}
	return nil
}

func (e *TransferRequest) Type() string {
	return "nimona.io/TransferRequest"
}

func (e TransferRequest) ToObject() *object.Object {
	r := &object.Object{
		Type:     "nimona.io/TransferRequest",
		Metadata: e.Metadata,
		Data:     object.Map{},
	}
	// else if $member.IsObject
	if e.File != nil {
		r.Data["file"] = e.File.ToObject()
	}
	// else
	// r.Data["nonce"] = object.String(e.Nonce)
	r.Data["nonce"] = object.String(e.Nonce)
	return r
}

func (e *TransferRequest) FromObject(o *object.Object) error {
	e.Metadata = o.Metadata
	if v, ok := o.Data["file"]; ok {
		if t, ok := v.(object.Map); ok {
			es := &File{}
			eo := object.FromMap(t)
			es.FromObject(eo)
			e.File = es
		} else if t, ok := v.(*object.Object); ok {
			es := &File{}
			es.FromObject(t)
			e.File = es
		}
	}
	if v, ok := o.Data["nonce"]; ok {
		if t, ok := v.(object.String); ok {
			e.Nonce = string(t)
		}
	}
	return nil
}

func (e *TransferResponse) Type() string {
	return "nimona.io/TransferResponse"
}

func (e TransferResponse) ToObject() *object.Object {
	r := &object.Object{
		Type:     "nimona.io/TransferResponse",
		Metadata: e.Metadata,
		Data:     object.Map{},
	}
	// else
	// r.Data["nonce"] = object.String(e.Nonce)
	r.Data["nonce"] = object.String(e.Nonce)
	// else
	// r.Data["accepted"] = object.Bool(e.Accepted)
	r.Data["accepted"] = object.Bool(e.Accepted)
	return r
}

func (e *TransferResponse) FromObject(o *object.Object) error {
	e.Metadata = o.Metadata
	if v, ok := o.Data["nonce"]; ok {
		if t, ok := v.(object.String); ok {
			e.Nonce = string(t)
		}
	}
	if v, ok := o.Data["accepted"]; ok {
		if t, ok := v.(object.Bool); ok {
			e.Accepted = bool(t)
		}
	}
	return nil
}
