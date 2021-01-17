package rabbit

import (
	"encoding/json"
	"log"

	ycq "github.com/jetbasrawi/go.cqrs"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
	"go.uber.org/zap"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/bank"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/config"
	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/pkg/commands"
)

type EvenConsumer struct {
	logger *zap.SugaredLogger
	conn *amqp.Connection
	ch *amqp.Channel
	q amqp.Queue
	bus ycq.EventBus
	dispatcher ycq.Dispatcher
}

func NewBankConsumer(logger *zap.SugaredLogger, bus ycq.EventBus, dispatcher ycq.Dispatcher) (*EvenConsumer, error) {
	conn, err := amqp.Dial(viper.GetString(config.RabbitMQUrl))
	if err != nil {
		logger.Fatalw("Could not connect to RabbitMQ message broker", "error", err)
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		logger.Fatalw("Could not connect to RabbitMQ channel", "error", err)
		return nil, err
	}

	q, err := ch.QueueDeclare(viper.GetString(config.RabbitBankQueue), true, false, false, false, nil)
	if err != nil {
		logger.Fatalw("Could not declare RabbitMQ Queue", "error", err)
		return nil, err
	}

	return &EvenConsumer{
		logger: logger,
		conn:   conn,
		ch:     ch,
		q:      q,
		bus: bus,
		dispatcher: dispatcher,
	}, nil
}

func (c *EvenConsumer) Start() error {
	msgs, err := c.ch.Consume(c.q.Name, "", true, false, false,false, nil)
	if err != nil {
		c.logger.Fatalw("Could not consume declared RabbitMQ Queue", "error", err)
	}

	log.Printf(" [*] Starting consuming rabbit messages.")
	for d := range msgs {
		log.Printf("Received a message: %s", d.Body)

		cmd, err := readMessageBody(d.Body)
		if err == nil {
			switch cmd.CommandType {
			case bank.Withdraw:
				_ = c.dispatcher.Dispatch(ycq.NewCommandMessage(cmd.UserId, &commands.Withdraw{
					Amount:  cmd.Amount,
					GameId:  cmd.Game,
					RoundId: cmd.Lobby,
					Time: cmd.Time,
				}))
			case bank.Deposit:
				_ = c.dispatcher.Dispatch(ycq.NewCommandMessage(cmd.UserId, &commands.Deposit{
					Amount:  cmd.Amount,
					GameId:  cmd.Game,
					RoundId: cmd.Lobby,
					Time: cmd.Time,
				}))
			}
		}
	}

	return nil
}

func readMessageBody(raw []byte) (*bank.Command, error){
	cmd := bank.Command{}
	err := json.Unmarshal(raw, &cmd)
	return &cmd, err
}


func (c EvenConsumer) Close() {

	c.ch.Close()
	c.conn.Close()
}