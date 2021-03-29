package authentication

import (
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/dtos"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/models"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/mw"
	"github.com/JohnnyS318/RoyalAfgInGo/services/auth/pkg/services/user"
)

//IAuthentication provides functions to register login and logout of a account. It connects to the necesairy services to do this.
type IAuthentication interface {
	//Login logs the user with the given name and password in after comparing credentials
	Login(username, password string) (*models.User, string, error)
	//VerifyAuthentication verifies whether the user has a active session
	VerifyAuthentication(user *mw.UserClaims) bool
	//Register registers a new user with the given information
	Register(dto *dtos.RegisterDto) (*models.User, string, error)
	//Logout logs of the user
	Logout(id string) error
}

//Authentication is responsible for controlling the flow of authentication.
type Authentication struct {
	UserService user.IUser
}

//NewAuthentication provides a new authentication service with the given dependencies
func NewAuthentication(userService user.IUser) *Authentication {
	return &Authentication{
		UserService: userService,
	}
}
