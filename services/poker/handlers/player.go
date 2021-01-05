package handlers

import (
	sdk "agones.dev/agones/sdks/go"

	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/lobby"
)

type Game struct {
	lby *lobby.Lobby
	sdk *sdk.SDK
}

func NewGame(lobbyInstance *lobby.Lobby, sdk *sdk.SDK) *Game {
	return &Game{
		lby: lobbyInstance,
		sdk: sdk,
	}
}
