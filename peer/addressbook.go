package peer

import (
	"errors"
	"sync"
	"time"

	"github.com/keybase/saltpack/basic"
)

var (
	ErrCannotPutLocalPeerInfo = errors.New("cannot put local peer info")

	peerInfoExpireAfter = time.Hour * 1
)

type PeerManager interface {
	GetLocalPeerInfo() *SecretPeerInfo
	PutLocalPeerInfo(*SecretPeerInfo) error // TODO Deprecate

	GetPeerInfo(peerID string) (*PeerInfo, error)
	GetAllPeerInfo() ([]*PeerInfo, error)
	PutPeerInfo(*PeerInfo) error

	PutPeerStatus(peerID string, status Status)
	GetPeerStatus(peerID string) Status

	// Resolve(ctx context.Context, peerID string) (string, error)
	// Discover(ctx context.Context, peerID, protocol string) ([]net.Address, error)
	LoadOrCreateLocalPeerInfo(path string) (*SecretPeerInfo, error)
	CreateNewPeer() (*SecretPeerInfo, error)
	LoadSecretPeerInfo(path string) (*SecretPeerInfo, error)
	StoreSecretPeerInfo(pi *SecretPeerInfo, path string) error
	GetKeyring() *basic.Keyring
}

// Status represents the connection state of a peer
type Status int

const (
	NotConnected Status = iota
	Connected
	CanConnect
	ErrorConnecting
)

// NewAddressBook creates a new AddressBook with an empty keyring
func NewAddressBook() *AddressBook {
	adb := &AddressBook{
		identities: &IdentityCollection{},
		peers:      &PeerInfoCollection{},
		keyring:    basic.NewKeyring(),
	}

	return adb
}

type AddressBook struct {
	identities    *IdentityCollection
	peers         *PeerInfoCollection
	peerStatus    sync.Map
	localPeerLock sync.RWMutex
	localPeer     *SecretPeerInfo
	keyring       *basic.Keyring
}

func (adb *AddressBook) GetKeyring() *basic.Keyring {
	return adb.keyring
}

func (adb *AddressBook) PutPeerInfo(peerInfo *PeerInfo) error {
	if adb.localPeer.ID == peerInfo.ID {
		return ErrCannotPutLocalPeerInfo
	}

	if peerInfo.ID == "" {
		return nil
	}

	peerInfo.UpdatedAt = time.Now()
	return adb.peers.Put(peerInfo)
}

func (adb *AddressBook) GetLocalPeerInfo() *SecretPeerInfo {
	return adb.localPeer
}

func (adb *AddressBook) PutLocalPeerInfo(peerInfo *SecretPeerInfo) error {
	adb.localPeerLock.Lock()
	defer adb.localPeerLock.Unlock()
	peerInfo.UpdatedAt = time.Now()
	adb.localPeer = peerInfo
	return nil
}

func (adb *AddressBook) GetPeerInfo(peerID string) (*PeerInfo, error) {
	return adb.peers.Get(peerID)
}

func (adb *AddressBook) GetAllPeerInfo() ([]*PeerInfo, error) {
	return adb.peers.All()
}

func (adb *AddressBook) PutPeerStatus(peerID string, status Status) {
	adb.peerStatus.Store(peerID, status)
}

func (adb *AddressBook) GetPeerStatus(peerID string) Status {
	status, ok := adb.peerStatus.Load(peerID)
	if !ok {
		return NotConnected
	}

	return status.(Status)
}
