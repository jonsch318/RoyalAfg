package aggregates

import (
	"errors"

	ycq "github.com/jetbasrawi/go.cqrs"

	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/pkg/events"
)

type Account struct {
	*ycq.AggregateBase
	UserId  string
	Balance int
}

func NewAccount(id string) *Account {
	return &Account{
		AggregateBase: ycq.NewAggregateBase(id),
		Balance:       0,
	}
}

func (a *Account) Create(userId string) error {
	if userId == "" {
		return errors.New("the userid can not be empty")
	}
	a.Apply(ycq.NewEventMessage(a.AggregateID(), &events.AccountCreated{
		a.AggregateID(), userId}, ycq.Int(a.CurrentVersion())), true)

	return nil
}

func (a Account) Purchase(amount int, gameID string, details interface{}) error {

	if amount <= 0 {
		return errors.New("A purchase must contain some amount greater than 0")
	}

	a.Apply(
		ycq.NewEventMessage(a.AggregateID(),
			&events.Purchased{ID: a.AggregateID(), Amount: amount, GameID: gameID, Details: details}, ycq.Int(a.CurrentVersion())),
		true)

	return nil
}

func (a Account) Deposit(amount int) error {

	if amount <= 0 {
		return errors.New("A deposit must contain some amount greater than 0")
	}

	a.Apply(
		ycq.NewEventMessage(a.AggregateID(),
			&events.Deposited{ID: a.AggregateID(), Amount: amount}, ycq.Int(a.CurrentVersion())),
		true)

	return nil
}

func (a Account) Apply(message ycq.EventMessage, isNew bool) {
	if isNew {
		a.TrackChange(message)
	}
	switch e := message.Event().(type) {
	case *events.AccountCreated:
		a.UserId = e.UserId
	case *events.Deposited:
		a.Balance += e.Amount
	case *events.Purchased:
		a.Balance -= e.Amount
	}
}
