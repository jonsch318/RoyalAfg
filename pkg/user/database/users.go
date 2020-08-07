package database

import (
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/user/models"
	"github.com/Kamva/mgm/v3"
	"go.uber.org/zap"
)

type Users struct {
	logger *zap.SugaredLogger
	coll   *mgm.Collection
}

func NewUserDatabase(logger *zap.SugaredLogger) *Users {
	return &Users{
		logger: logger,
		coll:   mgm.Coll(&models.User{}),
	}
}

func (db *Users) CreateUser(user *models.User) error {
	err := user.Validate()

	if err != nil {
		return err
	}

	db.coll.Create(user)
	db.logger.Info("Inserted new User ", user.Username)
	return nil
}

func (db *Users) DeleteUser(user *models.User) error {
	db.coll.Delete(user)
	return nil
}
