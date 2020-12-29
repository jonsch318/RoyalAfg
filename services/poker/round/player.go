package round

import (
	"errors"
	"fmt"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/events"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/models"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/utils"
	"log"
	"strconv"
	"time"
)

func (h *Round) actions(preflop bool) {

	var startIndexPlayers int
	for j := 1; j <= len(h.Players); j++ {
		startIndexPlayers = (h.bigBlindIndex + j) % len(h.Players)
		if h.Players[startIndexPlayers].Active {
			break
		}
	}

	startIndexBlocking := -1
	blocking := make([]int, 0)
	for i, n := range h.Players {
		if n.Active && !h.Bank.IsAllIn(n.ID) {
			blocking = append(blocking, i)
			if startIndexPlayers == i {
				startIndexBlocking = i
			}
		}
	}

	if(len(blocking) < 0){
		//Everything is handled no further actions from players necessary
		return
	}

	h.recAction(blocking, startIndexBlocking%len(blocking), preflop, true, 0)
}

func (h *Round) recAction(blocking []int, i int, preflop, canCheck bool, checkCount byte) {

	log.Printf("Checkcount: %v; InCount: %v", checkCount, h.InCount)
	if canCheck && checkCount >= h.InCount {
		canCheck = false
	}

	if h.InCount < 2 {
		return
	}

	removed := false
	// Check if blocking is an empty list
	if checkIfEmpty(blocking) {
		return
	}

	k := blocking[i]

	if k < 0 || !h.Players[k].Active || h.Bank.IsAllIn(h.Players[k].ID) {
		// remove from blocking list
		h.recAction(blocking, (i+1)%len(blocking), preflop, canCheck, checkCount)
		return
	}

	payload := 0
	var succeededAction events.Action
	success := false
	for j := 3; j > 0; j-- {
		a, err := h.waitForAction(k, preflop, canCheck)
		if err != nil {
			h.playerError(i, fmt.Sprintf("The action was not valid. %v more tries", j))
			continue
		}

		succeededAction = a
		if a.Action == events.FOLD {
			h.Fold(h.Players[k].ID)
			blocking = removeBlocking(blocking, i)
			removed = true
			success = true
			succeededAction = events.Action{
				Action:  events.FOLD,
				Payload: a.Payload,
			}
			break
		}

		if !preflop && a.Action == events.CHECK {
			checkCount++
			success = true
			succeededAction = events.Action{
				Action:  events.CHECK,
				Payload: a.Payload,
			}
			if err == nil {
				addBlocking(blocking, k)
				break
			}
		}

		if a.Action == events.RAISE {
			max := h.Bank.GetMaxBet()
			log.Printf("Raised to [%v] > [%v]", a.Payload, max)
			if a.Payload > max {
				amount := a.Payload
				err := h.Bank.Bet(h.Players[k].ID, amount)
				if err == nil {
					success = true
					succeededAction = events.Action{
						Action:  events.RAISE,
						Payload: amount,
					}
					payload = amount
					blocking = addAllButThisBlockgin(blocking, h.Players, k, h.Bank)
					removed = true
					break
				}
				h.playerError(i, fmt.Sprintf("Raise must be higher than the highest bet. %v more tries", j))
			}
		}

		if a.Action == events.BET || a.Action == events.ALL_IN {
			bet := h.Bank.GetMaxBet()
			if a.Action == events.ALL_IN {
				bet = h.Bank.AllIn(h.Players[k].ID)
			}
			err := h.Bank.Bet(h.Players[k].ID, bet)
			if err == nil {
				success = true
				succeededAction = events.Action{
					Action:  a.Action,
					Payload: a.Payload,
				}
				blocking = removeBlocking(blocking, i)
				removed = true
				payload = bet
				break
			}
			h.playerError(i, fmt.Sprintf("Bet must be equal to the current highest bet. %v more tries", j))
		}
	}

	if !success {
		succeededAction = events.Action{
			Action:  events.FOLD,
			Payload: 0,
		}
		h.Fold(h.Players[k].ID)
		removeBlocking(blocking, i)
		removed = true
	}

	utils.SendToAll(h.Players, events.NewActionProcessedEvent(succeededAction.Action, payload, k, h.Bank.GetPlayerBet(h.Players[k].ID), h.Bank.GetPlayerWallet(h.Players[i].ID)))

	time.Sleep(1 * time.Second)

	if !checkIfEmpty(blocking) {

		next := (i + 1) % len(blocking)
		if removed {
			next = i % len(blocking)
		}
		// blocking has changed now so the length is different and the
		h.recAction(blocking, next, preflop, canCheck, checkCount)
	}

	return

}

func (h *Round) Fold(id string) error {
	i, err := h.searchByActiveID(id)

	if err != nil {
		return err
	}

	if i < 0 || i >= len(h.Players) {
		return errors.New("Something went wrong")
	}
	h.Players[i].Active = false
	h.InCount--
	utils.SendToAll(h.Players, events.NewActionProcessedEvent(events.FOLD, 0, i, h.Bank.GetPlayerBet(h.Players[i].ID), h.Bank.GetPlayerWallet(h.Players[i].ID)))
	return nil
}

func (h *Round) playerError(i int, message string) {
	utils.SendToPlayerInList(h.Players, i, models.NewEvent("INVALID_ACTION", message))
}

func (h *Round) waitForAction(i int, preflop, check bool) (events.Action, error) {

	var possibilities byte
	possibilities = 0b11111
	mustAllIn, err := h.Bank.MustAllIn(h.Players[i].ID)
	if err != nil {
		return events.Action{}, errors.New("The player was not found in the bank")
	}

	//has bet all
	if mustAllIn {
		log.Printf("Must All In")
		possibilities = possibilities & 0b11001
	}
	//cant check?
	if preflop || !check {
		log.Printf("Player cant check")

		possibilities = possibilities & 0b10111
	}

	log.Printf("Possibilities:  %v", strconv.FormatInt(int64(possibilities), 2))

	utils.SendToAll(h.Players, events.NewWaitForActionEvent(i, possibilities))

	e, err := utils.WaitUntilEvent(&h.Players[i])
	if err != nil {
		log.Printf("Timeout: %v", err)
		return events.Action{}, err
	}
	action, err := events.ToAction(e)
	if err != nil {
		log.Printf("Decoding err: %v", err)
		return events.Action{}, err
	}
	return *action, nil
}

func (h *Round) PlayerLeaves(id string) error {
	_, i, err := utils.SearchByID(h.Players, id)
	if err != nil {
		return err
	}
	err = h.Fold(id)
	if err != nil {
		return err
	}
	utils.SendToAll(h.Players, events.NewActionProcessedEvent(events.FOLD, 0, i, h.Bank.GetPlayerBet(h.Players[i].ID), h.Bank.GetPlayerWallet(h.Players[i].ID)))
	if len(h.Players) < 2 {
		h.End()
	}
	return nil
}
