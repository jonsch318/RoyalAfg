package serviceconfig

import "github.com/spf13/viper"

const RedisUrl = "redis_url"
const RedisCred = "redis_credentials"

const BankServiceUrl = "bank_service_url"

func RegisterDefaults() {
	viper.SetDefault(RedisUrl, "localhost:6379")
	viper.SetDefault(RedisCred, "")
}
