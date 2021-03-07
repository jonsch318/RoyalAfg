package showdown

import (
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/log"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/models"
)

type WinnerInfo struct{
	Player models.Player
	Position int
}

func (i *WinnerInfo) String() string {
	return i.Player.String()
}

//Evaluate evaluates the given poker scenario and determines the winners based on a rank given to each player.
func Evaluate(players []models.Player, cards map[string][2]models.Card, board [5]models.Card, inCount byte) []WinnerInfo {

	if len(players) < 1 || inCount < 1{
		log.Logger.Info("No players. Nobody wins")
		return nil
	}

	if len(players) == 1 {
		if players[0].Active{
			log.Logger.Debugf("Player 0 wins. One player remaining")
			//player wins
			return []WinnerInfo{
				{
					Player:   players[0],
					Position: 0,
				},
			}
		}else {
			return nil
		}
	}

	log.Logger.Debugf("Comparision of players")

	ranks := make(map[WinnerInfo]int)
	for i := range players {
		log.Logger.Debugf("Player [%s] has Left(%v) and is Active(%v)", players[i].Username, players[i].Left, players[i].Active)
		if players[i].Active && !players[i].Left {
			log.Logger.Debugf("Player [%s] is still active", players[i].Username)
			c := cards[players[i].ID]
			rank := evaluatePlayer(append(c[:], board[:]...))
			info := WinnerInfo{
				Player:   players[i],
				Position: i,
			}
			ranks[info] = rank
		}else {
			log.Logger.Debugf("Skipped player %v", players[i].Username)
		}
	}
	winners := make([]WinnerInfo, 0)

	highestRank := 0
	// Determine winner or winners
	for k, v := range ranks {
		log.Logger.Debugf("Player [%v] has card rank %v", k.Player.Username, v)
		if v == highestRank {
			// add to winners because rank is equal
			winners = append(winners, k)
		}
		if v > highestRank {
			//reset winners
			winners = make([]WinnerInfo, 0)
			highestRank = v
			winners = append(winners, k)
		}

	}

	return winners
}

//evaluatePlayer generates a number as an identification of the players hole cards + the boards cards rank. it selects the best card section and return the rank.
//this probably could be better optimised with technics like tree shaking, etc. but it works.
func evaluatePlayer(cards []models.Card) int {

	maxRank := rankSpecificHand(cards[2:])
	for i := -1; i < 5; i++ {
		for j := -1; j < 5; j++ {
			if i == j {
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