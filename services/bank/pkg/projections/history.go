package projections

import (
	"fmt"
	"net/url"
	"reflect"
	"time"

	ycq "github.com/jetbasrawi/go.cqrs"
	goes "github.com/jetbasrawi/go.geteventstore"

	"github.com/jonsch318/royalafg/pkg/dtos"
	"github.com/jonsch318/royalafg/pkg/log"
	"github.com/jonsch318/royalafg/services/bank/pkg/aggregates"
	"github.com/jonsch318/royalafg/services/bank/pkg/events"
	"github.com/jonsch318/royalafg/services/bank/pkg/repositories"
)

// AccountHistoryQuery is the projection for the history
type AccountHistoryQuery struct {
	accounts map[string][]dtos.AccountHistoryEvent
	repo     *repositories.Account
	client   *goes.Client
}

func NewAccountHistoryQuery(repo *repositories.Account, client *goes.Client) *AccountHistoryQuery {
	return &AccountHistoryQuery{
		accounts: make(map[string][]dtos.AccountHistoryEvent),
		repo:     repo,
		client:   client,
	}
}

// Handle projects the new events to the read model
func (q *AccountHistoryQuery) Handle(message ycq.EventMessage) {

	switch ev := message.Event().(type) {
	case *events.AccountCreated:
		q.accounts[message.AggregateID()] = make([]dtos.AccountHistoryEvent, 0)
	case *events.Deposited:
		q.accounts[message.AggregateID()] = append(q.accounts[message.AggregateID()], dtos.AccountHistoryEvent{
			Amount:  dtos.FromMoney(ev.Amount),
			Type:    "Deposited",
			Time:    time.Now(),
			Game:    ev.GameId,
			LobbyID: ev.RoundId,
		})
	case *events.Withdrawn:
		q.accounts[message.AggregateID()] = append(q.accounts[message.AggregateID()], dtos.AccountHistoryEvent{
			Amount:  dtos.FromMoney(ev.Amount),
			Type:    "Withdrawn",
			Time:    time.Now(),
			Game:    ev.GameId,
			LobbyID: ev.RoundId,
		})
	case *events.Backroll:
		q.accounts[message.AggregateID()] = append(q.accounts[message.AggregateID()], dtos.AccountHistoryEvent{
			Amount:  dtos.FromMoney(ev.Amount),
			Type:    "Backroll",
			Time:    time.Now(),
			Game:    ev.GameId,
			LobbyID: ev.RoundId,
		})
	}
}

func (q *AccountHistoryQuery) GetPreviousEvent(id string) (dtos.AccountHistoryEvent, error) {
	fullHistory, ok := q.accounts[id]
	if !ok {
		err := q.LoadAggregate(id)
		if err != nil {
			return dtos.AccountHistoryEvent{}, err
		}
		fullHistory = q.accounts[id]
	}

	if len(fullHistory) == 0 {
		return dtos.AccountHistoryEvent{}, fmt.Errorf("no previous event found")
	}

	return fullHistory[len(fullHistory)-1], nil

}

// GetAccountHistory queries the read model for the l recorded history from the given index, where 0 is newest and last is oldest.
func (q *AccountHistoryQuery) GetAccountHistory(id string, i int, l int) ([]dtos.AccountHistoryEvent, error) {
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

// Load Aggregate loads the aggregate from the event store. This would be unnecessary because you would save the commands to a separate database
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
