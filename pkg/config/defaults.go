package config

import (
	"time"

	"github.com/spf13/viper"
)

const Prod = "prod"

const JWTSigningKey = "jwt_signing_key"
const JWTIssuer = "jwt_issuer"
const JWTExpiresAt = "jwt_expiring_time"

const RabbitMQUrl = "rabbitmq_url"
const RabbitExchange = "rabbit_exchange"
const IdentityCookieKey = "identity_cookie"

func RegisterDefaults() {
	viper.SetDefault(Prod, false)
	viper.SetDefault(JWTIssuer, "github.com/JohnnyS318/RoyalAfgInGo.games")
	viper.SetDefault(JWTExpiresAt, (time.Hour*24*7).Seconds())
	viper.SetDefault(RabbitBankQueue, "ryl_bank")
	viper.SetDefault(RabbitExchange, "ryl")
	viper.SetDefault(IdentityCookieKey, "RYL_SESSION")
}
