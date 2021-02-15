package round

import (
	"errors"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/log"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/events"
	moneyUtils "github.com/JohnnyS318/RoyalAfgInGo/services/poker/money"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/utils"
)

//actions is the starter for one round of actions (til all players have equal bets or folded)
func (r *Round) actions(preFlop bool) {

	//actions are quite complicated to pull of right.
	//We start defining a list with all the players that are blocking the program to go on. If one bets we remove him from the list
	//If one raises we remove him and put everyone else on the list again.
	//We try to be close to the rules, so we start next to the dealer.

	//search the closest to the dealer position.
	var startIndexPlayers int
	for j := 1; j <= len(r.Players); j++ {
		startIndexPlayers = (r.bigBlindIndex + j) % len(r.Players)
		if r.Players[startIndexPlayers].Active {
			break
		}
	}

	log.Logger.Debugf("Start index found %v", startIndexPlayers)

	//we have found our starting point.
	//generate the blocking list with all players that are active and everybody that is not already all in.
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

	log.Logger.Debugf("blocking list defined")

	if len(blocking) < 0{
		//Everything is handled no further actions from players necessary
		return
	}

	//Define state for the recursive action function.
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

	log.Logger.Debug("action round state set")

	r.RecursiveAction(roundState)
}

//Fold removes the player with the given id from the active player list
func (r *Round) Fold(id string) error {
	i, err := r.searchByActiveID(id)
	if err != nil {
		return err
	}

	if i < 0 || i >= len(r.Players) {
		return errors.New("something went wrong")
	}
	r.Players[i].Active = false
	r.InCount--

	log.Logger.Debugf("set active state of player")

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