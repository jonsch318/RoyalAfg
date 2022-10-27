package bank

import (
	"encoding/json"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/config"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/errors"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

type RabbitBankConnection struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func NewRabbitBankConnection(url string) (*RabbitBankConnection, error) {
	conn, err := amqp.Dial(url)

	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()

	if err != nil {
		return nil, err
	}

	if err := ch.ExchangeDeclare(viper.GetString(config.RabbitExchange), "direct", true, false, false, false, nil); err != nil {
		return nil, err
	}

	return &RabbitBankConnection{
		conn: conn,
		ch:   ch,
	}, nil
}

func (r *RabbitBankConnection) publishCommand(commandType string, body []byte) error {
	if body == nil {
		return &errors.BodyNullError{}
	}

	headers := make(map[string]interface{})
	headers["CommandType"] = commandType
	if err := r.ch.Publish(
		viper.GetString(config.RabbitExchange),
		viper.GetString(config.RabbitBankQueue),
		false,
		false,
		amqp.Publishing{
			Headers:      headers,
			ContentType:  "application/json",
			Body:         body,
			DeliveryMode: amqp.Transient,
		}); err != nil {
		return err
	}
	return nil
}

func (r *RabbitBankConnection) PublishCommand(command *Command) error {
	buf, err := json.Marshal(command)
	if err != nil {
		return err
	}

	return r.publishCommand(command.CommandType, buf)

}
