package commands

import (
	"fmt"
	"log"
	"net/url"
	"reflect"

	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/helpers"
	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/pkg/aggregates"
	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/pkg/repositories"
	ycq "github.com/jetbasrawi/go.cqrs"
	goes "github.com/jetbasrawi/go.geteventstore"
)

// AccountCommandHandlers is the handler for incomming commands
type AccountCommandHandlers struct {
	repo        repositories.AccountRepository
	accountRepo repositories.Account
	client      *goes.Client
}

// NewAccountCommandHandlers creates a new account command handler
func NewAccountCommandHandlers(repo repositories.AccountRepository) *AccountCommandHandlers {
	return &AccountCommandHandlers{
		repo: repo,
	}
}

// Handle handles incoming commands
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
	case *Backroll:
		item, err := h.repo.Load(accountName, message.AggregateID())

		if err != nil {
			return &ycq.ErrAggregateNotFound{
				AggregateID:   message.AggregateID(),
				AggregateType: accountName,
			}
		}

		prevEvents := item.GetChanges()
		prev := &helpers.GeneralTransaction{}
		if prevEvents == nil {
			//no local changes -> load from db
			prev, err = h.LoadPrevCommand(item.AggregateID())

		}

	default:
		log.Printf("account command handler received a command that it cannot handle %v", cmd)
	}

	return nil
}

func (h *AccountCommandHandlers) LoadPrevCommand(aggregatId string) (*helpers.GeneralTransaction, error) {
	aggregateType := reflect.TypeOf(&aggregates.Account{}).Elem().Name()

	streamName, _ := h.accountRepo.StreamNameDelegate.GetStreamName(aggregateType, aggregatId)

	stream := h.client.NewStreamReader(streamName)
	for stream.Next() {
		switch err := stream.Err().(type) {
		case nil:
			break
		case *url.Error, *goes.ErrTemporarilyUnavailable:
			return nil, fmt.Errorf("client not available")
		case *goes.ErrNoMoreEvents:
			return nil, fmt.Errorf("no more events")
		case *goes.ErrUnauthorized:
			return nil, fmt.Errorf("unauthorized")
		case *goes.ErrNotFound:
			return nil, fmt.Errorf("account with the id %v could not be found", aggregatId)
		default:
			return nil, fmt.Errorf("unexpected err %v", err.Error())
		}

		event := make(map[string]interface{})

		//TODO: No test for meta
		meta := make(map[string]string)
		_ = stream.Scan(event, &meta)
		if stream.Err() != nil {
			return nil, stream.Err()
		}

		if event["type"] == "AccountCreated" {
			continue
		}

		event := helpers.ToGeneralTransactionParse(event)

		return

	}
	return nil, nil
}
