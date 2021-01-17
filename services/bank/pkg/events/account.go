package events

import (
	"time"
)

type AccountCreated struct {
	ID string
	Balance int
}

type Withdrawn struct {
	ID      string
	Amount  int
	GameId  string
	RoundId string
	Time    time.Time
}

type Deposited struct {
	ID      string
	Amount  int
	GameId  string
	RoundId string
	Time    time.Time
}