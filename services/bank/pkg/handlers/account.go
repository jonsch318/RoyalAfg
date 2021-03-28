package handlers

import (
	ycq "github.com/jetbasrawi/go.cqrs"

	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/pkg/projections"
)

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
