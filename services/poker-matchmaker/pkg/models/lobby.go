package models

type Lobby struct {
	LobbyID string
	Class   int
	Players int
}

func NewLobby(lobbyId string, class int) *Lobby {
	return &Lobby{lobbyId, class, 0}
}
