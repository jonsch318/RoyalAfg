package round

import (
	"errors"

	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/events"
	moneyUtils "github.com/JohnnyS318/RoyalAfgInGo/services/poker/money"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/utils"
)

//actions is the starter for one round of actions (til all players have equal bets or folded)
func (r *Round) actions(preFlop bool) {

	var startIndexPlayers int
	for j := 1; j <= len(r.Players); j++ {
		startIndexPlayers = (r.bigBlindIndex + j) % len(r.Players)
		if r.Players[startIndexPlayers].Active {
			break
		}
	}

	startIndexBlocking := -1
	blocking := make([]int, 0)
	for i, n := range r.Players {
		if n.Active && !r.Bank.IsAllIn(n.ID) {
			blocking = append(blocking, i)
			if startIndexPlayers == i {
				startIndexBlocking = i
			}
		}
	}

	if len(blocking) < 0{
		//Everything is handled no further actions from players necessary
		return
	}

	roundState := &ActionRoundOptions{
		Success:          false,
		Payload:          moneyUtils.Zero(),
		SuccessfulAction: nil,
		PlayerId:         "",
		PreFlop:          preFlop,
		CanCheck:         true,
		CheckCount:       0,
		Current:          0,
		BlockingIndex:    startIndexBlocking%len(blocking),
		BlockingList:     NewBlockingList(blocking),
	}

	r.RecursiveAction(roundState)
}

func (r *Round) Fold(id string) error {
	i, err := r.searchByActiveID(id)

	if err != nil {
		return err
	}

	if i < 0 || i >= len(r.Players) {
		return errors.New("Something went wrong")
	}
	r.Players[i].Active = false
	r.InCount--
	utils.SendToAll(r.Players,
		events.NewActionProcessedEvent(
			events.FOLD,
			i,
			moneyUtils.Zero().Display(),
			r.Bank.GetPlayerBet(r.Players[i].ID),
			r.Bank.GetPlayerWallet(r.Players[i].ID),
			),
	)
	return nil
}