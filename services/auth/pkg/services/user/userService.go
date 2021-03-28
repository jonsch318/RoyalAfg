package user

import (
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/models"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/protos"
)

//IUser is responsible to communicate with the user service
type IUser interface {
	//GetUserById returns the user with the given id if found
	GetUserById(id string) (*models.User, error)
	//GetUserById returns the user with the given username or email if found
	GetUserByUsernameOrEmail(usernameOrEmail string) (*models.User, error)
	//SaveUser saves the new user to the user service
	SaveUser(user *models.User) error
}

//User communicates with the user services
type User struct {
	Client protos.UserServiceClient
}

//NewUser creates a new user service
func NewUser(client protos.UserServiceClient) *User {
	return &User{
		Client: client,
	}
}
