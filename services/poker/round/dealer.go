package round

import (
	"errors"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/events"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/models"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/utils"
)

func (h *Round) searchByID(id string) (*models.Player, int, error) {
	for i, n := range h.Players {
		if n.ID == id {
			return &n, i, nil
		}
	}
	return nil, -1, errors.New("Player not in game")
}

func (h *Round) searchByActiveID(id string) (int, error) {
	for i, n := range h.Players {
		if n.ID == id && n.Active {
			return i, nil
		}
	}
	return -1, errors.New("Player not in game")
}

func (h *Round) sendDealer() {
	if !h.Ended {
		utils.SendToAll(h.Players, models.NewEvent(events.DEALER_SET, h.Dealer))
	}
}
