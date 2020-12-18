package round

import (
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/events"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/models"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/showdown"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/utils"
	"log"
	"time"
)

func (h *Round) Start(players []models.Player, publicPlayers []models.PublicPlayer, dealer int) {

	h.Bank.Reset()

	h.Dealer = dealer
	h.Players = players
	h.InCount = byte(len(players))
	h.HoleCards = make(map[string][2]models.Card, len(players))

	//publish players and position

	time.Sleep(3 * time.Second)

	h.InnerStart()

	for i := range h.Players {
		utils.SendToPlayerInList(h.Players, i, events.NewGameStartEvent(publicPlayers, i))
	}

	time.Sleep(3 * time.Second)

	// Publish choosen Dealer

	h.sendDealer()
	time.Sleep(3 * time.Second)

	//set predefined blinds
	err := h.setBlinds()

	if err != nil {
		h.Bank.ConcludeRound(nil)
		return
	}

	// Set players hole cards
	holeCards(h.Players, h.HoleCards, h.cardGen)

	time.Sleep(3 * time.Second)

	h.actions(true)

	// Flop, turn and river are done here
	for i := 0; i < 5; i++ {
		h.Board[i] = h.cardGen.SelectRandom()
	}

	// send flop result

	utils.SendToAll(h.Players, events.NewFlopEvent(h.Board))

	//
	h.actions(false)
	// send turn result
	utils.SendToAll(h.Players, events.NewTurnEvent(h.Board))

	h.WhileNotEnded(func() {
		h.actions(false)
		// send river result
		utils.SendToAll(h.Players, events.NewRiverEvent(h.Board))
	})

	h.WhileNotEnded(func() {
		h.actions(false)
	})

	winners := showdown.Evaluate(h.Players, h.HoleCards, h.Board)
	winningPlayers := make([]int, 0)
	for i := range winners {
		_, i, err := utils.SearchByID(h.Players, winners[i])
		if err == nil {
			winningPlayers = append(winningPlayers, i)
		}
	}

	winningPublic := make([]models.PublicPlayer, len(winningPlayers))
	for i, n := range winningPlayers {
		winningPublic[i] = publicPlayers[n]
	}

	log.Printf("Winners: %v", winningPlayers)

	share := h.Bank.ConcludeRound(winners)

	utils.SendToAll(h.Players, events.NewGameEndEvent(winningPublic, share))

}

func (h *Round) InnerStart() {
	h.sendDealer()
}

func (h *Round) End() {
	log.Printf("Ending Hand due to error")
	h.Ended = true
}
