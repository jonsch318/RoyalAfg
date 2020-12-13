package config

import (
	"time"

	"github.com/spf13/viper"
)

const Port = "port"
const HttpPort = "http_port"
const GracefulTimeout = "gracefultimeout"
const WriteTimeout = "HttpServer.WriteTimeout"
const ReadTimeout = "HttpServer.ReadTimeout"
const IdleTimeout = "HttpServer.IdleTimeout"

const DatabaseName = "dbname"
const DatabaseTimeout = "Database.Timeout"
const DatabaseUrl = "mongodburl"

func ConfigureDefaults() {
	// HttpServer configuration
	viper.SetDefault(Port, 5001)
	viper.SetDefault(GracefulTimeout, time.Second*20)
	viper.SetDefault(HttpPort, 8080)
	viper.SetDefault(WriteTimeout, time.Second*20)
	viper.SetDefault(ReadTimeout, time.Second*20)
	viper.SetDefault(IdleTimeout, time.Second*60)

	// Database configuration
	viper.SetDefault(DatabaseName, "RoyalafgUserService")
	viper.SetDefault(DatabaseTimeout, time.Second*20)
}
