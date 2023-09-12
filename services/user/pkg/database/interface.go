package database

import (
	"github.com/jonsch318/royalafg/pkg/models"
)

// UserDB is the interface for the user database.
type UserDB interface {
	//CreateUser saves the new user to the database. If a user with the same id exists already, a error will be returned.
	CreateUser(user *models.User) error
	//UpdateUser updates the existing user.
	UpdateUser(user *models.User) error
	//DeleteUser deletes the user from the database.
	DeleteUser(user *models.User) error
	//FindById returns the user, if found, with the given id. This is Cached
	FindById(id string) (*models.User, error)
	//FindByEmail returns the user, if found, with the given email. This is NOT Cached
	FindByEmail(email string) (*models.User, error)
	//FindByUsername returns the user, if found, with the given username. This is NOT Cached
	FindByUsername(username string) (*models.User, error)
}

type OnlineStatusDB interface {
	//SetOnlineStatus sets the online status of the user
	SetOnlineStatus(id string, status *OnlineStatus) error

	//GetOnlineStatus returns the online status of the user
	GetOnlineStatus(id string) (*OnlineStatus, error)
}
