package events

import (
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/models"
	"github.com/Rhymond/go-money"
)

type GameStartEvent struct {
	Players  []models.PublicPlayer `json:"players"`
	Pot      string                `json:"pot"`
	PotNum   float64               `json:"potNum"`
	Position int                   `json:"position"`
}

func NewGameStartEvent(publicPlayers []models.PublicPlayer, position int, pot *money.Money) *models.Event {
	return models.NewEvent(GAME_START, &GameStartEvent{
		Players:  publicPlayers,
		Pot:      pot.Display(),
		PotNum:   pot.AsMajorUnits(),
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
