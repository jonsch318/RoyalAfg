package events

import (
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/models"
)

type DealerSetEvent struct {
	Player *models.PublicPlayer `json:"player"`
	Index  int                  `json:"index"`
}

func NewDealerSetEvent(player *models.PublicPlayer, index int) *models.Event{
	return models.NewEvent(DEALER_SET, &DealerSetEvent{
		Player: player,
		Index:  index,
	})
}
