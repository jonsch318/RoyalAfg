package serviceconfig

import "github.com/spf13/viper"

const ReddisUrl = "reddis_url"
const ReddisCred = "reddis_credentials"

const BankServiceUrl = "bank_service_url"

func RegisterDefaults() {
	viper.SetDefault(ReddisUrl, "localhost:6379")
	viper.SetDefault(ReddisCred, "")
}
