package protos

import (
	"time"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/shared/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FromMessageUser(m *User) *models.User {
	user := models.NewUser(m.GetUsername(), m.GetEmail(), m.GetFullName(), m.GetBirthdate())

	user.Hash = m.Hash

	return user
}

func FromMessageUserExact(m *User) *models.User {
	user := &models.User{
		Username:  m.Username,
		Email:     m.Email,
		Birthdate: m.Birthdate,
		FullName:  m.FullName,
		Hash:      m.Hash,
	}

	id, _ := primitive.ObjectIDFromHex(m.Id)

	user.ID = id

	user.CreatedAt = time.Unix(m.CreatedAt, 0)
	user.UpdatedAt = time.Unix(m.UpdatedAt, 0)

	return user
}

func ToMessageUser(user *models.User) *User {
	return &User{
		CreatedAt: user.CreatedAt.Unix(),
		UpdatedAt: user.UpdatedAt.Unix(),
		Id:        user.ID.Hex(),
		Username:  user.Username,
		FullName:  user.FullName,
		Birthdate: user.Birthdate,
		Email:     user.Email,
		Hash:      user.Hash,
	}
}
