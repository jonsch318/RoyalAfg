package rabbit

import (
	"encoding/json"
	"time"

	"github.com/Rhymond/go-money"
	ycq "github.com/jetbasrawi/go.cqrs"
	"github.com/streadway/amqp"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/auth"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/currency"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/logging"
	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/pkg/commands"
)

// AuthCommandHandlers handles rabbitmq messages
type AuthCommandHandler struct {
	bus        ycq.EventBus
	dispatcher ycq.Dispatcher
}

func NewAuthCommandHandler(bus ycq.EventBus, dispatcher ycq.Dispatcher) *AuthCommandHandler {
	return &AuthCommandHandler{
		bus:        bus,
		dispatcher: dispatcher,
	}
}

// Handle handles the rabbitmq message
func (h *AuthCommandHandler) Handle(d *amqp.Delivery) {
	logging.Logger.Debugf("Received a message: %s", d.Body)

	cmd, err := readAuthCommand(d.Body)
	if err != nil {
		logging.Logger.Errorw("Meesage deserialization error", "error", err)
		return
	}

	switch cmd.EventType {
	case auth.AccountCreatedEvent:
		logging.Logger.Infof("Received AccountCreatedEvent")
		//Dispatch Command to the internal command bus
		err = h.dispatcher.Dispatch(ycq.NewCommandMessage(cmd.UserID, &commands.CreateBankAccount{}))
		if err != nil {
			logging.Logger.Errorw("Could not dispatch command", "error", err)
			return
		}

		//Dispatch a default 200€ start Credit
		h.logger.Infof("Account created with id %v", cmd.UserID)
		_ = h.dispatcher.Dispatch(ycq.NewCommandMessage(cmd.UserID, &commands.Deposit{
			Amount: money.New(20000, currency.Code),
			Time:   time.Now(),
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
