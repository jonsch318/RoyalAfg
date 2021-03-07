package bank

import (
	"time"

	"github.com/Rhymond/go-money"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/dtos"
)

type Command struct {
	UserId      string    `json:"userId"`
	CommandType string    `json:"type"`
	Amount      *dtos.CurrencyDto       `json:"amount"`
	Time        time.Time `json:"time"`
	Game        string    `json:"game"`
	Lobby       string    `json:"lobby"`
}

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
