package aggregates

import (
	"errors"
	"fmt"
	"log"
	"time"

	ycq "github.com/jetbasrawi/go.cqrs"

	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/pkg/events"
)

type Account struct {
	*ycq.AggregateBase
	Balance int
}

func NewAccount(id string) *Account{
	return &Account{
		AggregateBase: ycq.NewAggregateBase(id),
		Balance: 0,
	}
}

func (a *Account) Create() error {
	a.Apply(ycq.NewEventMessage(a.AggregateID(), &events.AccountCreated{
		ID: a.AggregateID(),
	}, ycq.Int(a.CurrentVersion())), true)

	log.Printf("Create Account %v", a.AggregateID())
	return nil
}

func (a *Account) Deposit(amount int, gameId, roundId string, time time.Time) error {
	if amount <= 0 {
		return errors.New("the amount has to be greater than 0")
	}
	a.Apply(ycq.NewEventMessage(a.AggregateID(), &events.Deposited{
		ID:      a.AggregateID(),
		Amount:  amount,
		GameId:  gameId,
		RoundId: roundId,
		Time:    time,
	}, ycq.Int(a.CurrentVersion())), true)
	log.Printf("Deposit to %v", a.AggregateID())

	return nil
}

func (a *Account) Withdraw(amount int, gameId, roundId string, time time.Time) error {


	if amount <= 0 {
		return errors.New("the amount which is to withdraw has to be greater than 0")
	}

	if amount > a.Balance{
		return fmt.Errorf("the user cannot withdraw the given amount [%v] > [%v]", amount, a.Balance)
	}

	a.Apply(ycq.NewEventMessage(a.AggregateID(), &events.Withdrawn{
		ID:      a.AggregateID(),
		Amount:  amount,
		GameId:  gameId,
		RoundId: roundId,
		Time:    time,
	}, ycq.Int(a.CurrentVersion())), true)

	log.Printf("Withdraw from %v", a.AggregateID())

	return nil
}

func (a *Account) Apply(message ycq.EventMessage, isNew bool) {
	if isNew {
		a.TrackChange(message)
	}

	switch ev := message.Event().(type) {
	case *events.AccountCreated:
		a.Balance = 0
	case *events.Deposited:
		a.Balance = a.Balance +  ev.Amount
	case *events.Withdrawn:
		a.Balance = a.Balance - ev.Amount

	}
}