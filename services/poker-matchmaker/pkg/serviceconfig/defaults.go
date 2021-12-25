package serviceconfig

import "github.com/spf13/viper"

const BankServiceUrl = "bank_service_url"
const IncludeBankServiceValidation = "include_bank_service_validation"

const NodeIPAddresses = "node_ips"

func RegisterDefaults() {
	viper.SetDefault(IncludeBankServiceValidation, true)
}
