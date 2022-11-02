package rabbit

import (
	"fmt"

	ycq "github.com/jetbasrawi/go.cqrs"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
	"go.uber.org/zap"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/config"
)

type RabbitMQBankClient struct {
	conn      *amqp.Connection
	consumers []*EventConsumer
	ch        *amqp.Channel
}

func NewRabbitMQBankClient(bus ycq.EventBus, dispatcher ycq.Dispatcher, url string) (*RabbitMQBankClient, error) {
	client := &RabbitMQBankClient{}

	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("could not connect to RabbitMQ message broker: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("could not connect to RabbitMQ channel: %w", err)
	}

	bankConsumer, err := NewEventConsumer(bus, dispatcher, conn, ch, viper.GetString(config.RabbitBankQueue), NewBankCommandHandler(bus, dispatcher))
	if err != nil {
		return nil, fmt.Errorf("bank queue could not be consumed: %w", err)
	}

	client.consumers = append(client.consumers, bankConsumer)

	authCommandHandler := NewAuthCommandHandler(logger, bus, dispatcher)
	authConsumer, err := NewEventConsumer(logger, bus, dispatcher, conn, ch, viper.GetString(config.RabbitAccountQueue), authCommandHandler)

	if err != nil {
		logger.Fatalw("auth queue could not be consumed.", "error", err)
		return nil, err
	}

	go func() {
		if err := authConsumer.Start(); err != nil {
			logger.Fatalw("Error during auth consuming", "error", err)
		}
	}()

}

func (c *Connections) StartListening() {
	for _, consumer := range c.consumers {
	go func() {
		if err = consumer.Start(); err != nil {
			logger.Fatalw("error during consumption", "error", err)
		}
	}()

}

func (c *Connections) Close() {
	_ = c.conn.Close()
	_ = c.ch.Close()
}
