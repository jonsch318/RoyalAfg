package lobby

import (
	"agones.dev/agones/pkg/client/clientset/versioned"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/poker/models"
)

type Manager struct {
	lobbies      [][]models.LobbyBase
	classes      []models.Class
	agonesClient versioned.Interface
	rdg          *redis.Client
	logger       *zap.SugaredLogger
}

func NewManager(logger *zap.SugaredLogger, agonesClient versioned.Interface, classes []models.Class, redisClient *redis.Client) *Manager {
	ret := &Manager{
		lobbies:      make([][]models.LobbyBase, len(classes)),
		classes:      classes,
		agonesClient: agonesClient,
		logger:       logger,
		rdg:          redisClient,
	}

	return ret
}
