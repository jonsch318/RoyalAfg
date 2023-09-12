package rabbit

import (
	"log"

	ycq "github.com/jetbasrawi/go.cqrs"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
	"go.uber.org/zap"

	"github.com/jonsch318/royalafg/pkg/config"
)

type CommandHandler interface {
	Handle(*amqp.Delivery)
}

type EventConsumer struct {
	logger         *zap.SugaredLogger
	conn           *amqp.Connection
	ch             *amqp.Channel
	q              amqp.Queue
	bus            ycq.EventBus
	dispatcher     ycq.Dispatcher
	commandHandler CommandHandler
}

func NewEventConsumer(logger *zap.SugaredLogger, bus ycq.EventBus, dispatcher ycq.Dispatcher, conn *amqp.Connection, ch *amqp.Channel, queue string, handler CommandHandler) (*EventConsumer, error) {

	q, err := ch.QueueDeclare(queue, false, false, false, false, nil)
	if err != nil {
		logger.Fatalw("Could not declare RabbitMQ Queue", "error", err)
		return nil, err
	}
	logger.Infof("Queue declared on %v", q.Name)

	if err := ch.ExchangeDeclare(viper.GetString(config.RabbitExchange), "direct", true, false, false, false, nil); err != nil {
		logger.Fatalw("Exchange Declare Error", "error", err)
	}

	err = ch.QueueBind(q.Name, q.Name, viper.GetString(config.RabbitExchange), false, nil)

	if err != nil {
		logger.Fatalw("Could not bind RabbitMQ Queue", "error", err)
		return nil, err
	}

	return &EventConsumer{
		logger:         logger,
		conn:           conn,
		ch:             ch,
		q:              q,
		bus:            bus,
		dispatcher:     dispatcher,
		commandHandler: handler,
	}, nil
}

// Start the rabbitmq consumer
func (c *EventConsumer) Start() error {
	messages, err := c.ch.Consume(c.q.Name, "", true, false, false, false, nil)
	if err != nil {
		c.logger.Fatalw("Could not consume declared RabbitMQ Queue", "error", err)
	}

	c.logger.Infof("Message received on queue: %v", c.q.Name)
	log.Printf(" [*] Starting consuming rabbit messages.")
	for d := range messages {
		log.Printf("Received a message [%s]: %s", c.q.Name, d.Body)
		c.commandHandler.Handle(&d)
	}

	return nil
}
