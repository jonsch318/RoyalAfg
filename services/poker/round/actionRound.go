package round

import (
	"fmt"
	"log"
	"time"

	"github.com/Rhymond/go-money"

	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/events"
	moneyUtils "github.com/JohnnyS318/RoyalAfgInGo/services/poker/money"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/utils"
)

type ActionRoundOptions struct {
	Success            bool
	Payload            *money.Money
	SuccessfulAction *events.Action
	PlayerId string
	PreFlop bool
	CanCheck bool
	CheckCount byte
	Current int
	BlockingIndex int
	BlockingList *BlockingList
}

func (r *Round) RecursiveAction(options *ActionRoundOptions){

	//Checks in Poker are difficult to keep track. We count all the checks and decide whether more checks are legal.
	log.Printf("Checkcount: %v; InCount: %v", options.CheckCount, r.InCount)
	if options.CanCheck && options.CheckCount >= r.InCount {
		options.CanCheck = false
	}

	//Check if only one player is in the game or the blocking list is empty. (One player always wins) (Anchor)
	if r.InCount < 2 || options.BlockingList.CheckIfEmpty() {
		return
	}

	options.Current = options.BlockingList.Get(options.BlockingIndex)
	options.PlayerId = r.Players[options.Current].ID
	options.Success = false

	// Check for any abnormalities. Player must be active and has money to bet.
	if options.Current < 0 || !r.Players[options.Current].Active || r.Bank.IsAllIn(options.PlayerId) {
		// remove from blocking list
		options.BlockingList.RemoveBlocking(options.BlockingIndex)
		options.BlockingIndex = options.BlockingIndex % options.BlockingList.Length()
		r.RecursiveAction(options)
		return
	}

	r.ActionTries(options)

	if !options.Success {
		options.SuccessfulAction = &events.Action{
			Action:  events.FOLD,
			Payload: moneyUtils.Zero(),
		}
		_ = r.Fold(options.PlayerId)
		options.BlockingList.RemoveBlocking(options.BlockingIndex)
	}

	utils.SendToAll(r.Players,
		events.NewActionProcessedEvent(
			options.SuccessfulAction.Action,
			options.Current,
			options.Payload.Display() ,
			r.Bank.GetPlayerBet(options.PlayerId),
			r.Bank.GetPlayerWallet(options.PlayerId),
			),
	)
	time.Sleep(1 * time.Second)

	if !options.BlockingList.CheckIfEmpty() {
		//Get Next in Line
		next := options.BlockingList.GetNext(options.SuccessfulAction.Action != events.CHECK, options.BlockingIndex)

		options.BlockingIndex = next
		r.RecursiveAction(options)
	}

	return
}

func (r *Round) ActionTries(options *ActionRoundOptions)  {
	for i := 3; i > 0 ; i-- {
		action, err := r.waitForAction(options.Current, options.PreFlop, options.CanCheck)
		if err != nil || action == nil {
			r.playerError(options.BlockingIndex, fmt.Sprintf("The action was not valid. %v more tries", i))
			continue
		}
		options.SuccessfulAction = action
		options.Payload = action.Payload

		r.Action(options)

		if options.Success {
			break
		}
	}
}

func (r *Round) Action(options *ActionRoundOptions) {
	switch options.SuccessfulAction.Action {
	case events.FOLD:
		_ = r.Fold(options.PlayerId)
		options.BlockingList.RemoveBlocking(options.BlockingIndex)
		options.Success = true
		return

	case events.CHECK:
		if !options.PreFlop {
			options.CheckCount++
			err := options.BlockingList.AddBlocking(options.Current)
			if err == nil {
				options.Success = true
			}
			return
		}

	case events.ALL_IN:
		raise, err := r.Bank.PerformAllIn(options.PlayerId)
		if err == nil {
			options.Success = true
			if raise {
				options.BlockingList.AddAllButThisBlocking(r.Players, options.Current, r.Bank)
			} else {
				options.BlockingList.RemoveBlocking(options.BlockingIndex)
			}
		}
		return

	case events.RAISE:
		if r.Bank.IsRaise(options.Payload) {
			err := r.Bank.Bet(options.PlayerId, options.Payload)
			if err == nil {
				options.Success = true
				options.BlockingList.AddAllButThisBlocking(r.Players, options.Current, r.Bank)
				return
			}
			r.playerError(options.BlockingIndex, fmt.Sprintf("Raise must be higher than the highest bet."))
		}
		return

	case events.BET:
		err := r.Bank.PerformBet(options.PlayerId)
		if err == nil {
			options.Success = true
			options.BlockingList.RemoveBlocking(options.BlockingIndex)
			return
		}
		r.playerError(options.BlockingIndex, fmt.Sprintf("Bet must be equal to the current highest bet." ))
	}
}