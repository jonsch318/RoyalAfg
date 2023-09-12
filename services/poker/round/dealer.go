package round

import (
	"errors"

	"github.com/jonsch318/royalafg/pkg/log"
	"github.com/jonsch318/royalafg/services/poker/events"
	"github.com/jonsch318/royalafg/services/poker/utils"
)

func (r *Round) searchByActiveID(id string) (int, error) {
	for i, n := range r.Players {
		if n.ID == id && n.Active {
			return i, nil
		}
	}
	return -1, errors.New("player not in game")
}

func (r *Round) search(id string) (int, error) {
	for i, n := range r.Players {
		if n.ID == id {
			return i, nil
		}
	}
	return -1, errors.New("player not in game")
}

func (r *Round) sendDealer() {
	log.Logger.Infow("dealer send", "dealer", r.Dealer)
	if !r.Ended {
		utils.SendToAll(r.Players, events.NewDealerSetEvent(&r.PublicPlayers[r.Dealer], r.Dealer))
	}
}
