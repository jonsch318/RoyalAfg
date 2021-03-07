package config

import (
	"github.com/spf13/viper"
)

const RabbitMQUsername = "rabbitmq_username"
const RabbitMQPassword = "rabbitmq_password"
const RabbitMQUrl = "rabbitmq_url"

const RabbitExchange = "rabbit_exchange"
const RabbitBankQueue = "bank"
const RabbitAccountQueue = "rabbitmq_auth_queue"

func RegisterRabbitDefaults() {
	viper.SetDefault(RabbitBankQueue, "ryl_bank")
	viper.SetDefault(RabbitExchange, "ryl")
	viper.SetDefault(RabbitAccountQueue, "ryl_auth")
}
