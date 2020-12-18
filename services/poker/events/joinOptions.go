package events

import "github.com/JohnnyS318/RoyalAfgInGo/services/poker/dto"

type JoinOptions struct {
	BuyInClasses [][]float32         `json:"buyInClasses"`
	Lobbies      [][]dto.PublicLobby `json:"lobbies"`
}
