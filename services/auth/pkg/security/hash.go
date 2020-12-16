package security

import scrypt "github.com/elithrar/simple-scrypt"

func HashPassword(password, pepper string) (string, error) {
	hashBytes, err := scrypt.GenerateFromPassword(AddPepper(password, pepper), scrypt.DefaultParams)
	if err != nil {
		return "", err
	}
	return string(hashBytes), nil
}

// ComparePassword compares the password to the registered hash. True if passwords match, false if not or any errors occur
func ComparePassword(password, hash, pepper string) bool {
	return scrypt.CompareHashAndPassword([]byte(hash), AddPepper(password, pepper)) == nil
}

func AddPepper(password, pepper string) []byte {
	return []byte(password + pepper)
}
