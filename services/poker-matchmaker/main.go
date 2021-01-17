package main

import (
	"net/http"

	"agones.dev/agones/pkg/client/clientset/versioned"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/slok/go-http-metrics/metrics/prometheus"
	metricsMW "github.com/slok/go-http-metrics/middleware"
	metricsNegroni "github.com/slok/go-http-metrics/middleware/negroni"
	"github.com/spf13/viper"
	"github.com/urfave/negroni"
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
	r.Handle("/api/poker/ticket", mw.RequireAuth(ticketHandler.GetTicketWithParams)).Methods(http.MethodGet).Queries("class", "{class:[0-9]+}", "buyIn", "{buyIn:[0-9]+}")
	r.Handle("/api/poker/ticket/{id}", mw.RequireAuth(ticketHandler.GetTicketWithID)).Methods(http.MethodGet).Queries("buyIn", "{buyIn:[0-9]+}")
	r.Handle("/metrics", promhttp.Handler())

	metricsMiddleware := metricsMW.New(metricsMW.Config{
		Recorder: prometheus.NewRecorder(prometheus.Config{
		}),
		Service:                "PokerMatchMaker",
	})
	n := negroni.New(negroni.NewLogger(), negroni.NewRecovery(), metricsNegroni.Handler("", metricsMiddleware))
	n.UseHandler(r)

	// Start Application
	server := &http.Server{
		Addr:    ":8080",
		Handler: n,
	}

	utils.StartGracefully(logger, server, viper.GetDuration(config.GracefulShutdownTimeout))

}
