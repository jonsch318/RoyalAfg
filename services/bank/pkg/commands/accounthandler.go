package commands

import (
	"log"
	"reflect"

	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/pkg/aggregates"
	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/pkg/repositories"
	ycq "github.com/jetbasrawi/go.cqrs"
)

//AccountCommandHandlers is the handler for incomming commands
type AccountCommandHandlers struct {
	repo repositories.AccountRepository
}

//NewAccountCommandHandlers creates a new account command handler
func NewAccountCommandHandlers(repo repositories.AccountRepository) *AccountCommandHandlers {
	return &AccountCommandHandlers{
		repo: repo,
	}
}

//Handle handles incoming commands
func (h *AccountCommandHandlers) Handle(message ycq.CommandMessage) error {
	accountName := reflect.TypeOf(&aggregates.Account{}).Elem().Name()
	switch cmd := message.Command().(type) {
	case *CreateBankAccount:
		item := aggregates.NewAccount(message.AggregateID())
		if err := item.Create(); err != nil {
			return &ycq.ErrCommandExecution{Command: message, Reason: err.Error()}
		}
		return h.repo.Save(item, ycq.Int(item.OriginalVersion()))
	case *Deposit:
		item, err := h.repo.Load(accountName, message.AggregateID())
		if err != nil {
			return &ycq.ErrAggregateNotFound{
				AggregateID:   message.AggregateID(),
				AggregateType: accountName,
			}
		}
		if err := item.Deposit(cmd.Amount, cmd.GameId, cmd.RoundId, cmd.Time); err != nil {
			return &ycq.ErrCommandExecution{Command: message, Reason: err.Error()}
		}

		return h.repo.Save(item, ycq.Int(item.OriginalVersion()))
	case *Withdraw:
		item, err := h.repo.Load(accountName, message.AggregateID())
		if err != nil {
			return &ycq.ErrAggregateNotFound{
				AggregateID:   message.AggregateID(),
				AggregateType: accountName,
			}
		}
		if err := item.Withdraw(cmd.Amount, cmd.GameId, cmd.RoundId, cmd.Time); err != nil {
			return &ycq.ErrCommandExecution{Command: message, Reason: err.Error()}
		}

		return h.repo.Save(item, ycq.Int(item.OriginalVersion()))

	default:
		log.Printf("account command handler received a command that it cannot handle %v", cmd)
	}

	return nil
}
