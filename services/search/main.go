package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/log"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/utils"
	"github.com/JohnnyS318/RoyalAfgInGo/services/auth/config"
)

func main() {

	//logging
	logger := log.RegisterService()
	defer log.CleanLogger()

	//Gorilla Routing
	r := mux.NewRouter()
	server := &http.Server{
		Addr: ":" + viper.GetString(config.Port),
		/*		WriteTimeout: viper.GetDuration(config.WriteTimeout),
				ReadTimeout: viper.GetDuration(config.ReadTimeout),
				IdleTimeout: viper.GetDuration(config.IdleTimeout),*/
		Handler: r,
	}

	utils.StartGracefully(logger, server, viper.GetDuration(config.GracefulTimeout))

}
