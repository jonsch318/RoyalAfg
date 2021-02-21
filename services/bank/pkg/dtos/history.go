package dtos

import (
	"errors"
	"time"

	ycq "github.com/jetbasrawi/go.cqrs"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/dtos"
	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/pkg/events"
)

type AccountHistoryEvent struct {
	Amount *dtos.CurrencyDto `json:"amount"`
	Type string `json:"type"`
	Time time.Time `json:"time"`
}

type AccountHistoryQuery struct {
	accounts map[string][]AccountHistoryEvent
}

func NewAccountHistoryQuery() *AccountHistoryQuery	 {
	return &AccountHistoryQuery{
		accounts: make(map[string][]AccountHistoryEvent),
	}
}

func (q *AccountHistoryQuery) Handle(message ycq.EventMessage) {

	switch ev := message.Event().(type) {
	case *events.AccountCreated:
		q.accounts[message.AggregateID()] = make([]AccountHistoryEvent,0)
	case *events.Deposited:
		q.accounts[message.AggregateID()] = append(q.accounts[message.AggregateID()], AccountHistoryEvent{
			Amount: dtos.FromMoney(ev.Amount),
			Type:   "Deposited",
			Time:   time.Now(),
		})
	case *events.Withdrawn:
		q.accounts[message.AggregateID()] = append(q.accounts[message.AggregateID()], AccountHistoryEvent{
			Amount: dtos.FromMoney(ev.Amount),
			Type:   "Withdrawn",
			Time:   time.Now(),
		})
	}
}

//GetAccountHistory queries the read model for the l recorded history from the given index, where 0 is newest and last is oldest.
func (q *AccountHistoryQuery) GetAccountHistory(id string, i int, l int ) ([]AccountHistoryEvent, error){
	fullHistory, ok := q.accounts[id]
	if !ok{
		return nil, errors.New("the account with the given id does not exist")
	}

	if len(fullHistory) == 0 {
		return nil, nil
	}

	if len(fullHistory) < l + i {
		//the query can not be fully answered, so we do our best to satisfy it as much as possible
		if len(fullHistory) <= l {
			//smaller than our result length. just use fullHistory
			return fullHistory, nil
		}

		//return l from last.
		return fullHistory[len(fullHistory) - l:], nil
	}

	return fullHistory[i:i+l], nil
}