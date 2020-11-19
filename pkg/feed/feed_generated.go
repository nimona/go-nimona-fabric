// Code generated by nimona.io/tools/codegen. DO NOT EDIT.

package feed

import (
	object "nimona.io/pkg/object"
)

type (
	FeedStreamRoot struct {
		Metadata   object.Metadata `nimona:"metadata:m,omitempty"`
		ObjectType string          `nimona:"objectType:s,omitempty"`
		Datetime   string          `nimona:"datetime:s,omitempty"`
	}
	Added struct {
		Metadata   object.Metadata `nimona:"metadata:m,omitempty"`
		ObjectHash []object.Hash   `nimona:"objectHash:ar,omitempty"`
		Sequence   int64           `nimona:"sequence:i,omitempty"`
		Datetime   string          `nimona:"datetime:s,omitempty"`
	}
	Removed struct {
		Metadata   object.Metadata `nimona:"metadata:m,omitempty"`
		ObjectHash []object.Hash   `nimona:"objectHash:ar,omitempty"`
		Sequence   int64           `nimona:"sequence:i,omitempty"`
		Datetime   string          `nimona:"datetime:s,omitempty"`
	}
)

func (e *FeedStreamRoot) Type() string {
	return "stream:nimona.io/feed"
}

func (e FeedStreamRoot) ToObject() *object.Object {
	r := &object.Object{
		Type:     "stream:nimona.io/feed",
		Metadata: e.Metadata,
		Data:     map[string]interface{}{},
	}
	r.Data["objectType:s"] = e.ObjectType
	r.Data["datetime:s"] = e.Datetime
	return r
}

func (e FeedStreamRoot) ToObjectMap() map[string]interface{} {
	d := map[string]interface{}{}
	d["objectType:s"] = e.ObjectType
	d["datetime:s"] = e.Datetime
	r := map[string]interface{}{
		"type:s":     "stream:nimona.io/feed",
		"metadata:m": object.MetadataToMap(&e.Metadata),
		"data:m":     d,
	}
	return r
}

func (e *FeedStreamRoot) FromObject(o *object.Object) error {
	return object.Decode(o, e)
}

func (e *Added) Type() string {
	return "event:nimona.io/feed.Added"
}

func (e Added) ToObject() *object.Object {
	r := &object.Object{
		Type:     "event:nimona.io/feed.Added",
		Metadata: e.Metadata,
		Data:     map[string]interface{}{},
	}
	if len(e.ObjectHash) > 0 {
		r.Data["objectHash:ar"] = e.ObjectHash
	}
	r.Data["sequence:i"] = e.Sequence
	r.Data["datetime:s"] = e.Datetime
	return r
}

func (e Added) ToObjectMap() map[string]interface{} {
	d := map[string]interface{}{}
	if len(e.ObjectHash) > 0 {
		d["objectHash:ar"] = e.ObjectHash
	}
	d["sequence:i"] = e.Sequence
	d["datetime:s"] = e.Datetime
	r := map[string]interface{}{
		"type:s":     "event:nimona.io/feed.Added",
		"metadata:m": object.MetadataToMap(&e.Metadata),
		"data:m":     d,
	}
	return r
}

func (e *Added) FromObject(o *object.Object) error {
	return object.Decode(o, e)
}

func (e *Removed) Type() string {
	return "event:nimona.io/feed.Removed"
}

func (e Removed) ToObject() *object.Object {
	r := &object.Object{
		Type:     "event:nimona.io/feed.Removed",
		Metadata: e.Metadata,
		Data:     map[string]interface{}{},
	}
	if len(e.ObjectHash) > 0 {
		r.Data["objectHash:ar"] = e.ObjectHash
	}
	r.Data["sequence:i"] = e.Sequence
	r.Data["datetime:s"] = e.Datetime
	return r
}

func (e Removed) ToObjectMap() map[string]interface{} {
	d := map[string]interface{}{}
	if len(e.ObjectHash) > 0 {
		d["objectHash:ar"] = e.ObjectHash
	}
	d["sequence:i"] = e.Sequence
	d["datetime:s"] = e.Datetime
	r := map[string]interface{}{
		"type:s":     "event:nimona.io/feed.Removed",
		"metadata:m": object.MetadataToMap(&e.Metadata),
		"data:m":     d,
	}
	return r
}

func (e *Removed) FromObject(o *object.Object) error {
	return object.Decode(o, e)
}
