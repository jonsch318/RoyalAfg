package handlers

import (
	"agones.dev/agones/pkg/client/clientset/versioned"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"

	"github.com/JohnnyS318/RoyalAfgInGo/services/poker-matchmaker/pkg/lobby"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker-matchmaker/pkg/models"
)

type Ticket struct {
	logger       *zap.SugaredLogger
	rdb          *redis.Client
	lobbies      []models.Lobby
	agonesClient versioned.Interface
	manager      *lobby.Manager
}

func NewTicket(logger *zap.SugaredLogger, rdb *redis.Client, agonesClient versioned.Interface, manager *lobby.Manager) *Ticket {
	return &Ticket{
		logger:       logger,
		rdb:          rdb,
		lobbies:      make([]models.Lobby, 0),
		agonesClient: agonesClient,
		manager:      manager,
	}
}
