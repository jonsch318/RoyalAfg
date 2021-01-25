package models

type LobbyBase struct {
	LobbyID string `json:"id"`
	Class *Class `json:"class"`
	ClassIndex int `json:"classIndex"`
	PlayerCount int `json:"playerCount"`
}

func NewLobby(lobbyId string, class *Class, classIndex int) *LobbyBase {
	return &LobbyBase{lobbyId, class, classIndex,0}
}
