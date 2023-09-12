package events

import "github.com/jonsch318/royalafg/services/poker/models"

type BoardEvent struct {
	Cards []models.Card `json:"cards" mapstructure:"cards"`
}

func NewFlopEvent(cards [5]models.Card) *models.Event {
	return models.NewEvent(FLOP, &BoardEvent{Cards: cards[:3]})
}

func NewTurnEvent(cards [5]models.Card) *models.Event {
	return models.NewEvent(TURN, &BoardEvent{Cards: cards[3:4]})
}

func NewRiverEvent(cards [5]models.Card) *models.Event {
	return models.NewEvent(RIVER, &BoardEvent{Cards: cards[4:5]})
}
