package handlers

import "github.com/JohnnyS318/RoyalAfgInGo/services/poker/lobbies"

type Lobby struct {
	Lobbies *lobbies.LobbyManager
}

func NewLobbyHandler(lobbyManager *lobbies.LobbyManager) *Lobby {
	return &Lobby{
		Lobbies: lobbyManager,
	}
}
