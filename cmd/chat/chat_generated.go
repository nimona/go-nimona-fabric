// Code generated by nimona.io/tools/codegen. DO NOT EDIT.

package main

import (
	"errors"

	object "nimona.io/pkg/object"
)

type (
	ConversationStreamRoot struct {
		raw      object.Object
		Metadata object.Metadata
		Nonce    string
	}
	ConversationSetNickname struct {
		raw      object.Object
		Metadata object.Metadata
		Datetime string
		Nickname string
	}
	ConversationMessageAdded struct {
		raw      object.Object
		Metadata object.Metadata
		Datetime string
		Body     string
	}
)

func (e ConversationStreamRoot) GetType() string {
	return "stream:poc.nimona.io/conversation"
}

func (e ConversationStreamRoot) IsStreamRoot() bool {
	return true
}

func (e ConversationStreamRoot) GetSchema() *object.SchemaObject {
	return &object.SchemaObject{
		Properties: []*object.SchemaProperty{{
			Name:       "nonce",
			Type:       "string",
			Hint:       "s",
			IsRepeated: false,
			IsOptional: false,
		}},
	}
}

func (e ConversationStreamRoot) ToObject() object.Object {
	o := object.Object{}
	o = o.SetType("stream:poc.nimona.io/conversation")
	if len(e.Metadata.Stream) > 0 {
		o = o.SetStream(e.Metadata.Stream)
	}
	if len(e.Metadata.Parents) > 0 {
		o = o.SetParents(e.Metadata.Parents)
	}
	if !e.Metadata.Owner.IsEmpty() {
		o = o.SetOwner(e.Metadata.Owner)
	}
	if !e.Metadata.Signature.IsEmpty() {
		o = o.SetSignature(e.Metadata.Signature)
	}
	o = o.SetPolicy(e.Metadata.Policy)
	if e.Nonce != "" {
		o = o.Set("nonce:s", e.Nonce)
	}
	// if schema := e.GetSchema(); schema != nil {
	// 	m["_schema:m"] = schema.ToObject().ToMap()
	// }
	return o
}

func (e *ConversationStreamRoot) FromObject(o object.Object) error {
	data, ok := o.Raw().Value("data:m").(object.Map)
	if !ok {
		return errors.New("missing data")
	}
	e.raw = object.Object{}
	e.raw = e.raw.SetType(o.GetType())
	e.Metadata.Stream = o.GetStream()
	e.Metadata.Parents = o.GetParents()
	e.Metadata.Owner = o.GetOwner()
	e.Metadata.Signature = o.GetSignature()
	e.Metadata.Policy = o.GetPolicy()
	if v := data.Value("nonce:s"); v != nil {
		e.Nonce = string(v.PrimitiveHinted().(string))
	}
	return nil
}

func (e ConversationSetNickname) GetType() string {
	return "poc.nimona.io/conversation.SetNickname"
}

func (e ConversationSetNickname) IsStreamRoot() bool {
	return false
}

func (e ConversationSetNickname) GetSchema() *object.SchemaObject {
	return &object.SchemaObject{
		Properties: []*object.SchemaProperty{{
			Name:       "datetime",
			Type:       "string",
			Hint:       "s",
			IsRepeated: false,
			IsOptional: false,
		}, {
			Name:       "nickname",
			Type:       "string",
			Hint:       "s",
			IsRepeated: false,
			IsOptional: false,
		}},
	}
}

