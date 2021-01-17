package bank

import (
	"time"
)

type Command struct {
	UserId      string    `json:"userId"`
	CommandType string    `json:"type"`
	Amount      int       `json:"amount"`
	Time        time.Time `json:"time"`
	Game        string    `json:"game"`
	Lobby       string    `json:"lobby"`
}

func NewCommand(commandType, userId string, amount int, game, lobby string) *Command {
	return &Command{
		UserId:      userId,
		CommandType: commandType,
		Amount:      amount,
		Time:        time.Now(),
		Game:        game,
		Lobby:       lobby,
	}
}
