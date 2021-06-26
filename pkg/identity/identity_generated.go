// Code generated by nimona.io/tools/codegen. DO NOT EDIT.

package identity

import (
	chore "nimona.io/pkg/chore"
	crypto "nimona.io/pkg/crypto"
	object "nimona.io/pkg/object"
)

type Profile struct {
	Metadata object.Metadata `nimona:"@metadata:m,type=nimona.io/identity.Profile"`
	Version  int64           `nimona:"version:i"`
	Datetime string          `nimona:"datetime:s"`
	Name     string          `nimona:"name:s"`
	Image    chore.Hash      `nimona:"image:s"`
}

func (e *Profile) Type() string {
	return "nimona.io/identity.Profile"
}

type ProfileStreamRoot struct {
	Metadata object.Metadata `nimona:"@metadata:m,type=stream:nimona.io/identity.Profile"`
}

func (e *ProfileStreamRoot) Type() string {
	return "stream:nimona.io/identity.Profile"
}

type ProfileUpdated struct {
	Metadata object.Metadata `nimona:"@metadata:m,type=event:nimona.io/identity.Profile.Updated"`
	Profile  Profile         `nimona:"profile:m"`
}

func (e *ProfileUpdated) Type() string {
	return "event:nimona.io/identity.Profile.Updated"
}

type AddressbookStreamRoot struct {
	Metadata object.Metadata `nimona:"@metadata:m,type=stream:nimona.io/identity/addressbook"`
}

func (e *AddressbookStreamRoot) Type() string {
	return "stream:nimona.io/identity/addressbook"
}

type AddressbookContactAdded struct {
	Metadata    object.Metadata  `nimona:"@metadata:m,type=event:nimona.io/identity/addressbook.ContactAdded"`
	Alias       string           `nimona:"alias:s"`
	RemoteParty crypto.PublicKey `nimona:"remoteParty:s"`
	Profile     Profile          `nimona:"profile:m"`
	Datetime    string           `nimona:"datetime:s"`
}

func (e *AddressbookContactAdded) Type() string {
	return "event:nimona.io/identity/addressbook.ContactAdded"
}

type AddressbookContactRemoved struct {
	Metadata    object.Metadata  `nimona:"@metadata:m,type=event:nimona.io/identity/addressbook.ContactRemoved"`
	RemoteParty crypto.PublicKey `nimona:"remoteParty:s"`
	Datetime    string           `nimona:"datetime:s"`
}

func (e *AddressbookContactRemoved) Type() string {
	return "event:nimona.io/identity/addressbook.ContactRemoved"
}
