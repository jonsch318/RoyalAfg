package aggregates

import (
	"reflect"

	ycq "github.com/jetbasrawi/go.cqrs"
	"go.uber.org/zap"

	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/pkg/commands"
)

type AccountRepo interface {
	Load(string, string) (*Account, error)
	Save(ycq.AggregateRoot, *int) error
}

type AccountCommandHandlers struct {
	logger *zap.SugaredLogger
	repo   AccountRepo
}

func NewAccountCommandHandlers(repo AccountRepo) *AccountCommandHandlers {
	return &AccountCommandHandlers{
		repo: repo,
	}
}

func (h *AccountCommandHandlers) Handle(message ycq.CommandMessage) error {
	var aggregate *Account
	switch cmd := message.Command().(type) {
	case *commands.CreateAccount:
		aggregate = NewAccount(message.AggregateID())
		if err := aggregate.Create(cmd.UserId); err != nil {
			return &ycq.ErrCommandExecution{Command: message, Reason: err.Error()}
		}
		return h.repo.Save(aggregate, ycq.Int(aggregate.CurrentVersion()))
	case *commands.Purchase:
		aggregate, _ = h.repo.Load(reflect.TypeOf(&Account{}).Elem().Name(), message.AggregateID())
		if err := aggregate.Purchase(cmd.Amount, cmd.GameId, cmd.Details); err != nil {
			return &ycq.ErrCommandExecution{Command: message, Reason: err.Error()}
		}
		return h.repo.Save(aggregate, ycq.Int(aggregate.CurrentVersion()))
	case *commands.Deposit:
		aggregate, _ = h.repo.Load(reflect.TypeOf(&Account{}).Elem().Name(), message.AggregateID())
		if err := aggregate.Deposit(cmd.Amount); err != nil {
			return &ycq.ErrCommandExecution{Command: message, Reason: err.Error()}
		}
		return h.repo.Save(aggregate, ycq.Int(aggregate.CurrentVersion()))
	default:
		h.logger.Fatalf("AccountCommandHandlers has received a command that it does not know how to handle: %v", cmd)
	}

	return nil
}