func (e ConversationSetNickname) ToObject() object.Object {
	o := object.Object{}
	o = o.SetType("poc.nimona.io/conversation.SetNickname")
	if len(e.Metadata.Stream) > 0 {
		o = o.SetStream(e.Metadata.Stream)
	}
	if len(e.Metadata.Parents) > 0 {
		o = o.SetParents(e.Metadata.Parents)
	}
	if !e.Metadata.Owner.IsEmpty() {
		o = o.SetOwner(e.Metadata.Owner)
	}
	if !e.Metadata.Signature.IsEmpty() {
		o = o.SetSignature(e.Metadata.Signature)
	}
	o = o.SetPolicy(e.Metadata.Policy)
	if e.Datetime != "" {
		o = o.Set("datetime:s", e.Datetime)
	}
	if e.Nickname != "" {
		o = o.Set("nickname:s", e.Nickname)
	}
	// if schema := e.GetSchema(); schema != nil {
	// 	m["_schema:m"] = schema.ToObject().ToMap()
	// }
	return o
}

func (e *ConversationSetNickname) FromObject(o object.Object) error {
	data, ok := o.Raw().Value("data:m").(object.Map)
	if !ok {
		return errors.New("missing data")
	}
	e.raw = object.Object{}
	e.raw = e.raw.SetType(o.GetType())
	e.Metadata.Stream = o.GetStream()
	e.Metadata.Parents = o.GetParents()
	e.Metadata.Owner = o.GetOwner()
	e.Metadata.Signature = o.GetSignature()
	e.Metadata.Policy = o.GetPolicy()
	if v := data.Value("datetime:s"); v != nil {
		e.Datetime = string(v.PrimitiveHinted().(string))
	}
	if v := data.Value("nickname:s"); v != nil {
		e.Nickname = string(v.PrimitiveHinted().(string))
	}
	return nil
}

func (e ConversationMessageAdded) GetType() string {
	return "poc.nimona.io/conversation.MessageAdded"
}

func (e ConversationMessageAdded) IsStreamRoot() bool {
	return false
}

func (e ConversationMessageAdded) GetSchema() *object.SchemaObject {
	return &object.SchemaObject{
		Properties: []*object.SchemaProperty{{
			Name:       "datetime",
			Type:       "string",
			Hint:       "s",
			IsRepeated: false,
			IsOptional: false,
		}, {
			Name:       "body",
			Type:       "string",
			Hint:       "s",
			IsRepeated: false,
			IsOptional: false,
		}},
	}
}

func (e ConversationMessageAdded) ToObject() object.Object {
	o := object.Object{}
	o = o.SetType("poc.nimona.io/conversation.MessageAdded")
	if len(e.Metadata.Stream) > 0 {
		o = o.SetStream(e.Metadata.Stream)
	}
	if len(e.Metadata.Parents) > 0 {
		o = o.SetParents(e.Metadata.Parents)
	}
	if !e.Metadata.Owner.IsEmpty() {
		o = o.SetOwner(e.Metadata.Owner)
	}
	if !e.Metadata.Signature.IsEmpty() {
		o = o.SetSignature(e.Metadata.Signature)
	}
	o = o.SetPolicy(e.Metadata.Policy)
	if e.Datetime != "" {
		o = o.Set("datetime:s", e.Datetime)
	}
	if e.Body != "" {
		o = o.Set("body:s", e.Body)
	}
	// if schema := e.GetSchema(); schema != nil {
	// 	m["_schema:m"] = schema.ToObject().ToMap()
	// }
	return o
}

func (e *ConversationMessageAdded) FromObject(o object.Object) error {
	data, ok := o.Raw().Value("data:m").(object.Map)
	if !ok {
		return errors.New("missing data")
	}
	e.raw = object.Object{}
	e.raw = e.raw.SetType(o.GetType())
	e.Metadata.Stream = o.GetStream()
	e.Metadata.Parents = o.GetParents()
	e.Metadata.Owner = o.GetOwner()
	e.Metadata.Signature = o.GetSignature()
	e.Metadata.Policy = o.GetPolicy()
	if v := data.Value("datetime:s"); v != nil {
		e.Datetime = string(v.PrimitiveHinted().(string))
	}
	if v := data.Value("body:s"); v != nil {
		e.Body = string(v.PrimitiveHinted().(string))
	}
	return nil
}
