package bank

import (
	"bytes"
	"encoding/json"
	"time"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/bank"
)


func (b *Bank) PublishCommand(command bank.Command) error {

	//Fill all fields.

	command.Lobby = b.LobbyId
	command.Game = "Poker"

	var buf *bytes.Buffer
	err := json.NewEncoder(buf).Encode(&command)
	if err != nil {
		return err
	}

	err = b.eventBus.PublishCommand(command.CommandType, buf.Bytes())

	if err != nil {
		return err
	}

	return nil
}

//Publishes all commands queued. Usually at the end of the game
func (b *Bank) ExecuteQueue()  {
	//compressed commands evaluate the end difference after the game (e.g. -5 -5 -5 +10 => -5). This reduces traffic and load on the bank.
	//in case of a crash the entire poker game will not persisted to the bank. The bank will be updated once a round. Buy In is an extra
	compressed := make(map[string]bank.Command, len(b.PlayerWallet))
	for _, command := range b.eventQueue {
		v, ok := compressed[command.UserId]
		if !ok {
			compressed[command.UserId] = command
		}else {
			switch command.CommandType {
			case bank.Withdraw:
				v.Amount -= command.Amount
			case bank.Deposit:
				v.Amount += command.Amount
			}
		}
	}

	//Publish compressed commands
	for _, command := range compressed {
		command.Time = time.Now()
		b.PublishCommand(command)
	}
}

func (b *Bank) AddBetEvent(userId string, amount int)  {
	b.eventQueue = append(b.eventQueue, *bank.NewCommand(bank.Withdraw, userId, amount, "Poker", b.LobbyId))
}

func (b *Bank) AddWinEvent(userId string, amount int){
	b.eventQueue = append(b.eventQueue, *bank.NewCommand(bank.Deposit, userId, amount, "Poker", b.LobbyId))
}