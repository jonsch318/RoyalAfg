package round

import (
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/events"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/models"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/utils"
)

//holeCards randomly pics 2 cards for each player
func holeCards(players []models.Player, h map[string][2]models.Card, cards []models.Card) {
	j := 0
	for i := range players {
		if players[i].Active && j + 1 < len(cards) {
			var c [2]models.Card
			c[0] = cards[j]
			c[1] = cards[j+1]
			j += 2
			h[players[i].ID] = c
			utils.SendToPlayer(&players[i], events.NewHoleCardsEvent(c))
		}
	}
}
