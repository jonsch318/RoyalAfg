package models

type Lobby struct {
	LobbyId string
	class int
	Players int
}

func NewLobby(lobbyId string, class int) *Lobby {
	return &Lobby{lobbyId, class, 0}
}