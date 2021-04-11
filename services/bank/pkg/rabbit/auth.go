package rabbit

import (
	"encoding/json"
	"time"

	"github.com/Rhymond/go-money"
	ycq "github.com/jetbasrawi/go.cqrs"
	"github.com/streadway/amqp"
	"go.uber.org/zap"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/auth"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/currency"
	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/pkg/commands"
)

//AuthCommandHandlers handles rabbitmq messages
type AuthCommandHandler struct {
	logger     *zap.SugaredLogger
	bus        ycq.EventBus
	dispatcher ycq.Dispatcher
}

func NewAuthCommandHandler(logger *zap.SugaredLogger, bus ycq.EventBus, dispatcher ycq.Dispatcher) *AuthCommandHandler {
	return &AuthCommandHandler{
		logger:     logger,
		bus:        bus,
		dispatcher: dispatcher,
	}
}

//Handle handles the rabbitmq message
func (h *AuthCommandHandler) Handle(d *amqp.Delivery) {
	h.logger.Infof("Received Auth message")
	cmd, err := readAuthCommand(d.Body)
	if err != nil {
		h.logger.Errorw("Message deserialization error", "error", err)
		return
	}

	switch cmd.EventType {
	case auth.AccountCreatedEvent:
		h.logger.Infof("Starting account creation")
		//Dispatch Command to the internal command bus
		err = h.dispatcher.Dispatch(ycq.NewCommandMessage(cmd.UserID, &commands.CreateBankAccount{}))
		if err != nil {
			h.logger.Errorw("error during account dispatch", "error", err)
			return
		}

		//Dispatch a default 200€ start Credit
		h.logger.Infof("Account created with id %v", cmd.UserID)
		_ = h.dispatcher.Dispatch(ycq.NewCommandMessage(cmd.UserID, &commands.Deposit{
			Amount:  money.New(20000, currency.Code),
			Time:    time.Now(),
		}))
		h.logger.Infof("Account starting credit %v", cmd.UserID)
	case auth.AccountDeletedEvent:
		//_ = h.dispatcher.Dispatch(ycq.NewCommandMessage(cmd.UserID, &commands.DeleteBankAccount{}))
		//TODO: Delete Bank Account
	}

}

func readAuthCommand(raw []byte) (*auth.AccountCommand, error) {
	cmd := auth.AccountCommand{}
	err := json.Unmarshal(raw, &cmd)
	return &cmd, err
}
