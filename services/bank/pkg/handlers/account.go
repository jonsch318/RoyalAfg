package handlers

import (
	ycq "github.com/jetbasrawi/go.cqrs"

	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/pkg/dtos"
)

type Account struct {
	dispatcher ycq.Dispatcher
	eventBus ycq.EventBus
	balanceReadModel *dtos.AccountBalanceQuery
	historyReadModel *dtos.AccountHistoryQuery
}

func NewAccountHandler(dispatcher ycq.Dispatcher, eventBus ycq.EventBus, balanceQuery *dtos.AccountBalanceQuery, historyReadModel *dtos.AccountHistoryQuery) *Account {
	return &Account{
		dispatcher: dispatcher,
		eventBus:   eventBus,
		balanceReadModel: balanceQuery,
		historyReadModel: historyReadModel,
	}
}
