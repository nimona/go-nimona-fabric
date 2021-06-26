// Code generated by nimona.io/tools/codegen. DO NOT EDIT.

package identity

import (
	chore "nimona.io/pkg/chore"
	crypto "nimona.io/pkg/crypto"
	object "nimona.io/pkg/object"
)

const ProfileType = "nimona.io/identity.Profile"

type Profile struct {
	Metadata object.Metadata `nimona:"@metadata:m,type=nimona.io/identity.Profile"`
	Version  int64           `nimona:"version:i"`
	Datetime string          `nimona:"datetime:s"`
	Name     string          `nimona:"name:s"`
	Image    chore.Hash      `nimona:"image:s"`
}

const ProfileStreamRootType = "stream:nimona.io/identity.Profile"

type ProfileStreamRoot struct {
	Metadata object.Metadata `nimona:"@metadata:m,type=stream:nimona.io/identity.Profile"`
}

const ProfileUpdatedType = "event:nimona.io/identity.Profile.Updated"

type ProfileUpdated struct {
	Metadata object.Metadata `nimona:"@metadata:m,type=event:nimona.io/identity.Profile.Updated"`
	Profile  Profile         `nimona:"profile:m"`
}

const AddressbookStreamRootType = "stream:nimona.io/identity/addressbook"

type AddressbookStreamRoot struct {
	Metadata object.Metadata `nimona:"@metadata:m,type=stream:nimona.io/identity/addressbook"`
}

const AddressbookContactAddedType = "event:nimona.io/identity/addressbook.ContactAdded"

type AddressbookContactAdded struct {
	Metadata    object.Metadata  `nimona:"@metadata:m,type=event:nimona.io/identity/addressbook.ContactAdded"`
	Alias       string           `nimona:"alias:s"`
	RemoteParty crypto.PublicKey `nimona:"remoteParty:s"`
	Profile     Profile          `nimona:"profile:m"`
	Datetime    string           `nimona:"datetime:s"`
}

const AddressbookContactRemovedType = "event:nimona.io/identity/addressbook.ContactRemoved"

type AddressbookContactRemoved struct {
	Metadata    object.Metadata  `nimona:"@metadata:m,type=event:nimona.io/identity/addressbook.ContactRemoved"`
	RemoteParty crypto.PublicKey `nimona:"remoteParty:s"`
	Datetime    string           `nimona:"datetime:s"`
}
