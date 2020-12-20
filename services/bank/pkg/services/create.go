package services

import (
	ycq "github.com/jetbasrawi/go.cqrs"
	"go.uber.org/zap"

	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/pkg/commands"
)

type CreateService struct {
	logger     *zap.SugaredLogger
	dispatcher ycq.Dispatcher
}

func (s CreateService) Create(userId string) error {
	id := ycq.NewUUID()
	em := ycq.NewCommandMessage(id, &commands.CreateAccount{
		UserId: userId,
	})
	return s.dispatcher.Dispatch(em)
}

type GetBalanceService struct {
	logger *zap.SugaredLogger
}

func (s CreateService) GetBalance(userId string) error {
	return nil
}
