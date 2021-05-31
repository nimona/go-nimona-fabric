// Code generated by nimona.io/tools/codegen. DO NOT EDIT.

package fixtures

import (
	object "nimona.io/pkg/object"
)

type (
	CompositeTest struct {
		Metadata                    object.Metadata `nimona:"@metadata:m"`
		CompositeStringTest         *Composite      `nimona:"compositeStringTest:s"`
		CompositeDataTest           *Composite      `nimona:"compositeDataTest:d"`
		RepeatedCompositeStringTest []*Composite    `nimona:"repeatedCompositeStringTest:as"`
		RepeatedCompositeDataTest   []*Composite    `nimona:"repeatedCompositeDataTest:ad"`
	}
	TestPolicy struct {
		Metadata   object.Metadata `nimona:"@metadata:m"`
		Subjects   []string        `nimona:"subjects:as"`
		Resources  []string        `nimona:"resources:as"`
		Conditions []string        `nimona:"conditions:as"`
		Action     string          `nimona:"action:s"`
	}
	TestStream struct {
		Metadata        object.Metadata `nimona:"@metadata:m"`
		Nonce           string          `nimona:"nonce:s"`
		CreatedDateTime string          `nimona:"createdDateTime:s"`
	}
	TestSubscribed struct {
		Metadata object.Metadata `nimona:"@metadata:m"`
		Nonce    string          `nimona:"nonce:s"`
	}
	TestUnsubscribed struct {
		Metadata object.Metadata `nimona:"@metadata:m"`
		Nonce    string          `nimona:"nonce:s"`
	}
	TestRequest struct {
		Metadata  object.Metadata `nimona:"@metadata:m"`
		RequestID string          `nimona:"requestID:s"`
		Foo       string          `nimona:"foo:s"`
	}
	TestResponse struct {
		Metadata  object.Metadata `nimona:"@metadata:m"`
		RequestID string          `nimona:"requestID:s"`
		Foo       string          `nimona:"foo:s"`
	}
	Parent struct {
		Metadata      object.Metadata `nimona:"@metadata:m"`
		Foo           string          `nimona:"foo:s"`
		Child         *Child          `nimona:"child:o"`
		RepeatedChild []*Child        `nimona:"repeatedChild:ao"`
	}
	Child struct {
		Metadata object.Metadata `nimona:"@metadata:m"`
		Foo      string          `nimona:"foo:s"`
	}
)

func (e *CompositeTest) Type() string {
	return "compositeTest"
}

func (e *CompositeTest) MarshalObject() (*object.Object, error) {
	o, err := object.Marshal(e)
	if err != nil {
		return nil, err
	}
	o.Type = "compositeTest"
	return o, nil
}

func (e *CompositeTest) UnmarshalObject(o *object.Object) error {
	return object.Unmarshal(o, e)
}

func (e *TestPolicy) Type() string {
	return "nimona.io/fixtures.TestPolicy"
}

func (e *TestPolicy) MarshalObject() (*object.Object, error) {
	o, err := object.Marshal(e)
	if err != nil {
		return nil, err
	}
	o.Type = "nimona.io/fixtures.TestPolicy"
	return o, nil
}

func (e *TestPolicy) UnmarshalObject(o *object.Object) error {
	return object.Unmarshal(o, e)
}

func (e *TestStream) Type() string {
	return "nimona.io/fixtures.TestStream"
}

func (e *TestStream) MarshalObject() (*object.Object, error) {
	o, err := object.Marshal(e)
	if err != nil {
		return nil, err
	}
	o.Type = "nimona.io/fixtures.TestStream"
	return o, nil
}

func (e *TestStream) UnmarshalObject(o *object.Object) error {
	return object.Unmarshal(o, e)
}

func (e *TestSubscribed) Type() string {
	return "nimona.io/fixtures.TestSubscribed"
}

func (e *TestSubscribed) MarshalObject() (*object.Object, error) {
	o, err := object.Marshal(e)
	if err != nil {
		return nil, err
	}
	o.Type = "nimona.io/fixtures.TestSubscribed"
	return o, nil
}

func (e *TestSubscribed) UnmarshalObject(o *object.Object) error {
	return object.Unmarshal(o, e)
}

func (e *TestUnsubscribed) Type() string {
	return "nimona.io/fixtures.TestUnsubscribed"
}

func (e *TestUnsubscribed) MarshalObject() (*object.Object, error) {
	o, err := object.Marshal(e)
	if err != nil {
		return nil, err
	}
	o.Type = "nimona.io/fixtures.TestUnsubscribed"
	return o, nil
}

func (e *TestUnsubscribed) UnmarshalObject(o *object.Object) error {
	return object.Unmarshal(o, e)
}

func (e *TestRequest) Type() string {
	return "nimona.io/fixtures.TestRequest"
}

func (e *TestRequest) MarshalObject() (*object.Object, error) {
	o, err := object.Marshal(e)
	if err != nil {
		return nil, err
	}
	o.Type = "nimona.io/fixtures.TestRequest"
	return o, nil
}

func (e *TestRequest) UnmarshalObject(o *object.Object) error {
	return object.Unmarshal(o, e)
}

func (e *TestResponse) Type() string {
	return "nimona.io/fixtures.TestResponse"
}

func (e *TestResponse) MarshalObject() (*object.Object, error) {
	o, err := object.Marshal(e)
	if err != nil {
		return nil, err
	}
	o.Type = "nimona.io/fixtures.TestResponse"
	return o, nil
}

func (e *TestResponse) UnmarshalObject(o *object.Object) error {
	return object.Unmarshal(o, e)
}

func (e *Parent) Type() string {
	return "parent"
}

func (e *Parent) MarshalObject() (*object.Object, error) {
	o, err := object.Marshal(e)
	if err != nil {
		return nil, err
	}
	o.Type = "parent"
	return o, nil
}

func (e *Parent) UnmarshalObject(o *object.Object) error {
	return object.Unmarshal(o, e)
}

func (e *Child) Type() string {
	return "child"
}

func (e *Child) MarshalObject() (*object.Object, error) {
	o, err := object.Marshal(e)
	if err != nil {
		return nil, err
	}
	o.Type = "child"
	return o, nil
}

func (e *Child) UnmarshalObject(o *object.Object) error {
	return object.Unmarshal(o, e)
}
