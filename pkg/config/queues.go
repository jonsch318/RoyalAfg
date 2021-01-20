package config

import (
	"github.com/spf13/viper"
)

const RabbitMQUrl = "rabbitmq_url"
const RabbitExchange = "rabbit_exchange"
const RabbitBankQueue = "bank"

func RegisterRabbitDefaults(){
	viper.SetDefault(RabbitBankQueue, "ryl_bank")
	viper.SetDefault(RabbitExchange, "ryl")
}