package dtos

import (
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/poker/models"
)

//ClassInfoResponse contains the registered poker lobby classes that are available
type ClassInfoResponse struct {
	Classes []models.Class `json:"classes"`
}

//LobbyInfoResponse returns the current active lobbies sorted by lobby classes
type LobbyInfoResponse struct {
	Lobbies [][]models.LobbyBase `json:"lobbies"`
}

//PokerInfoResponse is the combination of the ClassInfoResponse and the LobbyInfoResponse
type PokerInfoResponse struct {
	Classes []models.Class `json:"classes"`
	Lobbies [][]models.LobbyBase `json:"lobbies"`
}