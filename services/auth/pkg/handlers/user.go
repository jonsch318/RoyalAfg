package handlers

import (
	"github.com/JohnnyS318/RoyalAfgInGo/services/auth/pkg/interfaces"

	"go.uber.org/zap"
)

type User struct {
	l           *zap.SugaredLogger
	Auth interfaces.AuthenticationService
}

func NewUserHandler(logger *zap.SugaredLogger, auth interfaces.AuthenticationService) *User {
	return &User{
		l:           logger,
		Auth: auth,
	}
}