package aggregates

import (
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/bank/system"
	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/events"
	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
)

type Account struct {
	//aggregate fields
	id      string
	version int
	changes []system.IEvent[any]

	//account fields
	balance    *money.Money
	holderId   string
	holderName string
	deleted    bool
}

func NewAccountFromEvents(events []system.IEvent[any]) *Account {
	a := &Account{}
	for _, e := range events {
		a.Apply(e)
	}
	return a
}

func (a *Account) GetId() string {
	return a.id
}

func (a *Account) GetVersion() int {
	return a.version
}

func (a *Account) GetChanges() []system.IEvent[any] {
	return a.changes
}

func (a *Account) ClearChanges() {
	a.changes = []system.IEvent[any]{}
}

func (a *Account) InitAccount() {
	a.id = uuid.New().String()
	a.version = 1
}

func (a *Account) Apply(event system.Event[any], isNew bool) {
	switch e := event.GetDat().(type) {
	//account creation
	case events.AccountCreated:
		a.holderId = e.HolderId
		a.holderName = e.HolderName

	//account deletion
	case events.AccountDeleted:
		a.deleted = true

	case events.Deposited:
		a.balance, _ = a.balance.Add(e.Amount)

	case events.Withdrawn:
		a.balance, _ = a.balance.Subtract(e.Amount)

	case events.Backroll:
		a.balance
	}
	if isNew {
		a.version++
	}
}
