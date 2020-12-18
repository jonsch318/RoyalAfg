package authentication

import (
	"errors"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/models"
	"github.com/JohnnyS318/RoyalAfgInGo/services/auth/config"
	"github.com/JohnnyS318/RoyalAfgInGo/services/auth/pkg/security"
	"github.com/JohnnyS318/RoyalAfgInGo/services/auth/pkg/services"
	"github.com/spf13/viper"
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
	var token string
	token, err = services.GenerateBearerToken(user)
	return user, token, err
}
