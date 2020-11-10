// Code generated by nimona.io/tools/codegen. DO NOT EDIT.

package peer

import (
	crypto "nimona.io/pkg/crypto"
	object "nimona.io/pkg/object"
)

type (
	ConnectionInfo struct {
		Metadata  object.Metadata   `nimona:"metadata:m,omitempty"`
		PublicKey crypto.PublicKey  `nimona:"publicKey:s,omitempty"`
		Addresses []string          `nimona:"addresses:as,omitempty"`
		Relays    []*ConnectionInfo `nimona:"relays:ao,omitempty"`
	}
)

func (e *ConnectionInfo) Type() string {
	return "nimona.io/peer.ConnectionInfo"
}

func (e ConnectionInfo) ToObject() *object.Object {
	o, err := object.Encode(&e)
	if err != nil {
		panic(err)
	}
	return o
}

func (e *ConnectionInfo) FromObject(o *object.Object) error {
	return object.Decode(o, e)
}
