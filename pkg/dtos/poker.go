package dtos

import (
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/poker/models"
)

type ClassInfoResponse struct {
	Classes []models.Class `json:"classes"`
}

type LobbyInfoResponse struct {
	Lobbies [][]models.LobbyBase `json:"lobbies"`
}

type PokerInfoResponse struct {
	Classes []models.Class `json:"classes"`
	Lobbies [][]models.LobbyBase `json:"lobbies"`
}