package bank

import (
	"encoding/json"
	"time"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/bank"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/log"
)


func (b *Bank) publishCommand(command *bank.Command) error {

	//Fill all fields.
	command.Lobby = b.LobbyId
	command.Game = "Poker"

	buf, err := json.Marshal(command)
	if err != nil {
		log.Logger.Errorw("Command could not be encoded", "error", err)
		return err
	}

	err = b.eventBus.PublishCommand(command.CommandType, buf)

	if err != nil {
		log.Logger.Errorw("Command could not be added to the event bus", "error", err)
		return err
	}
	return nil
}

//executeQueue Publishes all commands queued. Usually at the end of the game
func (b *Bank) executeQueue()  {
	//This could be done via a command queue, that everything adds to. (Could include more information about games. But this is simpler.
	for k, val := range b.PlayerBets	 {
		if val.IsZero(){
			continue
		}
		cmd := new(bank.Command)
		if val.IsNegative() {
			// Win because player bets are the expenses of the player. We deposit the winning amount to the player
			res := val.Multiply(-1)
			cmd = bank.NewCommand(bank.Deposit, k, res, "Poker", b.LobbyId)
			cmd.Time = time.Now()
		} else {
			cmd = bank.NewCommand(bank.Withdraw, k, val, "Poker", b.LobbyId)
			cmd.Time = time.Now()
		}

		log.Logger.Debugw("Publishing command", "cmd", cmd)

		//Retry functionality should be implemented here. A crucial peace to send all events to the bank. Add redundancy
		_ = b.publishCommand(cmd)
	}
}