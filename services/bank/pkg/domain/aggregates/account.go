package aggregates

import (
	"errors"
	"fmt"
	"time"

	"github.com/Rhymond/go-money"
	ycq "github.com/jetbasrawi/go.cqrs"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/currency"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/log"
	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/pkg/events"
)

type Account struct {
	*ycq.AggregateBase
	Balance *money.Money
}

func NewAccount(id string) *Account {
	return &Account{
		AggregateBase: ycq.NewAggregateBase(id),
		Balance:       currency.Zero(),
	}
}

func (a *Account) Create() error {
	a.Apply(ycq.NewEventMessage(a.AggregateID(), &events.AccountCreated{
		ID: a.AggregateID(),
	}, ycq.Int(a.CurrentVersion())), true)

	log.Logger.Infof("Create Account %v", a.AggregateID())
	return nil
}

func (a *Account) Deposit(amount *money.Money, gameId, roundId string, time time.Time) error {
	if !amount.IsPositive() {
		return errors.New("the amount has to be greater than 0")
	}
	prev := a.Balance.Display()
	a.Apply(ycq.NewEventMessage(a.AggregateID(), &events.Deposited{
		ID:      a.AggregateID(),
		Amount:  amount,
		GameId:  gameId,
		RoundId: roundId,
		Time:    time,
	}, ycq.Int(a.CurrentVersion())), true)
	log.Logger.Debugw("Deposit operation success", "id", a.AggregateID(), "amount", amount.Display(), "total", a.Balance.Display(), "previousTotal", prev)

	return nil
}

func (a *Account) Withdraw(amount *money.Money, gameId, roundId string, time time.Time) error {
	if !amount.IsPositive() {
		return errors.New("the amount which is to withdraw has to be greater than 0")
	}
	prev := a.Balance.Display()
	if res, err := amount.GreaterThan(a.Balance); res || err != nil {
		if err != nil {
			log.Logger.Errorw("Error during comparison", "error", err)
		}
		return fmt.Errorf("the user cannot withdraw the given amount [%v] > [%v]", amount.Display(), a.Balance.Display())
	}

	a.Apply(ycq.NewEventMessage(a.AggregateID(), &events.Withdrawn{
		ID:      a.AggregateID(),
		Amount:  amount,
		GameId:  gameId,
		RoundId: roundId,
		Time:    time,
	}, ycq.Int(a.CurrentVersion())), true)

	log.Logger.Debugw("Withdraw operation success", "id", a.AggregateID(), "amount", amount.Display(), "total", a.Balance.Display(), "previousTotal", prev)

	return nil
}

func (a *Account) Apply(message ycq.EventMessage, isNew bool) {
	if isNew {
		a.TrackChange(message)
	}

	switch ev := message.Event().(type) {
	case *events.AccountCreated:
		a.Balance = currency.Zero()
	case *events.Deposited:
		res, err := a.Balance.Add(ev.Amount)
		if err != nil {
			log.Logger.Errorw("Error during money addition", "error", err)
			return
		}
		a.Balance = res
	case *events.Withdrawn:
		res, err := a.Balance.Subtract(ev.Amount)
		if err != nil {
			log.Logger.Errorw("Error during money addition", "error", err)
			return
		}
		a.Balance = res
	}
}
