// Code generated by nimona.io/tools/codegen. DO NOT EDIT.

package main

import (
	object "nimona.io/pkg/object"
)

type (
	ConversationStreamRoot struct {
		Metadata object.Metadata `nimona:"metadata:m,omitempty"`
		Nonce    string          `nimona:"nonce:s,omitempty"`
	}
	ConversationNicknameUpdated struct {
		Metadata object.Metadata `nimona:"metadata:m,omitempty"`
		Nickname string          `nimona:"nickname:s,omitempty"`
	}
	ConversationMessageAdded struct {
		Metadata object.Metadata `nimona:"metadata:m,omitempty"`
		Body     string          `nimona:"body:s,omitempty"`
	}
)

func (e *ConversationStreamRoot) Type() string {
	return "stream:poc.nimona.io/conversation"
}

func (e ConversationStreamRoot) ToObject() *object.Object {
	r := &object.Object{
		Type:     "stream:poc.nimona.io/conversation",
		Metadata: e.Metadata,
		Data:     map[string]interface{}{},
	}
	r.Data["nonce:s"] = e.Nonce
	return r
}

func (e ConversationStreamRoot) ToObjectMap() map[string]interface{} {
	d := map[string]interface{}{}
	d["nonce:s"] = e.Nonce
	r := map[string]interface{}{
		"type:s":     "stream:poc.nimona.io/conversation",
		"metadata:m": object.MetadataToMap(&e.Metadata),
		"data:m":     d,
	}
	return r
}

func (e *ConversationStreamRoot) FromObject(o *object.Object) error {
	return object.Decode(o, e)
}

func (e *ConversationNicknameUpdated) Type() string {
	return "poc.nimona.io/conversation.NicknameUpdated"
}

func (e ConversationNicknameUpdated) ToObject() *object.Object {
	r := &object.Object{
		Type:     "poc.nimona.io/conversation.NicknameUpdated",
		Metadata: e.Metadata,
		Data:     map[string]interface{}{},
	}
	r.Data["nickname:s"] = e.Nickname
	return r
}

func (e ConversationNicknameUpdated) ToObjectMap() map[string]interface{} {
	d := map[string]interface{}{}
	d["nickname:s"] = e.Nickname
	r := map[string]interface{}{
		"type:s":     "poc.nimona.io/conversation.NicknameUpdated",
		"metadata:m": object.MetadataToMap(&e.Metadata),
		"data:m":     d,
	}
	return r
}

func (e *ConversationNicknameUpdated) FromObject(o *object.Object) error {
	return object.Decode(o, e)
}

func (e *ConversationMessageAdded) Type() string {
	return "poc.nimona.io/conversation.MessageAdded"
}

func (e ConversationMessageAdded) ToObject() *object.Object {
	r := &object.Object{
		Type:     "poc.nimona.io/conversation.MessageAdded",
		Metadata: e.Metadata,
		Data:     map[string]interface{}{},
	}
	r.Data["body:s"] = e.Body
	return r
}

func (e ConversationMessageAdded) ToObjectMap() map[string]interface{} {
	d := map[string]interface{}{}
	d["body:s"] = e.Body
	r := map[string]interface{}{
		"type:s":     "poc.nimona.io/conversation.MessageAdded",
		"metadata:m": object.MetadataToMap(&e.Metadata),
		"data:m":     d,
	}
	return r
}

func (e *ConversationMessageAdded) FromObject(o *object.Object) error {
	return object.Decode(o, e)
}
