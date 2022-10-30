package handlers

import (
	"github.com/JohnnyS318/RoyalAfgInGo/services/user/pkg/database"
	"go.uber.org/zap"
)

type UserHandler struct {
	l        *zap.SugaredLogger
	db       database.IUserDB
	statusDB database.IOnlineStatusDB
}

func NewUserHandler(logger *zap.SugaredLogger, userDB database.IUserDB, statusDB database.IOnlineStatusDB) *UserHandler {
	return &UserHandler{
		l:        logger,
		db:       userDB,
		statusDB: statusDB,
	}
}
