package authentication

import (
	"errors"

	"github.com/spf13/viper"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/models"
	"github.com/JohnnyS318/RoyalAfgInGo/services/auth/config"
	"github.com/JohnnyS318/RoyalAfgInGo/services/auth/pkg/security"
	"github.com/JohnnyS318/RoyalAfgInGo/services/auth/pkg/services"
)

func (auth Service) Login(username, password string) (*models.User, string, error) {

	user, err := auth.UserService.GetUserByUsernameOrEmail(username)
	if err != nil {
		return nil, "", err
	}

	if !security.ComparePassword(password, user.Hash, viper.GetString(config.Pepper)) {
		return nil, "", errors.New("passwords did not match")
	}
	//TODO: Execute other login schemes (2FA)

	token, err := services.GetJwt(user)
	return user, token, err
}
