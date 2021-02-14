package round

import (
	"errors"
	"log"
	"strconv"

	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/events"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/models"
	moneyUtils "github.com/JohnnyS318/RoyalAfgInGo/services/poker/money"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/utils"
)



func (r *Round) playerError(i int, message string) {
	utils.SendToPlayerInList(r.Players, i, models.NewEvent("INVALID_ACTION", message))
}

func (r *Round) waitForAction(i int, preFlop, check bool) (*events.Action, error) {

	var possibilities byte
	possibilities = 0b11111
	mustAllIn, err := r.Bank.MustAllIn(r.Players[i].ID)
	if err != nil {
		return nil, errors.New("the player was not found in the bank")
	}

	//has bet all
	if mustAllIn {
		log.Printf("Must All In")
		possibilities = possibilities & 0b11001
	}
	//cant check?
	if preFlop || !check {
		log.Printf("Player cant check")

		possibilities = possibilities & 0b10111
	}

	log.Printf("Possibilities:  %v", strconv.FormatInt(int64(possibilities), 2))

	utils.SendToAll(r.Players, events.NewWaitForActionEvent(i, possibilities))

	e, err := utils.WaitUntilEvent(&r.Players[i])
	if err != nil {
		log.Printf("Timeout: %v", err)
		return nil, err
	}
	action, err := events.ToAction(e)
	if err != nil {
		log.Printf("Decoding err: %v", err)
		return nil, err
	}
	return action, nil
}

func (r *Round) PlayerLeaves(id string) error {
	_, i, err := utils.SearchByID(r.Players, id)
	if err != nil {
		return err
	}
	err = r.Fold(id)
	if err != nil {
		return err
	}
	utils.SendToAll(r.Players, events.NewActionProcessedEvent(events.FOLD, i, moneyUtils.Zero().Display(), r.Bank.GetPlayerBet(r.Players[i].ID), r.Bank.GetPlayerWallet(r.Players[i].ID)))
	if len(r.Players) < 2 {
		r.End()
	}
	return nil
}
