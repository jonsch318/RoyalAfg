package models

type LobbyClass struct {
	Min int
	Max int
	Blind int
}

func NewLobbyClass(min, max, blind int) *LobbyClass{
	return &LobbyClass{min, max, blind}
}