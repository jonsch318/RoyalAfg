package database

import (
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/models"
)

//IUserDB is the interface for the user database.
type IUserDB interface {
	//CreateUser saves the new user to the database. If a user with the same id exists already, a error will be returned.
	CreateUser(user *models.User) error
	//UpdateUser updates the existing user.
	UpdateUser(user *models.User) error
	//DeleteUser deletes the user from the database.
	DeleteUser(user *models.User) error
	//FindById returns the user, if found, with the given id
	FindById(id string) (*models.User, error)
	//FindByEmail returns the user, if found, with the given email
	FindByEmail(email string) (*models.User, error)
	//FindByUsername returns the user, if found, with the given username
	FindByUsername(username string) (*models.User, error)
}
