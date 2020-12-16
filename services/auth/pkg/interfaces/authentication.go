package interfaces

import (
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/models"
	"github.com/JohnnyS318/RoyalAfgInGo/services/auth/pkg/dto"
)

type AuthenticationService interface {
	Login(username, password string) (*models.User,string,error)
	VerifyAuthentication() (bool, error)
	Register(dto *dto.RegisterDto) (*models.User, string, error)
	Logout(id string) error
}
