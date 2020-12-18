package showdown

import (
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/models"
	"testing"
)

func TestRankSpecificHand(t *testing.T) {
	cards := [][]models.Card{{c(0, 11), c(2, 11), c(2, 5), c(1, 11), c(3, 8)}, {c(0, 2), c(1, 5), c(2, 5), c(3, 11), c(3, 8)},
		{c(0, 3), c(0, 5), c(0, 4), c(0, 6), c(0, 7)}, {c(0, 11), c(0, 8), c(0, 10), c(0, 9), c(0, 12)}}

	for i := range cards {
		rank := rankSpecificHand(cards[i])
		t.Logf("Rank %v", rank)
	}
	return
}
