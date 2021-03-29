package serviceconfig

import (
	"github.com/spf13/viper"
)

const UserServiceUrl = "userservice_url"

const Pepper = "pepper"

// ConfigureDefaults configures the viper default configuration values
func ConfigureDefaults() {
	viper.SetDefault(UserServiceUrl, "localhost:5001")
	viper.SetDefault(Pepper, "")
}
