// Code generated by nimona.io/tools/codegen. DO NOT EDIT.

package fixtures

import (
	object "nimona.io/pkg/object"
)

type (
	TestPolicy struct {
		Metadata   object.Metadata `nimona:"metadata:m"`
		Subjects   []string        `nimona:"subjects:as,omitempty"`
		Resources  []string        `nimona:"resources:as,omitempty"`
		Conditions []string        `nimona:"conditions:as,omitempty"`
		Action     string          `nimona:"action:s,omitempty"`
	}
	TestStream struct {
		Metadata        object.Metadata `nimona:"metadata:m"`
		Nonce           string          `nimona:"nonce:s,omitempty"`
		CreatedDateTime string          `nimona:"createdDateTime:s,omitempty"`
	}
	TestSubscribed struct {
		Metadata object.Metadata `nimona:"metadata:m"`
		Nonce    string          `nimona:"nonce:s,omitempty"`
	}
	TestUnsubscribed struct {
		Metadata object.Metadata `nimona:"metadata:m"`
		Nonce    string          `nimona:"nonce:s,omitempty"`
	}
)

func (e *TestPolicy) Type() string {
	return "nimona.io/fixtures.TestPolicy"
}

func (e TestPolicy) ToObject() *object.Object {
	o, err := object.Encode(&e)
	if err != nil {
		panic(err)
	}
	return o
}

func (e *TestPolicy) FromObject(o *object.Object) error {
	return object.Decode(o, e)
}

func (e *TestStream) Type() string {
	return "nimona.io/fixtures.TestStream"
}

func (e TestStream) ToObject() *object.Object {
	o, err := object.Encode(&e)
	if err != nil {
		panic(err)
	}
	return o
}

func (e *TestStream) FromObject(o *object.Object) error {
	return object.Decode(o, e)
}

func (e *TestSubscribed) Type() string {
	return "nimona.io/fixtures.TestSubscribed"
}

func (e TestSubscribed) ToObject() *object.Object {
	o, err := object.Encode(&e)
	if err != nil {
		panic(err)
	}
	return o
}

func (e *TestSubscribed) FromObject(o *object.Object) error {
	return object.Decode(o, e)
}

func (e *TestUnsubscribed) Type() string {
	return "nimona.io/fixtures.TestUnsubscribed"
}

func (e TestUnsubscribed) ToObject() *object.Object {
	o, err := object.Encode(&e)
	if err != nil {
		panic(err)
	}
	return o
}

func (e *TestUnsubscribed) FromObject(o *object.Object) error {
	return object.Decode(o, e)
}
