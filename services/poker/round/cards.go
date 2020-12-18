package round

import (
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/events"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/models"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/utils"
)

//holeCards randomly pics 2 cards for each player
func holeCards(players []models.Player, h map[string][2]models.Card, gen *utils.CardGenerator) {
	for i := range players {
		if players[i].Active {
			var cards [2]models.Card
			cards[0] = gen.SelectRandom()
			cards[1] = gen.SelectRandom()
			h[players[i].ID] = cards
			utils.SendToPlayer(&players[i], events.NewHoleCardsEvent(cards))
		}
	}
}
