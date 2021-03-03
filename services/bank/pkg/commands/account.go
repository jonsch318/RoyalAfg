package commands

import (
	"time"

	"github.com/Rhymond/go-money"
)

type CreateBankAccount struct {
}

type Deposit struct {
	Amount *money.Money
	GameId string //Identification of Game
	RoundId string //Identification in which round
	Time    time.Time
}

type Withdraw struct {
	Amount *money.Money
	GameId string //Identification of Game
	RoundId string //Identification in which round
	Time    time.Time
}

type DeleteBankAccount struct {}