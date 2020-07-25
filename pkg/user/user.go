package user

import (
	"github.com/Kamva/mgm/v3"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"github.com/elithrar/simple-scrypt"
)

type User struct {
	mgm.DefaultModel `bson:",inline"`
	Username string `json:"username"`
	Email string `json:"email, omitempty"`
	hash string
}

// NewUser creates a new user with the given details and hashes the password
func NewUser(username, password, email string) (*User, error)  {
	user := &User{
		Username: username,
		Email: email,
	}

	err := user.hashPassword(password)

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (user *User) hashPassword(password string) error {
	var params scrypt.Params
	if viper.IsSet("PasswordHashParams"){
		params = viper.Get("PasswordHashParams").(scrypt.Params)
	} else {
		params = scrypt.DefaultParams;
	}
	hashBytes, err := scrypt.GenerateFromPassword(addPepper(password), params)
	if err != nil {
		return err
	}

	user.hash = string(hashBytes)
	return nil
}

// ComparePassword compares the password to the registered hash.
func (user *User) ComparePassword(password string) bool {
	return scrypt.CompareHashAndPassword([]byte(user.hash), addPepper(password)) == nil
}

func addPepper(password string) []byte {
	return []byte(password + viper.GetString("PasswordHashPepper"))
}