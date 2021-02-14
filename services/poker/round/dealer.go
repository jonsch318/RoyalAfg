package round

import (
	"errors"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/log"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/events"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/models"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/utils"
)

func (r *Round) searchByID(id string) (*models.Player, int, error) {
	for i, n := range r.Players {
		if n.ID == id {
			return &n, i, nil
		}
	}
	return nil, -1, errors.New("Player not in game")
}

func (r *Round) searchByActiveID(id string) (int, error) {
	for i, n := range r.Players {
		if n.ID == id && n.Active {
			return i, nil
		}
	}
	return -1, errors.New("Player not in game")
}

func (r *Round) sendDealer() {
	log.Logger.Infow("dealer send", r.Dealer)
	if !r.Ended {
		utils.SendToAll(r.Players, models.NewEvent(events.DEALER_SET, r.Dealer))
	}
}
