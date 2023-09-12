package rabbit

import (
	"encoding/json"

	ycq "github.com/jetbasrawi/go.cqrs"
	"github.com/streadway/amqp"
	"go.uber.org/zap"

	"github.com/jonsch318/royalafg/pkg/bank"
	"github.com/jonsch318/royalafg/pkg/dtos"
	"github.com/jonsch318/royalafg/services/bank/pkg/commands"
)

type BankCommandHandler struct {
	logger     *zap.SugaredLogger
	bus        ycq.EventBus
	dispatcher ycq.Dispatcher
}

func NewBankCommandHandler(logger *zap.SugaredLogger, bus ycq.EventBus, dispatcher ycq.Dispatcher) *BankCommandHandler {
	return &BankCommandHandler{
		logger:     logger,
		bus:        bus,
		dispatcher: dispatcher,
	}
}

func (h *BankCommandHandler) Handle(d *amqp.Delivery) {
	h.logger.Infof("Received bank")
	cmd, err := readBankCommand(d.Body)
	if err != nil {
		h.logger.Errorw("Message deserialization error", "error", err)
		return
	}

	h.logger.Infof("Deserialized: %v", cmd)
	switch cmd.CommandType {
	case bank.Withdraw:
		//Dispatch to the internal command bus
		_ = h.dispatcher.Dispatch(ycq.NewCommandMessage(cmd.UserId, &commands.Withdraw{
			Amount:  dtos.FromDTO(cmd.Amount),
			GameId:  cmd.Game,
			RoundId: cmd.Lobby,
			Time:    cmd.Time,
		}))
	case bank.Deposit:
		//Dispatch to the internal command bus
		_ = h.dispatcher.Dispatch(ycq.NewCommandMessage(cmd.UserId, &commands.Deposit{
			Amount:  dtos.FromDTO(cmd.Amount),
			GameId:  cmd.Game,
			RoundId: cmd.Lobby,
			Time:    cmd.Time,
		}))
	case bank.Rollback:
		_ = h.dispatcher.Dispatch(ycq.NewCommandMessage(cmd.UserId, &commands.Rollback{
			Reason: cmd.Reason,
		}))
	}
}

func readBankCommand(raw []byte) (*bank.Command, error) {
	cmd := bank.Command{}
	err := json.Unmarshal(raw, &cmd)
	return &cmd, err
}
