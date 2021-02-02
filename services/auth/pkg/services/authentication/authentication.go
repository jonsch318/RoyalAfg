package authentication

import (
	"github.com/JohnnyS318/RoyalAfgInGo/services/auth/pkg/interfaces"
)

//AuthenticationService is responsible for controlling the flow of authentication.
type Service struct {
	UserService interfaces.UserService
}

func NewService(userService interfaces.UserService ) *Service {
	return &Service{
		UserService: userService,
	}
}
