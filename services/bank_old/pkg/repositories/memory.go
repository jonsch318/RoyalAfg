package repositories

import (
	"log"

	ycq "github.com/jetbasrawi/go.cqrs"

	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/pkg/aggregates"
)

//InMemoryAccount is a in memory repository for bank accounts used for testing
type InMemoryAccount struct {
	current   map[string][]ycq.EventMessage
	publisher ycq.EventBus
}

func NewInMemoryAccount(eventbus ycq.EventBus) *InMemoryAccount {
	return &InMemoryAccount{
		publisher: eventbus,
		current:   make(map[string][]ycq.EventMessage),
	}
}

// Load loads an aggregate of the specified type.
func (r *InMemoryAccount) Load(aggregateType, id string) (*aggregates.Account, error) {

	events, ok := r.current[id]
	if !ok {
		return nil, &ycq.ErrAggregateNotFound{}
	}

	inventoryItem := aggregates.NewAccount(id)

	for _, v := range events {
		inventoryItem.Apply(v, false)
		inventoryItem.IncrementVersion()
	}
	log.Printf("replayed account state")

	return inventoryItem, nil
}

// Save persists an aggregate.
func (r *InMemoryAccount) Save(aggregate ycq.AggregateRoot, _ *int) error {

	//TODO: Look at the expected version
	for _, v := range aggregate.GetChanges() {
		r.current[aggregate.AggregateID()] = append(r.current[aggregate.AggregateID()], v)
		r.publisher.PublishEvent(v)
	}

	return nil

}
