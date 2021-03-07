package random

import (
	"crypto/rand"
	"errors"
	"log"

	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/models"
)

//SelectCards is used to generate random cards that dont repeat
func SelectCards(count int) ([]models.Card, error) {
	if count > 52 {
		return nil, errors.New("more requested cards than possible")
	}

	cards := GetCards()
	selected := make([]models.Card, count)
	rest := byte(51)
	picked := make([]bool, 51)

	for i := 0; i < count; i++ {
		success := false
		for !success {
			b := make([]byte, 1)
			_, err := rand.Read(b)
			if err != nil {
				log.Printf("Could not generate random bytes %v", err.Error())
				return nil, err
			}

			if b[0] > (rest - 1) {
				continue
			}
			//card was already chosen
			if picked[b[0]] {
				continue
			}

			//Card is free

			success = true
			picked[b[0]] = true
			selected[i] = cards[b[0]]
			log.Printf("Picked b[0]=%v for card %v", b[0], cards[b[0]])
		}
	}
	return selected, nil
}

//Reset resets the card stack and therefore the probabilities
func GetCards() [52]models.Card {
	cards := [52]models.Card{}
	for i := 0; i < 52; i += 4 {
		for j := 0; j < 4; j++ {

			c := new(models.Card)
			c, err := models.NewCard(j, i/4)

			if err != nil {
				log.Printf("Something went very wrong %v", err)

				// Card selection is a key feature and should not fail
				// => closing lobby
			}
			cards[i+j] = *c
		}
	}
	return cards
}
