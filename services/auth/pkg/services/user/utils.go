package user

import (
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/models"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/protos"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ToMessage(user *models.User) *protos.User {
	return &protos.User{
		Id:        user.ID.Hex(),
		Username:  user.Username,
		FullName:  user.FullName,
		Birthdate: user.Birthdate,
		Hash:      user.Hash,
		Email:     user.Email,
	}
}

func FromMessage(message *protos.User) *models.User {
	user := &models.User{
		Username:  message.Username,
		Email:     message.Email,
		Birthdate: message.Birthdate,
		FullName:  message.FullName,
		Hash:      message.Hash,
	}
	var err error
	user.ID, err = primitive.ObjectIDFromHex(message.Id)
	if err != nil {
		user.ID = primitive.NilObjectID
	}
	return user
}
