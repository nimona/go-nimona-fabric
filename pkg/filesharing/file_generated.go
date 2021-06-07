// Code generated by nimona.io/tools/codegen. DO NOT EDIT.

package filesharing

import (
	object "nimona.io/pkg/object"
	value "nimona.io/pkg/object/value"
)

type (
	File struct {
		Metadata object.Metadata `nimona:"@metadata:m"`
		Name     string          `nimona:"name:s"`
		Chunks   []value.CID     `nimona:"chunks:as"`
	}
	TransferDone struct {
		Metadata object.Metadata `nimona:"@metadata:m"`
		Nonce    string          `nimona:"nonce:s"`
	}
	TransferRequest struct {
		Metadata object.Metadata `nimona:"@metadata:m"`
		File     File            `nimona:"file:m"`
		Nonce    string          `nimona:"nonce:s"`
	}
	TransferResponse struct {
		Metadata object.Metadata `nimona:"@metadata:m"`
		Nonce    string          `nimona:"nonce:s"`
		Accepted bool            `nimona:"accepted:b"`
	}
)

func (e *File) Type() string {
	return "nimona.io/File"
}

func (e *File) MarshalObject() (*object.Object, error) {
	o, err := object.Marshal(e)
	if err != nil {
		return nil, err
	}
	o.Type = "nimona.io/File"
	return o, nil
}

func (e *File) UnmarshalObject(o *object.Object) error {
	return object.Unmarshal(o, e)
}

func (e *TransferDone) Type() string {
	return "nimona.io/TransferDone"
}

func (e *TransferDone) MarshalObject() (*object.Object, error) {
	o, err := object.Marshal(e)
	if err != nil {
		return nil, err
	}
	o.Type = "nimona.io/TransferDone"
	return o, nil
}

func (e *TransferDone) UnmarshalObject(o *object.Object) error {
	return object.Unmarshal(o, e)
}

func (e *TransferRequest) Type() string {
	return "nimona.io/TransferRequest"
}

func (e *TransferRequest) MarshalObject() (*object.Object, error) {
	o, err := object.Marshal(e)
	if err != nil {
		return nil, err
	}
	o.Type = "nimona.io/TransferRequest"
	return o, nil
}

func (e *TransferRequest) UnmarshalObject(o *object.Object) error {
	return object.Unmarshal(o, e)
}

func (e *TransferResponse) Type() string {
	return "nimona.io/TransferResponse"
}

func (e *TransferResponse) MarshalObject() (*object.Object, error) {
	o, err := object.Marshal(e)
	if err != nil {
		return nil, err
	}
	o.Type = "nimona.io/TransferResponse"
	return o, nil
}

func (e *TransferResponse) UnmarshalObject(o *object.Object) error {
	return object.Unmarshal(o, e)
}