package main

import (
"net/http"

	"agones.dev/agones/pkg/client/clientset/versioned"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
"github.com/justinas/alice"
"github.com/spf13/viper"
	"k8s.io/client-go/rest"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/config"
"github.com/JohnnyS318/RoyalAfgInGo/pkg/log"
"github.com/JohnnyS318/RoyalAfgInGo/pkg/mw"
"github.com/JohnnyS318/RoyalAfgInGo/pkg/utils"
"github.com/JohnnyS318/RoyalAfgInGo/services/poker-matchmaker/pkg/handlers"
)

func main() {

	//logging
	logger := log.RegisterService()
	defer log.CleanLogger()

	//configuration
	//config.ReadStandardConfig("search", logger)

	//Gorilla Routing

	loggerHandler := mw.NewLoggerHandler(logger)
	stdChain := alice.New(loggerHandler.LogRoute)


	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	kubeConfig, _ := rest.InClusterConfig()

	agonesClient, _ := versioned.NewForConfig(kubeConfig)

	ticketHandler := handlers.NewTicket(logger, rdb, agonesClient)

	r := mux.NewRouter()
	r.Handle("/api/poker/ticket",  stdChain.ThenFunc(ticketHandler.GetTicket)).Methods(http.MethodPost)

	// Start Application
	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	utils.StartGracefully(logger, server, viper.GetDuration(config.GracefulShutdownTimeout))
}

