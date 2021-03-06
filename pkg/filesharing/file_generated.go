// Code generated by nimona.io/tools/codegen. DO NOT EDIT.

package filesharing

import (
	chore "nimona.io/pkg/chore"
	object "nimona.io/pkg/object"
)

const FileType = "nimona.io/File"

type File struct {
	Metadata object.Metadata `nimona:"@metadata:m,type=nimona.io/File"`
	Name     string          `nimona:"name:s"`
	Chunks   []chore.Hash    `nimona:"chunks:as"`
}

const TransferDoneType = "nimona.io/TransferDone"

type TransferDone struct {
	Metadata object.Metadata `nimona:"@metadata:m,type=nimona.io/TransferDone"`
	Nonce    string          `nimona:"nonce:s"`
}

const TransferRequestType = "nimona.io/TransferRequest"

type TransferRequest struct {
	Metadata object.Metadata `nimona:"@metadata:m,type=nimona.io/TransferRequest"`
	File     File            `nimona:"file:m"`
	Nonce    string          `nimona:"nonce:s"`
}

const TransferResponseType = "nimona.io/TransferResponse"

type TransferResponse struct {
	Metadata object.Metadata `nimona:"@metadata:m,type=nimona.io/TransferResponse"`
	Nonce    string          `nimona:"nonce:s"`
	Accepted bool            `nimona:"accepted:b"`
}
