package showdown

import (
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/models"
	"log"
)

//Evaluate evaluates the given poker scenario and determins the winners based on a rank given to each player.
func Evaluate(players []models.Player, cards map[string][2]models.Card, board [5]models.Card) []string {

	if len(players) < 1 {
		log.Printf("Player under 1")
		return nil
	}

	if len(players) == 1 {
		//h.In[0] wins
		id := make([]string, 1)
		id[0] = players[0].ID
		return id
	}

	ranks := make(map[string]int)
	for i := range players {
		if players[i].Active {
			cards := cards[players[i].ID]
			rank := evaluatePlayer(append(cards[:], board[:]...))
			ranks[players[i].ID] = rank
			log.Printf("Ranking Player %v => %d", players[i].ID, rank)
		}
	}
	winners := make([]string, 0)

	highestRank := 0

	// Determin winner or winners
	for k, v := range ranks {
		if v == highestRank {
			// add to winner because rank is equal
			winners = append(winners, k)
		}
		if v > highestRank {
			//reset winners
			winners = nil
			highestRank = v
			winners = append(winners, k)
		}

	}

	return winners
}

//evaluatePlayer generates a number as an identification of the players hole cards + the boards cards rank. it selects the best card seection and return the rank.
func evaluatePlayer(cards []models.Card) int {

	maxRank := rankSpecificHand(cards[2:])
	for i := -1; i < 5; i++ {
		for j := -1; j < 5; j++ {
			if i == j {
				log.Printf("Skipped: %d", i)
				continue
			}

			//swap

			if i > -1 {
				cards[i+2], cards[0] = cards[0], cards[i+2]
			}

			if j > -1 {
				cards[j+2], cards[1] = cards[1], cards[j+2]
			}

			r := rankSpecificHand(cards[2:])
			log.Printf("Set : %v => %v", cards[2:], r)

			if r > maxRank {
				maxRank = r
			}

			// swap back
			if i > -1 {
				cards[i+2], cards[0] = cards[0], cards[i+2]
			}

			if j > -1 {
				cards[j+2], cards[1] = cards[1], cards[j+2]
			}
		}
	}

	return maxRank
}

/*
//evaluatePlayer generates a number as an identification of the players hole cards + the boards cards rank. it selects the best card seection and return the rank.
func evaluatePlayer(cards []models.Card) int {

	maxRank := rankSpecificHand(cards[2:])
	if rank := rankSpecificHand(cards[1:6]); rank > maxRank {
		maxRank = rank
	}
	if rank := rankSpecificHand(cards[0:5]); rank > maxRank {
		maxRank = rank
	}
	return maxRank
} */

func normalizeAce(number int) int {
	if number-1 < 0 {
		return 12
	}
	return number
}
