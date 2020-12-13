package handlers

import (
	"github.com/JohnnyS318/RoyalAfgInGo/services/user/pkg/user/database"
	"go.uber.org/zap"
)

type UserHandler struct {
	l  *zap.SugaredLogger
	db *database.UserDatabase
}

func NewUserHandler(logger *zap.SugaredLogger, database *database.UserDatabase) *UserHandler {
	return &UserHandler{
		l:  logger,
		db: database,
	}
}
