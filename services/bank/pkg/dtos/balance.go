package dtos

import (
	"errors"
	"log"

	ycq "github.com/jetbasrawi/go.cqrs"

	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/pkg/events"
)


type AccountBalanceQuery struct {
	accounts map[string]int
}

func NewAccountBalanceQuery() *AccountBalanceQuery  {
	return &AccountBalanceQuery{
		accounts: make(map[string]int),
	}
}

func (q *AccountBalanceQuery) Handle(message ycq.EventMessage)  {

	log.Printf("Read Model handle %v", message)

	switch ev := message.Event().(type) {
	case *events.AccountCreated:
		log.Printf("created [%v]", message.AggregateID())
		q.accounts[message.AggregateID()] = 0

	case *events.Deposited:
		a, err := q.GetAccountBalance(message.AggregateID())
		if err != nil {
			log.Printf("the account was not created")
		}
		log.Printf("Deposited [%v] %v", message.AggregateID(), ev.Amount)
		q.accounts[message.AggregateID()] = a + ev.Amount

	case *events.Withdrawn:
		a, err := q.GetAccountBalance(message.AggregateID())
		if err != nil{
			log.Printf("the account was not created")
		}
		log.Printf("Withdrawn [%v] %v", message.AggregateID(), ev.Amount)
		q.accounts[message.AggregateID()] = a - ev.Amount
	}

}

func (q *AccountBalanceQuery) GetAccountBalance(id string) (int, error) {
	res, ok := q.accounts[id]
	if !ok {
		return -1, errors.New("the account with the given id does not exist")
	}

	if res == -1 {
		return -1, errors.New("the account with the given id does not exist")
	}

	return res, nil
}