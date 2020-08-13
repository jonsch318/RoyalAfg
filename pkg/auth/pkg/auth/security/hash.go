package security

import scrypt "github.com/elithrar/simple-scrypt"

func HashPassword(password, pepper string) (string, error) {
	hashBytes, err := scrypt.GenerateFromPassword(addPepper(password, pepper), scrypt.DefaultParams)
	if err != nil {
		return "", err
	}
	return string(hashBytes), nil
}

// ComparePassword compares the password to the registered hash.
func ComparePassword(password, hash, pepper string) bool {
	return scrypt.CompareHashAndPassword([]byte(hash), addPepper(password, pepper)) == nil
}

func addPepper(password, pepper string) []byte {
	return []byte(password + pepper)
}
