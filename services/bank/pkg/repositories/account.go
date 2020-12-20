package repositories

import (
	"fmt"
	"reflect"

	ycq "github.com/jetbasrawi/go.cqrs"
	goes "github.com/jetbasrawi/go.geteventstore"

	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/pkg/aggregates"
	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/pkg/events"
)

type AccountRepo struct {
	repo *ycq.GetEventStoreCommonDomainRepo
}

func NewAccountRepo(eventStore *goes.Client, eventBus ycq.EventBus) (*AccountRepo, error) {
	r, err := ycq.NewCommonDomainRepository(eventStore, eventBus)
	if err != nil {
		return nil, err
	}

	repo := &AccountRepo{repo: r}

	aggregateFactory := ycq.NewDelegateAggregateFactory()
	aggregateFactory.RegisterDelegate(&aggregates.Account{}, func(id string) ycq.AggregateRoot {
		return aggregates.NewAccount(id)
	})
	repo.repo.SetAggregateFactory(aggregateFactory)

	streamNameDelegate := ycq.NewDelegateStreamNamer()
	streamNameDelegate.RegisterDelegate(func(t string, id string) string {
		return fmt.Sprintf("%v-%v", t, id)
	}, &aggregates.Account{})
	repo.repo.SetStreamNameDelegate(streamNameDelegate)

	eventFactory := ycq.NewDelegateEventFactory()
	eventFactory.RegisterDelegate(&events.AccountCreated{}, func() interface{} {
		return &events.AccountCreated{}
	})
	eventFactory.RegisterDelegate(&events.Purchased{}, func() interface{} {
		return &events.Deposited{}
	})
	eventFactory.RegisterDelegate(&events.Deposited{}, func() interface{} {
		return &events.Purchased{}
	})
	repo.repo.SetEventFactory(eventFactory)
	return repo, nil
}

func (r AccountRepo) Load(aggregateType, id string) (*aggregates.Account, error) {
	aggregate, err := r.repo.Load(reflect.TypeOf(&aggregates.Account{}).Elem().Name(), id)
	if _, ok := err.(*ycq.ErrAggregateNotFound); ok {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	if result, ok := aggregate.(*aggregates.Account); ok {
		return result, nil
	}

	return nil, fmt.Errorf("could not cast aggregate returned to type of %s", reflect.TypeOf(&aggregates.Account{}).Elem().Name())
}

func (r AccountRepo) Save(aggregate ycq.AggregateRoot, expectedVersion *int) error {
	return r.repo.Save(aggregate, expectedVersion)
}
