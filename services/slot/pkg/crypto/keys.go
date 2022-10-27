package crypto

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"os"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/errors"
)

func ReadECDSAKeys(skPath string, pkPath string) (*ecdsa.PrivateKey, *ecdsa.PublicKey, error) {

	if skPath == "" || pkPath == "" {
		return nil, nil, &errors.MissingKeyError{}
	}

	sk, err := os.ReadFile(skPath)
	if err != nil {
		return nil, nil, err
	}

	pk, err := os.ReadFile(pkPath)
	if err != nil {
		return nil, nil, err
	}

	return decodePEMKey(sk, pk)
}

func decodePEMKey(sk []byte, pk []byte) (*ecdsa.PrivateKey, *ecdsa.PublicKey, error) {
	block, rest := pem.Decode(sk)

	if rest != nil {
		if bytes.Equal(rest, sk) {
			return nil, nil, &errors.InvalidKeyError{Details: "no PEM data found"}
		} else {
			return nil, nil, &errors.InvalidKeyError{Details: "trailing data found after PEM data"}
		}
	}

	privateKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, nil, &errors.InvalidKeyError{Details: err.Error()}
	}

	//Public Key Decoding
	blockPub, restPub := pem.Decode(pk)

	if restPub != nil {
		if bytes.Equal(restPub, sk) {
			return nil, nil, &errors.InvalidKeyError{Details: "no PEM data found"}
		} else {
			return nil, nil, &errors.InvalidKeyError{Details: "trailing data found after PEM data"}
		}
	}

	genericPubKey, err := x509.ParsePKIXPublicKey(blockPub.Bytes)

	if err != nil {
		return nil, nil, &errors.InvalidKeyError{Details: err.Error()}
	}

	publicKey, ok := genericPubKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, nil, &errors.InvalidKeyError{Details: "public key is not an ECDSA key"}
	}

	return privateKey, publicKey, nil
}

// Encode the given public key to a base64 string of PKIX, ASN.1 DER form.
func ExportPublicKey(pk *ecdsa.PublicKey) (string, error) {
	keyBytes, err := x509.MarshalPKIXPublicKey(pk)

	if err != nil {
		return "", err
	}

	key := base64.StdEncoding.EncodeToString(keyBytes)
	return key, nil
}
