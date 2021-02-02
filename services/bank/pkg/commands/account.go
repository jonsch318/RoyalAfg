package commands

import (
	"time"
)

type CreateBankAccount struct {
}

type Deposit struct {
	Amount int
	//Identification of Game
	GameId string
	//Identification in which round
	RoundId string
	Time    time.Time
}

type Withdraw struct {
	Amount int
	//Identification of Game
	GameId string
	//Identification in which round
	RoundId string
	Time    time.Time
}

type DeleteBankAccount struct {}