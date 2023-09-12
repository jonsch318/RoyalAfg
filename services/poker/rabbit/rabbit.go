package rabbit

import (
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
	"go.uber.org/zap"

	"github.com/jonsch318/royalafg/pkg/errors"
	"github.com/jonsch318/royalafg/services/poker/serviceconfig"
)

type RabbitMessageBroker struct {
	logger *zap.SugaredLogger
	conn   *amqp.Connection
	ch     *amqp.Channel
}

func NewRabbitMessageBroker(logger *zap.SugaredLogger, url string) (*RabbitMessageBroker, error) {
	conn, err := amqp.Dial(url)

	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()

	if err != nil {
		return nil, err
	}

	if err := ch.ExchangeDeclare(viper.GetString(serviceconfig.RabbitExchange), "direct", true, false, false, false, nil); err != nil {
		logger.Panicw("Exchange Declare Error", "error", err)
	}

	return &RabbitMessageBroker{
		conn:   conn,
		ch:     ch,
		logger: logger,
	}, nil
}

func (r *RabbitMessageBroker) PublishCommand(commandType string, body []byte) error {

	if body == nil {
		return &errors.BodyNullError{}
	}

	headers := make(map[string]interface{})
	headers["CommandType"] = commandType
	if err := r.ch.Publish(
		viper.GetString(serviceconfig.RabbitExchange),
		viper.GetString(serviceconfig.RabbitBankQueue),
		false,
		false,
		amqp.Publishing{
			Headers:      headers,
			ContentType:  "application/json",
			Body:         body,
			DeliveryMode: amqp.Transient,
		}); err != nil {
		r.logger.Errorw("Error during command publishing", "error", err)
		return err
	}

	return nil
}

// Close closes the connection to rabbitmq.
func (r *RabbitMessageBroker) Close() {
	//Close Channel
	r.ch.Close()

	//Close RabbitMQ Connection
	r.conn.Close()
}
