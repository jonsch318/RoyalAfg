package handlers

import (
	"go.uber.org/zap"

	"github.com/JohnnyS318/RoyalAfgInGo/services/auth/pkg/interfaces"
	"github.com/JohnnyS318/RoyalAfgInGo/services/auth/pkg/rabbit"
)

//Auth is the user handler to handle authentication requests of users.
type Auth struct {
	Auth interfaces.AuthenticationService
	l    *zap.SugaredLogger
	Rabbit *rabbit.RabbitMessageBroker
}

//NewAuth creates a new user handler with the specified dependencies.
func NewAuth(logger *zap.SugaredLogger, auth interfaces.AuthenticationService, rabbit *rabbit.RabbitMessageBroker) *Auth {
	return &Auth{
		Auth: auth,
		l:    logger,
		Rabbit: rabbit,
	}
}
