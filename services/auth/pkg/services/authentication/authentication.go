package authentication

import "github.com/JohnnyS318/RoyalAfgInGo/services/auth/pkg/interfaces"

type AuthenticationService struct {
	UserService interfaces.UserService
}

func NewAuthenticationService(userService interfaces.UserService) *AuthenticationService{
	return &AuthenticationService{
		UserService: userService,
	}
}