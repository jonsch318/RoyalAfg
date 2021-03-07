package authentication

import (
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/models"
	"github.com/JohnnyS318/RoyalAfgInGo/services/auth/pkg/dto"
	"github.com/JohnnyS318/RoyalAfgInGo/services/auth/pkg/security"
	"github.com/JohnnyS318/RoyalAfgInGo/services/auth/pkg/services"
)

func (auth Service) Register(dto *dto.RegisterDto) (*models.User, string, error) {
	user := models.NewUser(dto.Username, dto.Email, dto.FullName, dto.Birthdate)
	user.ID = primitive.NewObjectID()

	hash, err := security.HashPassword(dto.Password, "")
	if err != nil {
		return nil, "", err
	}

	if hash == "" {
		return nil, "", fmt.Errorf("Hash from password %v could not be created", dto.Password)
	}

	user.Hash = hash

	log.Printf("User hash: %v", user.Hash)

	if err = user.Validate(); err != nil {
		return nil, "", err
	}

	err = auth.UserService.SaveUser(user)
	if err != nil {
		return nil, "", err
	}

	token, err := services.GetJwt(user)

	return user, token, err
}
