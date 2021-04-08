package serviceconfig

import "github.com/spf13/viper"

const BankServiceUrl = "bank_service_url"
const IncludeBankServiceValidation = "include_bank_service_validation"

func RegisterDefaults() {
	viper.SetDefault(IncludeBankServiceValidation, true)
}
