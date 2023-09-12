package bank

import (
	"time"

	"github.com/Rhymond/go-money"

	"github.com/jonsch318/royalafg/pkg/dtos"
)

// Command is a default Bank Command used to communicate with the bank
type Command struct {
	UserId      string            `json:"userId"`
	CommandType string            `json:"type"`
	Amount      *dtos.CurrencyDto `json:"amount"`
	Time        time.Time         `json:"time"`
	Game        string            `json:"game"`
	Lobby       string            `json:"lobby"`
	Reason      string            `json:"reason"`
}

// NewCommand creates a new bank command
func NewCommand(commandType, userId string, amount *money.Money, game, lobby string) *Command {
	return &Command{
		UserId:      userId,
		CommandType: commandType,
		Amount:      dtos.FromMoney(amount),
		Time:        time.Now(),
		Game:        game,
		Lobby:       lobby,
	}
}
