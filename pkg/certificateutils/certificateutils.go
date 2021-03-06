package certificateutils

import (
	"encoding/json"
	"fmt"

	"nimona.io/pkg/context"
	"nimona.io/pkg/crypto"
	"nimona.io/pkg/errors"
	"nimona.io/pkg/network"
	"nimona.io/pkg/object"
	"nimona.io/pkg/objectstore"
)

func WaitForCertificateResponse(
	ctx context.Context,
	net network.Network,
	csr *object.CertificateRequest,
) <-chan *object.CertificateResponse {
	ch := make(chan *object.CertificateResponse)
	go func() {
		sub := net.Subscribe(
			network.FilterByObjectType(
				object.CertificateResponseType,
			),
		)
		subCh := sub.Channel()
		defer sub.Cancel()
		defer close(ch)
		for {
			select {
			case <-ctx.Done():
				return
			case env := <-subCh:
				csrRes := &object.CertificateResponse{}
				if csrRes == nil {
					return
				}
				b, _ := json.MarshalIndent(env.Payload, "", "  ")
				fmt.Println(string(b))
				if err := object.Unmarshal(env.Payload, csrRes); err != nil {
					continue
				}
				select {
				case ch <- csrRes:
				default:
				}
				return
			}
		}
	}()
	return ch
}

func FindCertificateResponseForPeer(
	ctx context.Context,
	str objectstore.Store,
	peerPublicKey crypto.PublicKey,
) (*object.CertificateResponse, error) {
	rdr, err := str.GetByType(
		object.CertificateResponseType,
	)
	if err != nil {
		return nil, err
	}
	for {
		obj, err := rdr.Read()
		if err != nil {
			break
		}
		c := &object.CertificateResponse{}
		if err := object.Unmarshal(obj, c); err != nil {
			return nil, err
		}
		if c.Certificate.Subject.Equals(peerPublicKey) {
			return c, nil
		}
	}
	return nil, errors.Error("not found")
}
