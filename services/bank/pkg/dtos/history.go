package dtos

import (
	"fmt"
	"net/url"
	"reflect"
	"time"

	ycq "github.com/jetbasrawi/go.cqrs"
	goes "github.com/jetbasrawi/go.geteventstore"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/dtos"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/log"
	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/pkg/domain/aggregates"
	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/pkg/events"
	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/pkg/repositories"
)

type AccountHistoryEvent struct {
	Amount *dtos.CurrencyDto `json:"amount"`
	Type   string            `json:"type"`
	Time   time.Time         `json:"time"`
	Game string `json:"gameId"`
	LobbyID string `json:"roundId"`
}

type AccountHistoryQuery struct {
	accounts map[string][]AccountHistoryEvent
	repo     *repositories.Account
	client   *goes.Client
}

func NewAccountHistoryQuery(repo *repositories.Account, client *goes.Client) *AccountHistoryQuery {
	return &AccountHistoryQuery{
		accounts: make(map[string][]AccountHistoryEvent),
		repo:     repo,
		client:   client,
	}
}

func (q *AccountHistoryQuery) Handle(message ycq.EventMessage) {

	switch ev := message.Event().(type) {
	case *events.AccountCreated:
		q.accounts[message.AggregateID()] = make([]AccountHistoryEvent, 0)
	case *events.Deposited:
		q.accounts[message.AggregateID()] = append(q.accounts[message.AggregateID()], AccountHistoryEvent{
			Amount:  dtos.FromMoney(ev.Amount),
			Type:    "Deposited",
			Time:    time.Now(),
			Game:    ev.GameId,
			LobbyID: ev.RoundId,
		})
	case *events.Withdrawn:
		q.accounts[message.AggregateID()] = append(q.accounts[message.AggregateID()], AccountHistoryEvent{
			Amount:  dtos.FromMoney(ev.Amount.Multiply(-1)),
			Type:    "Withdrawn",
			Time:    time.Now(),
			Game:    ev.GameId,
			LobbyID: ev.RoundId,
		})
	}
}

//GetAccountHistory queries the read model for the l recorded history from the given index, where 0 is newest and last is oldest.
func (q *AccountHistoryQuery) GetAccountHistory(id string, i int, l int) ([]AccountHistoryEvent, error) {
	fullHistory, ok := q.accounts[id]
	if !ok {
		//Load
		err := q.LoadAggregate(id)
		if err != nil {
			return nil, err
		}
		fullHistory = q.accounts[id]
	}

	if len(fullHistory) == 0 {
		return nil, nil
	}

	if len(fullHistory) < l+i {
		//the query can not be fully answered, so we do our best to satisfy it as much as possible
		if len(fullHistory) <= l {
			//smaller than our result length. just use fullHistory
			return fullHistory, nil
		}

		//return l from last.
		return fullHistory[len(fullHistory)-l:], nil
	}

	return fullHistory[i : i+l], nil
}

func (q *AccountHistoryQuery) LoadAggregate(id string) error {
	aggregateType := reflect.TypeOf(&aggregates.Account{}).Elem().Name()
	streamName, _ := q.repo.StreamNameDelegate.GetStreamName(aggregateType, id)
	log.Logger.Debugf("Stream name %v", streamName)
	stream := q.client.NewStreamReader(streamName)
	for stream.Next() {
		switch err := stream.Err().(type) {
		case nil:
			break
		case *url.Error, *goes.ErrTemporarilyUnavailable:
			return fmt.Errorf("client not available")
		case *goes.ErrNoMoreEvents:
			return nil
		case *goes.ErrUnauthorized:
			return fmt.Errorf("unauthorized")
		case *goes.ErrNotFound:
			return fmt.Errorf("account with the id %v could not be found", id)
		default:
			return fmt.Errorf("unexpected err %v", err.Error())
		}

		event := q.repo.EventDelegate.GetEvent(stream.EventResponse().Event.EventType)

		//TODO: No test for meta
		meta := make(map[string]string)
		_ = stream.Scan(event, &meta)
		if stream.Err() != nil {
			return stream.Err()
		}
		em := ycq.NewEventMessage(id, event, ycq.Int(stream.EventResponse().Event.EventNumber))
		for k, v := range meta {
			em.SetHeader(k, v)
		}
		q.Handle(em)
	}
	return nil
}