package events

import (
	"github.com/Rhymond/go-money"
	"github.com/jonsch318/royalafg/services/poker/models"
)

type GameEndEvent struct {
	Winners   []models.PublicPlayer `json:"winners"`
	Shares    string                `json:"shares"`
	SharesNum float64               `json:"sharesNum"`
}

func NewGameEndEvent(winners []models.PublicPlayer, shares *money.Money) *models.Event {
	return models.NewEvent(GAME_END, &GameEndEvent{
		Winners:   winners,
		Shares:    shares.Display(),
		SharesNum: shares.AsMajorUnits(),
	})
}

type LobbyPause struct {
	Players     []models.PublicPlayer `json:"players"`
	PlayerCount int                   `json:"playerCount"`
}

func NewLobbyPauseEvent(players []models.PublicPlayer, playerCount int) *models.Event {
	return models.NewEvent(LOBBY_PAUSE, &LobbyPause{
		Players:     players,
		PlayerCount: playerCount,
	})
}
