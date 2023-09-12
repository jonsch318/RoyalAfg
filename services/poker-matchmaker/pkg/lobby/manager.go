package lobby

import (
	"agones.dev/agones/pkg/client/clientset/versioned"
	"go.uber.org/zap"

	"github.com/jonsch318/royalafg/pkg/poker/models"
)

type Manager struct {
	lobbies      [][]models.LobbyBase
	classes      []models.Class
	agonesClient versioned.Interface
	logger       *zap.SugaredLogger
}

func NewManager(logger *zap.SugaredLogger, agonesClient versioned.Interface, classes []models.Class) *Manager {
	ret := &Manager{
		lobbies:      make([][]models.LobbyBase, len(classes)),
		classes:      classes,
		agonesClient: agonesClient,
		logger:       logger,
	}

	return ret
}
