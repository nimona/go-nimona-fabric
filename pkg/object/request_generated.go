// Code generated by nimona.io/tools/codegen. DO NOT EDIT.

package object

import (
	chore "nimona.io/pkg/chore"
)

type (
	Request struct {
		Metadata   Metadata   `nimona:"@metadata:m"`
		RequestID  string     `nimona:"requestID:s"`
		ObjectHash chore.Hash `nimona:"objectHash:s"`
	}
	Response struct {
		Metadata  Metadata `nimona:"@metadata:m"`
		RequestID string   `nimona:"requestID:s"`
		Object    *Object  `nimona:"object:m"`
	}
)

func (e *Request) Type() string {
	return "nimona.io/Request"
}

func (e *Request) MarshalObject() (*Object, error) {
	o, err := Marshal(e)
	if err != nil {
		return nil, err
	}
	o.Type = "nimona.io/Request"
	return o, nil
}

func (e *Request) UnmarshalObject(o *Object) error {
	return Unmarshal(o, e)
}

func (e *Response) Type() string {
	return "nimona.io/Response"
}

func (e *Response) MarshalObject() (*Object, error) {
	o, err := Marshal(e)
	if err != nil {
		return nil, err
	}
	o.Type = "nimona.io/Response"
	return o, nil
}

func (e *Response) UnmarshalObject(o *Object) error {
	return Unmarshal(o, e)
}
