package crypto

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/binary"
	"time"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/errors"
	"github.com/vechain/go-ecvrf"
)

type VRFNumberGenerator struct {
	sk   *ecdsa.PrivateKey
	pk   *ecdsa.PublicKey
	seed []byte
}

// Create a new VRF number generator from the given private/public key.
func NewVRFNumberGenerator(sk *ecdsa.PrivateKey, pk *ecdsa.PublicKey) *VRFNumberGenerator {

	return &VRFNumberGenerator{
		sk: sk,
		pk: pk,
	}
}

// Generate
func (v *VRFNumberGenerator) Generate() (alpha, beta, pi []byte, err error) {

	//Generate the alpha value
	nowBytes := make([]byte, 8)
	now := time.Now().UnixNano()
	binary.LittleEndian.PutUint64(nowBytes, uint64(now))
	alpha = append(nowBytes, v.seed...)

	h := sha256.New()
	_, err = h.Write(alpha)

	if err != nil {
		return nil, nil, nil, err
	}

	alpha = h.Sum(nil)

	beta, pi, err = ecvrf.P256Sha256Tai.Prove(v.sk, alpha)
	return
}

func (v *VRFNumberGenerator) Verify(alpha []byte, beta []byte, pi []byte) (bool, error) {
	betaCheck, err := ecvrf.P256Sha256Tai.Verify(v.pk, alpha, pi)

	if err != nil {
		return false, err
	}

	if betaCheck == nil {
		return false, &errors.VerifyFailedError{Details: "could not generate beta"}
	}

	if !bytes.Equal(beta, betaCheck) {
		return false, &errors.VerifyFailedError{Details: "beta does not match"}
	}

	return true, nil
}

// Get the public key of the VRF number generator.
func (v *VRFNumberGenerator) GetPublicKey() *ecdsa.PublicKey {
	return v.pk
}
