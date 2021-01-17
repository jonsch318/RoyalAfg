package rabbit

import (
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
	"go.uber.org/zap"

	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/serviceConfig"
)

type RabbitMessageBroker struct {
	logger *zap.SugaredLogger
	conn *amqp.Connection
	ch *amqp.Channel
}

func NewRabbitMessageBroker(logger *zap.SugaredLogger) (*RabbitMessageBroker, error) {
	conn, err := amqp.Dial(viper.GetString(serviceConfig.RabbitUrl))

	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()

	if err != nil {
		return nil, err
	}

	if err := ch.ExchangeDeclare(viper.GetString(serviceConfig.RabbitExchange), "direct", true, false, false, false, nil); err != nil {
		logger.Panicw("Exchange Declare Error", "error", err)
	}

	return &RabbitMessageBroker{
		conn: conn,
		ch: ch,
		logger: logger,
	}, nil
}


type BodyNullError struct {}

func (e *BodyNullError) Error() string {
	return "the message should have a body"
}

func (r *RabbitMessageBroker) PublishCommand(commandType string, body []byte) error {

	if body == nil {
		return &BodyNullError{}
	}

	headers := make(map[string]interface{})
	headers["CommandType"] = commandType
	if err := r.ch.Publish(
		viper.GetString(serviceConfig.RabbitExchange),
		viper.GetString(serviceConfig.RabbitBankQueue),
		false,
		false,
		amqp.Publishing{
			Headers: headers,
			ContentType: "application/json",
			Body: body,
			DeliveryMode: amqp.Transient,
		}); err != nil {
		r.logger.Errorw("Error during command publishing", "error", err)
		return err
	}

	return nil
}

//Close closes the connection to rabbitmq.
func (r *RabbitMessageBroker) Close()  {
	//Close Channel
	r.ch.Close()

	//Close RabbitMQ Connection
	r.conn.Close()
}