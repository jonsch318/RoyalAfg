package models

import (
	"time"

	"github.com/Kamva/mgm"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type User struct {
	mgm.DefaultModel `bson:",inline"`
	Username         string    `json:"username" bson:"username"`
	Email            string    `json:"email" bson:"email"`
	Hash             string    `json:"-" bson:"hash"`
	FullName         string    `json:"fullName" bson:"fullName"`
	Birthdate        time.Time `json:"birthdate" bson:"bithdate"`
}

func (user User) Validate() error {
	return validation.ValidateStruct(&user,
		validation.Field(&user.Username, validation.Required, validation.Length(4, 100)),
		validation.Field(&user.Email, is.Email),
		validation.Field(&user.Hash, validation.Required),
		validation.Field(&user.FullName, validation.Required, validation.Length(1, 100)),
		validation.Field(&user.Birthdate, validation.Required, validation.Date(time.RFC3339).Min(time.Now().AddDate(-16, 0, 0)).Max(time.Now().AddDate(-150, 0, 0))),
	)
}

// NewUser creates a new user with the given details.
//IMPORTANT THIS DOES NOT SAVE OR HASH THE PASSWORD. This has to be done seperatly
func NewUser(username, email, fullName string, birthdate time.Time) *User {
	return &User{
		Username:  username,
		Email:     email,
		FullName:  fullName,
		Birthdate: birthdate,
	}
}
