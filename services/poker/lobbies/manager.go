package lobbies

import (
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/dto"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/lobby"
	"sync"
)

//LobbyManager keeps track of the poker lobbies and can distribute players among those lobbies.
type LobbyManager struct {
	lock           sync.RWMutex
	BuyInClasses   [][]int
	LobbiesIndexed [][]string
	Lobbies        map[string]*lobby.Lobby
	PublicLobbies  [][]dto.PublicLobby
	MaxCount       int
	LobbyCount     int
}

//NewManager creates a new LobbyManagerInstance that can manage lobbies and distribute players among the lobbies.
func NewManager(maxCount int, classes [][]int) *LobbyManager {
	return &LobbyManager{
		BuyInClasses:   classes,
		MaxCount:       maxCount,
		Lobbies:        make(map[string]*lobby.Lobby),
		LobbiesIndexed: make([][]string, len(classes)),
		PublicLobbies:  make([][]dto.PublicLobby, len(classes)),
	}
}
