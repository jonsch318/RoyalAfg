package events

import (
	"time"

	"github.com/Rhymond/go-money"
)

type AccountCreated struct {
	ID string
}

type Withdrawn struct {
	ID      string
	Amount  *money.Money
	GameId  string
	RoundId string
	Time    time.Time
}

type Deposited struct {
	ID      string
	Amount  *money.Money
	GameId  string
	RoundId string
	Time    time.Time
}