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
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker-matchmaker/pkg/lobby"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker-matchmaker/pkg/models"
)

func main() {

	//logging
	logger := log.RegisterService()
	defer log.CleanLogger()

	//configuration
	//serviceConfig.ReadStandardConfig("search", logger)

	//Gorilla Routing

	loggerHandler := mw.NewLoggerHandler(logger)
	stdChain := alice.New(loggerHandler.LogRoute)

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	kubeConfig, _ := rest.InClusterConfig()
	agonesClient, _ := versioned.NewForConfig(kubeConfig)

	classes := []models.LobbyClass{*models.NewLobbyClass(50, 999, 20),
		*models.NewLobbyClass(1000, 9999, 20),
		*models.NewLobbyClass(10000, 49999, 20),
		*models.NewLobbyClass(50000, 1000000, 20),
	}

	manager := lobby.NewManager(agonesClient, classes)
	ticketHandler := handlers.NewTicket(logger, rdb, agonesClient, manager)

	r := mux.NewRouter()
	r.Handle("/api/poker/ticket", stdChain.ThenFunc(ticketHandler.GetTicketWithParams)).Methods(http.MethodGet).Queries("class", "{class:[0-9]+}")
	r.Handle("/api/poker/ticket/{id}", stdChain.ThenFunc(ticketHandler.GetTicketWithID)).Methods(http.MethodGet)
	// Start Application
	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	utils.StartGracefully(logger, server, viper.GetDuration(config.GracefulShutdownTimeout))

}
