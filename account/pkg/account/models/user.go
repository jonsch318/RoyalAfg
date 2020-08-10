package models

import (
	"encoding/json"
	"io"

	"github.com/Kamva/mgm/v3"
	scrypt "github.com/elithrar/simple-scrypt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type User struct {
	mgm.DefaultModel `bson:",inline"`
	Username         string `json:"username" bson:"username"`
	Email            string `json:"email" bson:"email"`
	Hash             string `json:"-" bson:"hash"`
	FullName         string `json:"fullName" bson:"fullName"`
}

func (user User) Validate() error {
	return validation.ValidateStruct(&user,
		validation.Field(&user.Username, validation.Required, validation.Length(4, 100)),
		validation.Field(&user.Email, is.Email),
		validation.Field(&user.Hash, validation.Required),
		validation.Field(&user.FullName, validation.Required, validation.Length(1, 100)),
	)
}

func (user *User) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(user)
}

func (user *User) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(user)
}

// NewUser creates a new user with the given details and hashes the password
func NewUser(username, password, email, fullName string) (*User, error) {
	user := &User{
		Username: username,
		Email:    email,
		FullName: fullName,
	}

	err := user.SetPassword(password)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// SetPassword sets the password of the user
func (user *User) SetPassword(password string) error {
	err := validation.Validate(password, validation.Required, validation.Length(4, 100))
	if err != nil {
		return err
	}
	user.hashPassword(password, "")

	return nil
}

func (user *User) hashPassword(password, pepper string) {
	hashBytes, _ := scrypt.GenerateFromPassword(addPepper(password, pepper), scrypt.DefaultParams)
	user.Hash = string(hashBytes)
}

// ComparePassword compares the password to the registered hash.
func (user *User) ComparePassword(password, pepper string) bool {
	return scrypt.CompareHashAndPassword([]byte(user.Hash), addPepper(password, pepper)) == nil
}

func addPepper(password, pepper string) []byte {
	return []byte(password + pepper)
}
