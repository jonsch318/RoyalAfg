package events

import "github.com/JohnnyS318/RoyalAfgInGo/services/poker/models"

type GameStartEvent struct {
	Players  []models.PublicPlayer `json:"players"`
	Player *models.PublicPlayer `json:"player"`
	Pot string `json:"pot"`
	Position int                   `json:"position"`
}

func NewGameStartEvent(publicPlayers []models.PublicPlayer, player *models.PublicPlayer, position int, pot string) *models.Event {
	return models.NewEvent(GAME_START, &GameStartEvent{
		Players:  publicPlayers,
		Player: player,
		Pot: pot,
		Position: position,
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
