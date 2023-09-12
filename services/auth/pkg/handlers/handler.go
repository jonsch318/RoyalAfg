package handlers

import (
	"go.uber.org/zap"

	"github.com/jonsch318/royalafg/services/auth/pkg/rabbit"
	"github.com/jonsch318/royalafg/services/auth/pkg/services/authentication"
)

// Auth is the user handler to handle authentication requests of users.
type Auth struct {
	Auth   authentication.IAuthentication
	l      *zap.SugaredLogger
	Rabbit *rabbit.RabbitMessageBroker
}

// NewAuth creates a new user handler with the specified dependencies.
func NewAuth(logger *zap.SugaredLogger, auth authentication.IAuthentication, rabbit *rabbit.RabbitMessageBroker) *Auth {
	return &Auth{
		Auth:   auth,
		l:      logger,
		Rabbit: rabbit,
	}
}
