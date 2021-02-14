package round

import (
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/events"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/models"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/showdown"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/utils"
	"log"
	"time"
)

func (r *Round) Start(players []models.Player, publicPlayers []models.PublicPlayer, dealer int) {

	r.Bank.Reset()

	r.Dealer = dealer
	r.Players = players
	r.InCount = byte(len(players))
	r.HoleCards = make(map[string][2]models.Card, len(players))

	//publish players and position

	time.Sleep(3 * time.Second)

	r.InnerStart()

	for i := range r.Players {
		utils.SendToPlayerInList(r.Players, i, events.NewGameStartEvent(publicPlayers, i))
	}

	time.Sleep(3 * time.Second)

	// Publish choosen Dealer

	r.sendDealer()
	time.Sleep(3 * time.Second)

	//set predefined blinds
	err := r.setBlinds()

	if err != nil {
		r.Bank.ConcludeRound(nil)
		return
	}

	// Set players hole cards
	holeCards(r.Players, r.HoleCards, r.cardGen)

	time.Sleep(3 * time.Second)

	r.actions(true)

	// Flop, turn and river are done here
	for i := 0; i < 5; i++ {
		r.Board[i] = r.cardGen.SelectRandom()
	}

	// send flop result

	utils.SendToAll(r.Players, events.NewFlopEvent(r.Board))

	r.actions(false)
	// send turn result
	utils.SendToAll(r.Players, events.NewTurnEvent(r.Board))

	r.WhileNotEnded(func() {
		r.actions(false)
		// send river result
		utils.SendToAll(r.Players, events.NewRiverEvent(r.Board))
	})

	r.WhileNotEnded(func() {
		r.actions(false)
	})


	//Evaluation
	winners := showdown.Evaluate(r.Players, r.HoleCards, r.Board)
	winningPlayers := make([]int, 0)
	for i := range winners {
		_, i, err := utils.SearchByID(r.Players, winners[i])
		if err == nil {
			winningPlayers = append(winningPlayers, i)
		}
	}
	shares := r.Bank.ConcludeRound(winners)
	winningPublic := make([]models.PublicPlayer, len(winningPlayers))
	for i, n := range winningPlayers {
		winningPublic[i] = publicPlayers[n]
	}
	log.Printf("Winners: %v", winningPlayers)
	utils.SendToAll(r.Players, events.NewGameEndEvent(winningPublic, shares[0]))
}

func (r *Round) InnerStart() {
	r.sendDealer()
}

func (r *Round) End() {
	log.Printf("Ending Hand due to error")
	r.Ended = true
}
