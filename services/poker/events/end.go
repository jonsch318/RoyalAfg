package events

import "github.com/JohnnyS318/RoyalAfgInGo/services/poker/models"

type GameEndEvent struct {
	Winners []models.PublicPlayer `json:"winners"`
	Shares   string               `json:"shares"`
}

func NewGameEndEvent(winners []models.PublicPlayer, shares string) *models.Event {
	return models.NewEvent(GAME_END, &GameEndEvent{
		Winners: winners,
		Shares:   shares,
	})
}
