package commands

import (
	"time"

	"github.com/Rhymond/go-money"
)

// CreateBankAccount is a command to create a new bank account
type CreateBankAccount struct {
}

// Deposit is a command to transact from a bank account
type Deposit struct {
	Amount  *money.Money
	GameId  string //Identification of Game
	RoundId string //Identification in which round
	Time    time.Time
}

// Withdraw is a command to transact to a bank account
type Withdraw struct {
	Amount  *money.Money
	GameId  string //Identification of Game
	RoundId string //Identification in which round
	Time    time.Time
}

type Rollback struct {
	Amount  *money.Money
	GameId  string //Identification of Game
	RoundId string //Identification in which round
	Time    time.Time
	Reason  string
}

type DeleteBankAccount struct{}
