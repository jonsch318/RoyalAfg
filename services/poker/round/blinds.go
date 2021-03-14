package round

import (
	"errors"
	"log"

	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/events"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/utils"
)

func (r *Round) setBlinds() error {
	success := false
	s := 0

	var smallBlindIndex int
	for !success {

		if s >= len(r.Players)-1 {
			return errors.New("Nobody can bet smallBlind and bigBlind")
		}

		var i int
		for j := 1; j < len(r.Players)-1; j++ {
			i = (r.Dealer + j) % len(r.Players)
			if r.Players[i].Active && r.Dealer != i {
				break
			}
		}

		err := r.Bank.PerformBlind(r.Players[i].ID, r.SmallBlind)

		if err != nil {
			log.Printf("Folding player due to invalid buyin in smallBlind %v", err)
			_ = r.Fold(r.Players[i].ID)
			s++
			continue
		}

		utils.SendToAll(r.Players, events.NewActionProcessedEvent(
				2,
				i,
				r.SmallBlind.Display(),
				r.Bank.GetPlayerBet(r.Players[i].ID),
				r.Bank.GetPlayerWallet(r.Players[i].ID),
				r.Bank.GetPot(),
		))
		success = true
		smallBlindIndex = i
		s = 0
	}

	success = false
	for !success {

		if s >= len(r.Players) {
			return errors.New("Nobody can bet bigBlind")
		}

		var i int
		for j := 1; j <= len(r.Players); j++ {
			i = (smallBlindIndex + j) % len(r.Players)
			if r.Players[i].Active && r.Dealer != i && smallBlindIndex != i {
				break
			}
		}

		bigBlind := r.SmallBlind.Multiply(2)
		err := r.Bank.PerformBlind(r.Players[i].ID, bigBlind)

		if err != nil {
			log.Printf("Folding player due to invalid buyin in bigBlind")
			_ = r.Fold(r.Players[i].ID)
			s++
			continue
		}
		utils.SendToAll(r.Players, events.NewActionProcessedEvent(
			2,
			i,
			bigBlind.Display(),
			r.Bank.GetPlayerBet(r.Players[i].ID),
			r.Bank.GetPlayerWallet(r.Players[i].ID),
			r.Bank.GetPot(),
		))

		success = true
		r.bigBlindIndex = i
	}

	return nil
}
