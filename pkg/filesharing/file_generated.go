// Code generated by nimona.io/tools/codegen. DO NOT EDIT.

package filesharing

import (
	chore "nimona.io/pkg/chore"
	object "nimona.io/pkg/object"
)

type (
	File struct {
		Metadata object.Metadata `nimona:"@metadata:m"`
		Name     string          `nimona:"name:s"`
		Chunks   []chore.Hash    `nimona:"chunks:as"`
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

func (e *TransferDone) Type() string {
	return "nimona.io/TransferDone"
}

func (e *TransferRequest) Type() string {
	return "nimona.io/TransferRequest"
}

func (e *TransferResponse) Type() string {
	return "nimona.io/TransferResponse"
}
