package errors

import "fmt"

type InvalidKeyError struct {
	Details string
}

func (e *InvalidKeyError) Error() string {
	return fmt.Sprintf("the key is invalid: %s", e.Details)
}

type MissingKeyError struct{}

func (e *MissingKeyError) Error() string {
	return "no key found"
}

type InvalidKeyPairError struct{}

func (e *InvalidKeyPairError) Error() string {
	return "the key pair is invalid"
}

type VerifyFailedError struct {
	Details string
}

func (e *VerifyFailedError) Error() string {
	return fmt.Sprintf("the crypto verification failed: %s", e.Details)
}
