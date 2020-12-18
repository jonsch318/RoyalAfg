package round

import (
	"errors"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/events"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/utils"
	"log"
)

func (h *Round) setBlinds() error {
	success := false
	s := 0

	var smallBlindIndex int
	for !success {

		if s >= len(h.Players)-1 {
			return errors.New("Nobody can bet smallBlind and bigBlind")
		}

		var i int
		for j := 1; j < len(h.Players)-1; j++ {
			i = (h.Dealer + j) % len(h.Players)
			if h.Players[i].Active && h.Dealer != i {
				break
			}
		}

		err := h.Bank.Bet(h.Players[i].ID, h.SmallBlind)

		if err != nil {
			log.Printf("Folding player due to invalid buyin in smallBlind %v", err)
			h.Fold(h.Players[i].ID)
			s++
			continue
		}

		utils.SendToAll(h.Players, events.NewActionProcessedEvent(2, h.SmallBlind, i, h.Bank.GetPlayerBet(h.Players[i].ID), h.Bank.GetPlayerWallet(h.Players[i].ID)))
		success = true
		smallBlindIndex = i
		s = 0
	}

	success = false
	for !success {

		if s >= len(h.Players) {
			return errors.New("Nobody can bet bigBlind")
		}

		var i int
		for j := 1; j <= len(h.Players); j++ {
			i = (smallBlindIndex + j) % len(h.Players)
			if h.Players[i].Active && h.Dealer != i && smallBlindIndex != i {
				break
			}
		}

		err := h.Bank.Bet(h.Players[i].ID, h.SmallBlind*2)

		if err != nil {
			log.Printf("Folding player due to invalid buyin in bigBlind")
			h.Fold(h.Players[i].ID)
			s++
			continue
		}
		utils.SendToAll(h.Players, events.NewActionProcessedEvent(2, h.SmallBlind*2, i, h.Bank.GetPlayerBet(h.Players[i].ID), h.Bank.GetPlayerWallet(h.Players[i].ID)))
		success = true
		h.bigBlindIndex = i
	}

	return nil
}
