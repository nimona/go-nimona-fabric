// Code generated by nimona.io/tools/codegen. DO NOT EDIT.

package object

import (
	crypto "nimona.io/pkg/crypto"
)

const CertificateType = "nimona.io/Certificate"

type Certificate struct {
	Metadata    Metadata                `nimona:"@metadata:m,type=nimona.io/Certificate"`
	Nonce       string                  `nimona:"nonce:s"`
	Subject     crypto.PublicKey        `nimona:"subject:s"`
	Permissions []CertificatePermission `nimona:"permissions:am"`
	Starts      string                  `nimona:"starts:s"`
	Expires     string                  `nimona:"expires:s"`
}

const CertificatePermissionType = "nimona.io/CertificatePermission"

type CertificatePermission struct {
	Metadata Metadata `nimona:"@metadata:m,type=nimona.io/CertificatePermission"`
	Types    []string `nimona:"types:as"`
	Actions  []string `nimona:"actions:as"`
}

const CertificateRequestType = "nimona.io/CertificateRequest"

type CertificateRequest struct {
	Metadata               Metadata                `nimona:"@metadata:m,type=nimona.io/CertificateRequest"`
	Nonce                  string                  `nimona:"nonce:s"`
	VendorName             string                  `nimona:"vendorName:s"`
	VendorURL              string                  `nimona:"vendorURL:s"`
	ApplicationName        string                  `nimona:"applicationName:s"`
	ApplicationDescription string                  `nimona:"applicationDescription:s"`
	ApplicationURL         string                  `nimona:"applicationURL:s"`
	Permissions            []CertificatePermission `nimona:"permissions:am"`
}

const CertificateResponseType = "nimona.io/CertificateResponse"

type CertificateResponse struct {
	Metadata    Metadata           `nimona:"@metadata:m,type=nimona.io/CertificateResponse"`
	Signed      bool               `nimona:"signed:b"`
	Notes       string             `nimona:"notes:s"`
	Request     CertificateRequest `nimona:"request:m"`
	Certificate Certificate        `nimona:"certificate:m"`
}
