package lobby

import (
	"agones.dev/agones/pkg/client/clientset/versioned"
	"github.com/go-redis/redis/v8"

	"github.com/JohnnyS318/RoyalAfgInGo/services/poker-matchmaker/pkg/models"
)

type Manager struct {
	lobbies [][]models.Lobby
	classes []models.LobbyClass
	agonesClient versioned.Interface
	rdg *redis.Client
}

func NewManager(agonesClient versioned.Interface, classes []models.LobbyClass) *Manager{
	ret := &Manager{
		lobbies: make([][]models.Lobby, len(classes)),
		classes: classes,
		agonesClient: agonesClient,
	}

	return ret
}




