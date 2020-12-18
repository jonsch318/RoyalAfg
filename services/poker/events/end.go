package events

import "github.com/JohnnyS318/RoyalAfgInGo/services/poker/models"

type GameEndEvent struct {
	Winners []models.PublicPlayer `json:"winners"`
	Share   float32               `json:"share"`
}

func NewGameEndEvent(winners []models.PublicPlayer, share int) *models.Event {
	return models.NewEvent(GAME_END, &GameEndEvent{
		Winners: winners,
		Share:   float32(share) / 100,
	})
}
