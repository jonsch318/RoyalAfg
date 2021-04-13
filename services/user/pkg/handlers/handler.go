package handlers

import (
	"github.com/JohnnyS318/RoyalAfgInGo/services/user/pkg/database"
	"go.uber.org/zap"
)

type UserHandler struct {
	l  *zap.SugaredLogger
	db database.IUserDB
}

func NewUserHandler(logger *zap.SugaredLogger, database database.IUserDB) *UserHandler {
	return &UserHandler{
		l:  logger,
		db: database,
	}
}
