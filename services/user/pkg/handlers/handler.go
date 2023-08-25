package handlers

import (
	"github.com/JohnnyS318/RoyalAfgInGo/services/user/pkg/database"
	"go.uber.org/zap"
)

type UserHandler struct {
	l        *zap.SugaredLogger
	db       database.UserDB
	statusDB database.OnlineStatusDB
}

func NewUserHandler(logger *zap.SugaredLogger, userDB database.UserDB, statusDB database.OnlineStatusDB) *UserHandler {
	return &UserHandler{
		l:        logger,
		db:       userDB,
		statusDB: statusDB,
	}
}
