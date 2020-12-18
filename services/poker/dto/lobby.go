package dto

type PublicLobby struct {
	ID          string `json:"id"`
	PlayerCount int    `json:"playerCount"`
	MinBuyIn    int    `json:"minBuyIn"`
	MaxBuyIn    int    `json:"maxBuyIn"`
	SmallBlind  int    `json:"smallBlind"`
	LobbyClass  int    `json:"lobbyClass"`
}

func ToPublic(id string, playerCount, minBuyIn, maxBuyIn, smallBlind, lobbyClass int) *PublicLobby {
	return &PublicLobby{
		ID:          id,
		PlayerCount: playerCount,
		MinBuyIn:    minBuyIn,
		MaxBuyIn:    maxBuyIn,
		SmallBlind:  smallBlind,
		LobbyClass:  lobbyClass,
	}
}
