package utils

import (
	"log"
	"math/rand"

	"github.com/jonsch318/royalafg/services/poker/models"
)

// CardGenerator is a utility for randomly choose a multiple cards out of a 52 card deck
type CardGenerator struct {
	Cards         [52]models.Card
	SelectedCards int
}

// NewCardSelector creates a new card select
func NewCardSelector() *CardGenerator {
	s := &CardGenerator{SelectedCards: 0}
	s.Reset()
	return s
}

// Reset resets the card stack and therefore the probabilities
func (s *CardGenerator) Reset() {
	for i := 0; i < 48; i += 4 {
		for j := 0; j < 4; j++ {

			c := new(models.Card)
			c, err := models.NewCard(j, i/4)

			if err != nil {
				log.Printf("Something went wrong %v", err)

				// Card selection is a key feature and should not fail
				// => closing lobby
			}
			s.Cards[i+j] = *c
		}
	}
	s.SelectedCards = 0
}

// SelectRandom randomly selects a card from the stack and returns a copy of it
func (s *CardGenerator) SelectRandom() (c models.Card) {
	c, s.SelectedCards = SelectRandom(s.Cards, s.SelectedCards)
	return c
}

// SelectRandom randomly selects a card from the stack and returns a copy of it
func SelectRandom(cards [52]models.Card, selected int) (models.Card, int) {
	i := rand.Intn(51 - selected)
	c := cards[i]
	// if selected card is last. the last card has not to be swapped
	if i != 51-selected {
		cards[51-selected], cards[i] = c, cards[51-selected]
		selected++
	}
	return c, selected
}

func SelectRandomN(s int) (int, int) {
	i := rand.Intn(51 - s)
	c := i
	// if selected card is last. the last card has not to be swapped
	if i != 51-s {
		s++
	}
	return c, s
}

func (s *CardGenerator) Select(n int) []models.Card {
	if n >= 51-s.SelectedCards {
		return nil
	}

	//b := n*(51-s.SelectedCards) - n

	return nil
}

/*//Select selects the given number of random bytes in range [0-start - picked]
func Select(n byte, start byte) []byte {
	var b [n*10]byte
	binary.Read(crand.Reader, binary.BigEndian, b)
}*/
