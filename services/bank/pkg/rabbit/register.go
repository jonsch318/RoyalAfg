package rabbit

import (
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/config"
	ycq "github.com/jetbasrawi/go.cqrs"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

type Connections struct {
	conn *amqp.Connection
	ch *amqp.Channel
}

func RegisterRabbitMqConsumers(logger *zap.SugaredLogger, bus ycq.EventBus, dispatcher ycq.Dispatcher, url string) (*Connections, error) {
	logger.Infof("Connecting to rabbitmq url: %s", url)

	conn, err := amqp.Dial(url)
	if err != nil {
		logger.Fatalw("Could not connect to RabbitMQ message broker", "error", err)
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		logger.Fatalw("Could not connect to RabbitMQ channel", "error", err)
		return nil, err
	}

	bankCommandHandler := NewBankCommandHandler(logger, bus, dispatcher)
	bankConsumer, err := NewEventConsumer(logger, bus, dispatcher, conn, ch, viper.GetString(config.RabbitBankQueue), bankCommandHandler)

	if err != nil {
		logger.Fatalw("bank queue could not be consumed.", "error", err)
		return nil, err
	}

	go func() {
		if err := bankConsumer.Start(); err != nil {
			logger.Fatalw("Error during bank consuming", "error", err)
		}
	}()

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

	return &Connections{
		conn: conn,
		ch: ch,
	} , nil
}

func (c *Connections) Close()  {
	c.conn.Close()
	c.ch.Close()
}