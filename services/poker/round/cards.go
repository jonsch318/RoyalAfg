package round

import (
	"github.com/jonsch318/royalafg/pkg/log"
	"github.com/jonsch318/royalafg/services/poker/events"
	"github.com/jonsch318/royalafg/services/poker/models"
	"github.com/jonsch318/royalafg/services/poker/utils"
)

// holeCards randomly pics 2 cards for each player
func (r *Round) holeCards(cards []models.Card) {
	j := 0
	for i := range r.Players {
		if r.Players[i].Active && j+1 < len(cards) {
			var c [2]models.Card
			c[0] = cards[j]
			c[1] = cards[j+1]
			j += 2
			log.Logger.Debugf("Player received cards %v j=%v count=%v", c, j, len(cards))
			r.HoleCards[r.Players[i].ID] = c
			err := utils.SendToPlayerInListTimeout(r.Players, i, events.NewHoleCardsEvent(c))
			if err != nil {
				log.Logger.Infof("Could not send to player flagging as removed and folding")
			}
		}
	}
}
