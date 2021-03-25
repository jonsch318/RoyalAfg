package round

import (
	"errors"
	"strconv"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/log"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/events"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/models"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/utils"
)



func (r *Round) playerError(i int, message string) {
	_ = utils.SendToPlayerInListTimeout(r.Players, i, models.NewEvent("INVALID_ACTION", message))
}

//waitForAction sends the player all action possibilities in a decoded form and awaits an response.
func (r *Round) waitForAction(i int, preFlop, check bool) (*events.Action, error) {

	var possibilities byte
	possibilities = 0b11111
	mustAllIn, err := r.Bank.MustAllIn(r.Players[i].ID)
	if err != nil {
		return nil, errors.New("the player was not found in the bank")
	}

	if mustAllIn {
		log.Logger.Debug("Must All In")
		//player cant bet or raise so we disable these flags
		possibilities = possibilities & 0b11001
	}
	if preFlop || !check {
		log.Logger.Debug("Player cant check")
		//player cant check so we disable the flag
		possibilities = possibilities & 0b10111
	}

	log.Logger.Infof("Possibilities:  %v", strconv.FormatInt(int64(possibilities), 2))

	utils.SendToAll(r.Players, events.NewWaitForActionEvent(&r.PublicPlayers[i],i, possibilities))

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

	log.Logger.Debugw("Received action", "action", action.Action, "payload", action.Payload.Display() )
	return action, nil
}
