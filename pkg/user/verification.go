package user

const (
	Unverified      uint8 = 0
	EmailVerified   uint8 = 1
	AgeVerified     uint8 = 2
	NameVerified    uint8 = 4
	PaymentVerified uint8 = 8
)

// Sets the verification number based on the given bools
func SetVerification(prev uint8, email, age, name, payment bool) uint8 {
	newVer := uint8(0)
	if email {
		newVer |= EmailVerified
	}
	if age {
		newVer |= AgeVerified
	}
	if name {
		newVer |= NameVerified
	}
	if payment {
		newVer |= PaymentVerified
	}

	return prev | newVer
}
