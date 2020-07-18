package user

import (
	"github.com/Rhymond/go-money"
)

type Hasher interface {
	Hash(string) (string, err)
}

type User struct {
	ID string
	Username string
	Email string
	balance money.Money
}

func (user *User) HashPassword(hasher user.Hasher, password string) {
	hash, err := hasher.Hash()
}
