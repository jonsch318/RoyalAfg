package random

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/models"
)

func TestSelectCards(t *testing.T) {
	t.Run("Should Generate 5", func(t *testing.T) {
		cards, err := SelectCards(5)
		assert.Empty(t, err)
		assert.Len(t, cards, 5)
	})
	t.Run("Should Generate Unique", func(t *testing.T) {
		for i := 0; i < 10000; i++ {
			cards, err := SelectCards(20)
			assert.Empty(t, err)
			seen := make([]models.Card,0)
			for _, card := range cards {
				for _, m := range seen {
					require.NotEqualValues(t, m, card)
				}
				seen = append(seen, card)
			}
			require.Len(t, cards, 20)
		}

	})

	t.Run("Should Generate Random", func(t *testing.T) {
		cards, err := SelectCards(10)
		assert.Empty(t, err)
		cards2, err := SelectCards(10)
		assert.Empty(t, err)
		assert.NotEqual(t, cards, cards2)
	})
}

func TestGetCards(t *testing.T) {
	t.Run("Should Generate Unique", func(t *testing.T) {
		cards := GetCards()
		seen := make([]models.Card,0)
		for _, card := range cards {
			for _, m := range seen {
				assert.NotEqualValues(t, m, card)
			}
			seen = append(seen, card)
		}
		assert.Len(t, cards, 52)
	})
	t.Run("Should Keep in Bound", func(t *testing.T) {
		cards := GetCards()
		assert.Equal(t, cards[0], models.Card{Color: 0,Value: 0})
		assert.Equal(t, cards[len(cards) -1], models.Card{Color: 3,Value: 12})
		assert.Len(t, cards, 52)
	})

}