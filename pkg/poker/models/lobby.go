package models

import (
	"fmt"
)

//LobbyBase is the base model for a poker lobby. The Poker Matchmaker does not need any more information. But the poker gameserver extends this
type LobbyBase struct {
	LobbyID     string `json:"id"`
	Class       *Class `json:"class"`
	ClassIndex  int    `json:"classIndex"`
	PlayerCount int    `json:"playerCount"`
}

//NewLobby returns a new LobbyBase
func NewLobby(lobbyId string, class *Class, classIndex int) *LobbyBase {
	return &LobbyBase{lobbyId, class, classIndex, 0}
}

//String formulates a custom string for logging
func (l *LobbyBase) String() string {
	return fmt.Sprintf("Lobby [%s] (%v) => %v", l.LobbyID, l.PlayerCount, l.Class)
}
