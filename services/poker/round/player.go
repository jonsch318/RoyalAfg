package round

import (
	"errors"
	"strconv"

	"github.com/jonsch318/royalafg/pkg/log"
	"github.com/jonsch318/royalafg/services/poker/events"
	"github.com/jonsch318/royalafg/services/poker/models"
	"github.com/jonsch318/royalafg/services/poker/utils"
)

func (r *Round) Leave(id string) error {

	i, err := r.searchByActiveID(id)
	if err != nil {
		log.Logger.Errorw("Leave called for player that does not exist.", "error", err)
		return err
	}

	log.Logger.Warnf("Player [%v][%v] is leaving", id, r.Players[i].Username)
	r.Players[i].Left = true
	if !r.Players[i].Active {
		//We dont have to fold an inactive player
		return nil
	}

	if r.InCount-1 <= 0 {
		r.end()
	}
	return r.fold(id)
}

func (r *Round) playerError(i int, message string) {
	_ = utils.SendToPlayerInListTimeout(r.Players, i, models.NewEvent("INVALID_ACTION", message))
}

// waitForAction sends the player all action possibilities in a decoded form and awaits an response.
func (r *Round) waitForAction(i int, preFlop, check bool) (*events.Action, error) {

	var possibilities byte
	possibilities = 0b11111
	mustAllIn, err := r.Bank.MustAllIn(r.Players[i].ID)
	if err != nil {
		return nil, errors.New("the player was not found in the bank")
	}

	if mustAllIn {
		//player cant bet or raise so we disable these flags
		possibilities = possibilities & 0b11001
	}
	if preFlop || !check {
		//player cant check so we disable the flag
		possibilities = possibilities & 0b10111
	}

	log.Logger.Debugf("Action possibilities:  %v", strconv.FormatInt(int64(possibilities), 2))

	utils.SendToAll(r.Players, events.NewWaitForActionEvent(&r.PublicPlayers[i], i, possibilities))

	e, err := utils.WaitUntilCloseOrEvent(&r.Players[i])
	if err != nil {
		log.Logger.Warnw("Timeout waiting for action or player left", "error", err)
		return nil, err
	}
	action, err := events.ToAction(e)
	if err != nil {
		log.Logger.Warn("Error decoding action", "error", err)
		return nil, err
	}

	log.Logger.Debugw("Received action", "action", action.Action, "payload", action.Payload.Display())
	return action, nil
}
