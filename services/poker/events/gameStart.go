package events

import "github.com/JohnnyS318/RoyalAfgInGo/services/poker/models"

type GameStartEvent struct {
	Players  []models.PublicPlayer `json:"players"`
	Position int                   `json:"position"`
	Pot string `json:"pot"`
}

func NewGameStartEvent(publicPlayers []models.PublicPlayer, position int, pot string) *models.Event {
	return models.NewEvent(GAME_START, &GameStartEvent{
		Position: position,
		Players:  publicPlayers,
		Pot: pot,
	})
}

type HoleCardsEvent struct {
	Cards [2]models.Card `json:"cards"`
}

func NewHoleCardsEvent(cards [2]models.Card) *models.Event {
	return models.NewEvent(HOLE_CARDS, &HoleCardsEvent{
		Cards: cards,
	})
}
