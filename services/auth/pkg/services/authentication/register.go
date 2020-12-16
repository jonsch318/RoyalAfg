package authentication

import (
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/models"
	"github.com/JohnnyS318/RoyalAfgInGo/services/auth/pkg/dto"
	"github.com/JohnnyS318/RoyalAfgInGo/services/auth/pkg/security"
	"github.com/JohnnyS318/RoyalAfgInGo/services/auth/pkg/services"
)

func (auth AuthenticationService) Register(dto *dto.RegisterDto) (*models.User, string, error) {
	user := models.NewUser(dto.Username, dto.Email, dto.FullName, dto.Birthdate)
	hash, err := security.HashPassword(dto.Password, "")
	if err != nil {
		return nil, "", err
	}
	user.Hash = hash

	err = auth.UserService.SaveUser(user)
	if err != nil {
		return nil, "", err
	}

	token, err := services.GenerateBearerToken(user)

	return user, token, err
}