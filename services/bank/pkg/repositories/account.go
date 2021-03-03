package repositories

import (
	"errors"
	"log"
	"reflect"

	ycq "github.com/jetbasrawi/go.cqrs"
	goes "github.com/jetbasrawi/go.geteventstore"

	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/pkg/domain/aggregates"
	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/pkg/events"
)

type Account struct {
	repo               *ycq.GetEventStoreCommonDomainRepo
	EventDelegate      *ycq.DelegateEventFactory
	StreamNameDelegate *ycq.DelegateStreamNamer
	AggregateDelegate  *ycq.DelegateAggregateFactory
}

func NewAccount(eventStore *goes.Client, eventBus ycq.EventBus) (*Account, error) {
	r, err := ycq.NewCommonDomainRepository(eventStore, eventBus)
	if err != nil {
		log.Printf("Error during CommonDomainRepository: %v", err)
		return nil, err
	}

	ret := &Account{
		repo: r,
	}

	aggregateFactory := ycq.NewDelegateAggregateFactory()
	_ = aggregateFactory.RegisterDelegate(&aggregates.Account{}, func(id string) ycq.AggregateRoot { return aggregates.NewAccount(id) })
	ret.repo.SetAggregateFactory(aggregateFactory)
	ret.AggregateDelegate = aggregateFactory

	streamNameDelegate := ycq.NewDelegateStreamNamer()
	_ = streamNameDelegate.RegisterDelegate(func(t string, id string) string {
		return t + "-" + id
	}, &aggregates.Account{})
	ret.repo.SetStreamNameDelegate(streamNameDelegate)
	ret.StreamNameDelegate = streamNameDelegate

	eventFactory := ycq.NewDelegateEventFactory()
	_ = eventFactory.RegisterDelegate(&events.AccountCreated{}, func() interface{} { return &events.AccountCreated{} })
	_ = eventFactory.RegisterDelegate(&events.Deposited{}, func() interface{} { return &events.Deposited{} })
	_ = eventFactory.RegisterDelegate(&events.Withdrawn{}, func() interface{} { return &events.Withdrawn{} })
	ret.repo.SetEventFactory(eventFactory)
	ret.EventDelegate = eventFactory

	return ret, nil
}

func (r Account) Load(aggregateType, id string) (*aggregates.Account, error) {
	aggregate, err := r.repo.Load(reflect.TypeOf(&aggregates.Account{}).Elem().Name(), id)

	if err != nil {
		return nil, err
	}
	a, ok := aggregate.(*aggregates.Account)
	if !ok {
		return nil, errors.New("the loaded aggregate could not be casted to the correct type")
	}

	return a, nil
}

func (r Account) Save(aggregate ycq.AggregateRoot, expectedVersion *int) error {
	err := r.repo.Save(aggregate, expectedVersion)
	if err != nil {
		log.Printf("error during saving: %v", err)
		return err
	}
	return nil
}
