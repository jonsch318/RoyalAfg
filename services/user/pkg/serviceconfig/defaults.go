package serviceconfig

import (
	"time"

	"github.com/spf13/viper"
)

const DatabaseUrl = "mongodb_url"
const DatabaseTimeout = "mongodb_timeout"
const DatabaseName = "mongodb_db_name"

const Port = "grpc_port"

//ConfigureDefaults configures the default values for service specific configurations
func ConfigureDefaults() {
	// HttpServer configuration
	viper.SetDefault(Port, 5001)

	// Database configuration
	viper.SetDefault(DatabaseName, "RoyalafgUser")
	viper.SetDefault(DatabaseTimeout, time.Second*20)
}
