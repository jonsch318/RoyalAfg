package dtos

import (
	"errors"
	"reflect"

	"github.com/Rhymond/go-money"
	ycq "github.com/jetbasrawi/go.cqrs"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/currency"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/log"
	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/pkg/domain/aggregates"
	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/pkg/events"
	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/pkg/repositories"
)

type AccountBalanceQuery struct {
	accounts map[string]*money.Money
	repo     *repositories.Account
}

func NewAccountBalanceQuery(repo *repositories.Account) *AccountBalanceQuery {
	return &AccountBalanceQuery{
		accounts: make(map[string]*money.Money),
		repo:     repo,
	}
}

func (q *AccountBalanceQuery) Handle(message ycq.EventMessage) {

	log.Logger.Debugf("Balance Read Model handle invoked")

	switch ev := message.Event().(type) {
	case *events.AccountCreated:
		log.Logger.Debugf("created [%v]", message.AggregateID())
		q.accounts[message.AggregateID()] = currency.Zero()

	case *events.Deposited:
		a, err := q.GetAccountBalance(message.AggregateID())
		if err != nil {
			log.Logger.Errorf("the account was not created")
			return
		}
		log.Logger.Debugf("Deposited [%v] %v", message.AggregateID(), ev.Amount.Display())
		res, err := a.Add(ev.Amount)
		if err != nil {
			log.Logger.Errorw("balance read model could not handle event", "error", err)
			return
		}
		q.accounts[message.AggregateID()] = res

	case *events.Withdrawn:
		a, err := q.GetAccountBalance(message.AggregateID())
		if err != nil {
			log.Logger.Debugf("the account was not created")
			return
		}
		log.Logger.Debugf("Withdrawn [%v] %v", message.AggregateID(), ev.Amount)
		res, err := a.Subtract(ev.Amount)
		if err != nil {
			log.Logger.Errorw("balance read model could not handle event", "error", err)
			return
		}
		q.accounts[message.AggregateID()] = res
	}

}

func (q *AccountBalanceQuery) GetAccountBalance(id string) (*money.Money, error) {
	res, ok := q.accounts[id]
	if !ok {
		item, err := q.repo.Load(reflect.TypeOf(&aggregates.Account{}).Elem().Name(), id)

		if err != nil {
			return nil, errors.New("the account with the given id does not exist")
		}

		log.Logger.Warnf("Loaded Unloaded Aggregate [%s]", id)
		q.accounts[id] = item.Balance
		return item.Balance, nil
	}

	return res, nil
}
