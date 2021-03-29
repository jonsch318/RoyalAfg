package handlers

import (
	"agones.dev/agones/pkg/client/clientset/versioned"
	"go.uber.org/zap"

	"github.com/JohnnyS318/RoyalAfgInGo/services/poker-matchmaker/pkg/lobby"
)


type Ticket struct {
	logger       *zap.SugaredLogger
	agonesClient versioned.Interface
	manager      *lobby.Manager
}

func NewTicket(logger *zap.SugaredLogger, agonesClient versioned.Interface, manager *lobby.Manager) *Ticket {
	return &Ticket{
		logger:       logger,
		agonesClient: agonesClient,
		manager:      manager,
	}
}
