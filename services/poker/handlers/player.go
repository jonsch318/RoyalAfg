package handlers

import (
	sdk "agones.dev/agones/sdks/go"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/lobbies"
)

type Lobby struct {
	Lobbies *lobbies.LobbyManager
	sdk *sdk.SDK
}

func NewLobbyHandler(lobbyManager *lobbies.LobbyManager, sdk *sdk.SDK) *Lobby {
	return &Lobby{
		Lobbies: lobbyManager,
		sdk: sdk,
	}
}
