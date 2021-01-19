package config

import (
	"github.com/spf13/viper"
)

const UserServiceUrl = "service_user_url"

const Pepper = "pepper"

// ConfigureDefaults configures the viper default configuration values
func ConfigureDefaults() {
	viper.SetDefault(UserServiceUrl, "royalafg-user.royalafg.svc.cluster.local:8080")
	viper.SetDefault(Pepper, "")
}
