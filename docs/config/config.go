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

func ConfigureDefaults() {
	// HttpServer configuration
	viper.SetDefault(Port, 9000)
	viper.SetDefault(GracefulTimeout, time.Second*20)
	viper.SetDefault(WriteTimeout, time.Second*20)
	viper.SetDefault(ReadTimeout, time.Second*20)
	viper.SetDefault(IdleTimeout, time.Second*60)

}
