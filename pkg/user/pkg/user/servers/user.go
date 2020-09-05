package servers

import (
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/protos"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/shared/pkg/models"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/user/pkg/user/database"
	"go.uber.org/zap"
)

type UserServer struct {
	l  *zap.SugaredLogger
	db *database.UserDatabase
}

func NewUserServer(logger *zap.SugaredLogger, database *database.UserDatabase) *UserServer {
	return &UserServer{
		l:  logger,
		db: database,
	}
}

func toMessageUser(user *models.User) *protos.User {
	return &protos.User{
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

func fromMessageUser(m *protos.User) *models.User {
	user := models.NewUser(m.GetUsername(), m.GetEmail(), m.GetFullName(), m.GetBirthdate())

	user.Hash = m.Hash

	return user
}

func fromMessageUserExact(m *protos.User) *models.User {
	user := models.NewUser(m.GetUsername(), m.GetEmail(), m.GetFullName(), m.GetBirthdate())

	user.Hash = m.Hash

	return user
}
