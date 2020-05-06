package net

import (
	"nimona.io/internal/rand"
	"nimona.io/pkg/context"
	"nimona.io/pkg/crypto"
	"nimona.io/pkg/keychain"
	"nimona.io/pkg/log"
	"nimona.io/pkg/object"
)

// Handshake middleware
type Handshake struct {
	keychain keychain.Keychain
}

// Handle incoming and outgoing connections
// TODO needs to be able to handle both server and client interactions
func (hs *Handshake) Handle() MiddlewareHandler {
	return func(ctx context.Context, conn *Connection) (
		*Connection, error) {
		if conn.IsIncoming {
			return hs.handleIncoming(ctx, conn)
		}
		return hs.handleOutgoing(ctx, conn)
	}
}

func (hs *Handshake) handleIncoming(
	ctx context.Context,
	conn *Connection,
) (*Connection, error) {
	logger := log.
		FromContext(ctx).
		Named("net/middleware/handleIncoming").
		With(
			log.String("remote_addr", conn.RemoteAddr()),
		)
	logger.Debug("handling inc connection, sending syn")

	ks := hs.keychain.List(keychain.PeerKey)
	k := hs.keychain.GetPrimaryPeerKey()

	pks := make([]crypto.PublicKey, len(ks))
	for i := 0; i < len(ks); i++ {
		pks[i] = ks[i].PublicKey()
	}
	pk := pks[0]

	nonce := rand.String(8)
	syn := &HandshakeSyn{
		Nonce:  nonce,
		Owners: pks,
	}
	so := syn.ToObject()
	sig, err := object.NewSignature(k, so)
	if err != nil {
		return nil, err
	}

	so = so.AddSignature(sig)
	if err := Write(so, conn); err != nil {
		return nil, err
	}

	logger.Debug("sent syn, waiting syn-ack")

	synAckObj, err := Read(conn)
	if err != nil {
		return nil, err
	}

	synAck := &HandshakeSynAck{}
	if err := synAck.FromObject(*synAckObj); err != nil {
		return nil, err
	}

	if synAck.Nonce != nonce {
		return nil, ErrNonce
	}

	logger.Debug("got syn-ack, sending ack")

	// store who is on the other side
	// TODO Exchange relies on this nees to be somewhere else?
	if len(synAck.Signatures) == 0 {
		return nil, ErrMissingSignature
	}
	conn.RemotePeerKey = synAck.Signatures[0].Signer
	conn.LocalPeerKey = pk

	// TODO(@geoah) do we need to do something about this?
	// hs.discoverer.Add(synAck.Peer)

	ack := &HandshakeAck{
		Nonce: nonce,
	}

	ao := ack.ToObject()
	sig, err = object.NewSignature(k, ao)
	if err != nil {
		return nil, err
	}

	ao = ao.AddSignature(sig)
	if err := Write(ao, conn); err != nil {
		return nil, err
	}

	logger.Debug("sent acl, done")

	return conn, nil
}

func (hs *Handshake) handleOutgoing(
	ctx context.Context,
	conn *Connection,
) (*Connection, error) {
	logger := log.
		FromContext(ctx).
		Named("net/middleware/handleOutgoing").
		With(
			log.String("remote_addr", conn.RemoteAddr()),
		)
	logger.Debug("handling out connection, waiting for syn")

	ks := hs.keychain.List(keychain.PeerKey)
	k := hs.keychain.GetPrimaryPeerKey()

	pks := make([]crypto.PublicKey, len(ks))
	for i := 0; i < len(ks); i++ {
		pks[i] = ks[i].PublicKey()
	}
	pk := pks[0]

	synObj, err := Read(conn)
	if err != nil {
		logger.Warn("waiting for syn failed", log.Error(err))
		// TODO close conn?
		return nil, err
	}

	syn := &HandshakeSyn{}
	if err := syn.FromObject(*synObj); err != nil {
		logger.Warn("could not convert obj to syn")
		// TODO close conn?
		return nil, err
	}

	logger.Debug("got syn, sending syn-ack")

	// store the remote peer
	if len(syn.Signatures) == 0 {
		return nil, ErrMissingSignature
	}
	conn.RemotePeerKey = syn.Signatures[0].Signer
	conn.LocalPeerKey = pk

	// TODO(@geoah) this one too
	// hs.discoverer.Add(syn.Peer)

	synAck := &HandshakeSynAck{
		Nonce:  syn.Nonce,
		Owners: pks,
	}
	sao := synAck.ToObject()
	sig, err := object.NewSignature(k, sao)
	if err != nil {
		logger.Warn(
			"could not sign for syn ack object",
			log.Error(err),
		)
		// TODO close conn?
		return nil, nil
	}

	sao = sao.AddSignature(sig)
	if err := Write(sao, conn); err != nil {
		logger.Warn("sending for syn-ack failed", log.Error(err))
		// TODO close conn?
		return nil, nil
	}

	logger.Debug("sent syn-ack, waiting ack")

	ackObj, err := Read(conn)
	if err != nil {
		logger.Warn("waiting for ack failed", log.Error(err))
		// TODO close conn?
		return nil, nil
	}

	ack := &HandshakeAck{}
	if err := ack.FromObject(*ackObj); err != nil {
		// TODO close conn?
		logger.Warn("could not convert obj to syn ack")
		return nil, nil
	}

	if ack.Nonce != syn.Nonce {
		logger.Warn("validating syn to ack nonce failed")
		// TODO close conn?
		return nil, nil
	}

	logger.Debug("got ack, done")

	return conn, nil
}