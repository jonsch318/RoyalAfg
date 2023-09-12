package handlers

import (
	ycq "github.com/jetbasrawi/go.cqrs"

	"github.com/jonsch318/royalafg/services/bank/pkg/projections"
)

// Account is the http handler for a bank account h
type Account struct {
	dispatcher       ycq.Dispatcher
	eventBus         ycq.EventBus
	balanceReadModel *projections.AccountBalanceQuery
	historyReadModel *projections.AccountHistoryQuery
}

func NewAccountHandler(dispatcher ycq.Dispatcher, eventBus ycq.EventBus, balanceQuery *projections.AccountBalanceQuery, historyReadModel *projections.AccountHistoryQuery) *Account {
	return &Account{
		dispatcher:       dispatcher,
		eventBus:         eventBus,
		balanceReadModel: balanceQuery,
		historyReadModel: historyReadModel,
	}
}
