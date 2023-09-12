package authentication

import (
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"

	pAuth "github.com/jonsch318/royalafg/pkg/auth"
	"github.com/jonsch318/royalafg/pkg/dtos"
	"github.com/jonsch318/royalafg/pkg/models"
	"github.com/jonsch318/royalafg/services/auth/pkg/security"
)

func (auth *Authentication) Register(dto *dtos.RegisterDto) (*models.User, string, error) {
	user := models.NewUser(dto.Username, dto.Email, dto.FullName, dto.Birthdate)
	user.ID = primitive.NewObjectID()

	hash, err := security.HashPassword(dto.Password, "")
	if err != nil {
		return nil, "", err
	}

	if hash == "" {
		return nil, "", fmt.Errorf("hash from password %v could not be created", dto.Password)
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

	token, err := pAuth.GetJwt(user)

	return user, token, err
}
