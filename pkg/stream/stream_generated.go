// Code generated by nimona.io/tools/codegen. DO NOT EDIT.

package stream

import (
	object "nimona.io/pkg/object"
	value "nimona.io/pkg/object/value"
)

type (
	Policy struct {
		Metadata   object.Metadata `nimona:"@metadata:m"`
		Subjects   []string        `nimona:"subjects:as"`
		Resources  []string        `nimona:"resources:as"`
		Conditions []string        `nimona:"conditions:as"`
		Action     string          `nimona:"action:s"`
	}
	Request struct {
		Metadata  object.Metadata `nimona:"@metadata:m"`
		RequestID string          `nimona:"requestID:s"`
		RootCID   value.CID       `nimona:"rootCID:s"`
	}
	Response struct {
		Metadata  object.Metadata `nimona:"@metadata:m"`
		RequestID string          `nimona:"requestID:s"`
		RootCID   value.CID       `nimona:"rootCID:s"`
		Leaves    []value.CID     `nimona:"leaves:as"`
	}
	Announcement struct {
		Metadata   object.Metadata `nimona:"@metadata:m"`
		StreamCID  value.CID       `nimona:"streamCID:s"`
		ObjectCIDs []value.CID     `nimona:"objectCIDs:as"`
	}
	Subscription struct {
		Metadata object.Metadata `nimona:"@metadata:m"`
		RootCIDs []value.CID     `nimona:"rootCIDs:as"`
		Expiry   string          `nimona:"expiry:s"`
	}
)

func (e *Policy) Type() string {
	return "nimona.io/stream.Policy"
}

func (e *Policy) MarshalObject() (*object.Object, error) {
	o, err := object.Marshal(e)
	if err != nil {
		return nil, err
	}
	o.Type = "nimona.io/stream.Policy"
	return o, nil
}

func (e *Policy) UnmarshalObject(o *object.Object) error {
	return object.Unmarshal(o, e)
}

func (e *Request) Type() string {
	return "nimona.io/stream.Request"
}

func (e *Request) MarshalObject() (*object.Object, error) {
	o, err := object.Marshal(e)
	if err != nil {
		return nil, err
	}
	o.Type = "nimona.io/stream.Request"
	return o, nil
}

func (e *Request) UnmarshalObject(o *object.Object) error {
	return object.Unmarshal(o, e)
}

func (e *Response) Type() string {
	return "nimona.io/stream.Response"
}

func (e *Response) MarshalObject() (*object.Object, error) {
	o, err := object.Marshal(e)
	if err != nil {
		return nil, err
	}
	o.Type = "nimona.io/stream.Response"
	return o, nil
}

func (e *Response) UnmarshalObject(o *object.Object) error {
	return object.Unmarshal(o, e)
}

func (e *Announcement) Type() string {
	return "nimona.io/stream.Announcement"
}

func (e *Announcement) MarshalObject() (*object.Object, error) {
	o, err := object.Marshal(e)
	if err != nil {
		return nil, err
	}
	o.Type = "nimona.io/stream.Announcement"
	return o, nil
}

func (e *Announcement) UnmarshalObject(o *object.Object) error {
	return object.Unmarshal(o, e)
}

func (e *Subscription) Type() string {
	return "nimona.io/stream.Subscription"
}

func (e *Subscription) MarshalObject() (*object.Object, error) {
	o, err := object.Marshal(e)
	if err != nil {
		return nil, err
	}
	o.Type = "nimona.io/stream.Subscription"
	return o, nil
}

func (e *Subscription) UnmarshalObject(o *object.Object) error {
	return object.Unmarshal(o, e)
}
