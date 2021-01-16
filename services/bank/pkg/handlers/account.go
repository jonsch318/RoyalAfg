package handlers

import (
	ycq "github.com/jetbasrawi/go.cqrs"

	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/pkg/dtos"
)

type Account struct {
	dispatcher ycq.Dispatcher
	eventBus ycq.EventBus
	balanceReadModel *dtos.AccountBalanceQuery
}

func NewAccountHandler(dispatcher ycq.Dispatcher, eventBus ycq.EventBus, balanceQuery *dtos.AccountBalanceQuery) *Account {
	return &Account{
		dispatcher: dispatcher,
		eventBus:   eventBus,
		balanceReadModel: balanceQuery,
	}
}
