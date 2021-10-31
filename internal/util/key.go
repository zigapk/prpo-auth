package util

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io"
)

func EncodePubKey(key *rsa.PublicKey, w io.Writer) error {
	pubKeyBytes, err := x509.MarshalPKIXPublicKey(key)
	if err != nil {
		return err
	}

	pubKeyBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubKeyBytes,
	}

	err = pem.Encode(w, pubKeyBlock)
	if err != nil {
		return err
	}

	return nil
}
