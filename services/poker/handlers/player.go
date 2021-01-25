package handlers

import (
	sdk "agones.dev/agones/sdks/go"

	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/lobby"
)

type Game struct {
	lby          *lobby.Lobby
	sdk          *sdk.SDK
	stopShutdown chan interface{}
}

func NewGame(lobbyInstance *lobby.Lobby, sdk *sdk.SDK, stopShutdown chan interface{}) *Game {
	return &Game{
		lby:          lobbyInstance,
		sdk:          sdk,
		stopShutdown: stopShutdown,
	}
}
