package config

import (
	"time"

	"github.com/spf13/viper"
)

const Port = "HttpServer.Port"
const GracefulTimeout = "HttpServer.GracefulTimeout"
const WriteTimeout = "HttpServer.WriteTimeout"
const ReadTimeout = "HttpServer.ReadTimeout"
const IdleTimeout = "HttpServer.IdleTimeout"

const DatabaseName = "Database.Name"
const DatabaseTimeout = "Database.Timeout"
const DatabaseUrl = "Database.Url"

func ConfigureDefaults() {
	// HttpServer configuration
	viper.SetDefault(Port, 5001)
	viper.SetDefault(GracefulTimeout, time.Second*20)
	viper.SetDefault(WriteTimeout, time.Second*20)
	viper.SetDefault(ReadTimeout, time.Second*20)
	viper.SetDefault(IdleTimeout, time.Second*60)

	// Database configuration
	viper.SetDefault(DatabaseName, "RoyalAfgUserService")
	viper.SetDefault(DatabaseTimeout, time.Second*20)
}
