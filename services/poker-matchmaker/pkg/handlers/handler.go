package handlers

import (
	"go.uber.org/zap"

	"github.com/jonsch318/royalafg/services/poker-matchmaker/pkg/lobby"
)

type Ticket struct {
	logger  *zap.SugaredLogger
	manager *lobby.Manager
}

func NewTicket(logger *zap.SugaredLogger, manager *lobby.Manager) *Ticket {
	return &Ticket{
		logger:  logger,
		manager: manager,
	}
}
