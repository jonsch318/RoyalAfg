package main

import (
	"context"
	"fmt"
	"net/http"

	"agones.dev/agones/pkg/client/clientset/versioned"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/cors"
	"github.com/slok/go-http-metrics/metrics/prometheus"
	metricsMW "github.com/slok/go-http-metrics/middleware"
	metricsNegroni "github.com/slok/go-http-metrics/middleware/negroni"
	"github.com/spf13/viper"
	"github.com/urfave/negroni"
	"k8s.io/client-go/rest"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/config"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/log"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/mw"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/poker/models"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/utils"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker-matchmaker/pkg/handlers"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker-matchmaker/pkg/lobby"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker-matchmaker/serviceconfig"
)

func main() {

	//logging
	logger := log.RegisterService()
	defer log.CleanLogger()

	//configuration
	config.ReadStandardConfig("poker-matchmaker", logger)
	serviceconfig.RegisterDefaults()

	rdb := redis.NewClient(&redis.Options{
		Addr:     viper.GetString(serviceconfig.RedisUrl),
		Password: viper.GetString(serviceconfig.RedisCred),
		DB:       0,
	})

	//Check redis config
	if err := ping(rdb); err != nil {
		logger.Fatalw("Redis ping failed", "error", err)
	}
	logger.Info("Redis ping success")

	kubeConfig, _ := rest.InClusterConfig()
	agonesClient, _ := versioned.NewForConfig(kubeConfig)

	classes := []models.Class{*models.NewClass(50, 999, 20),
		*models.NewClass(1000, 9999, 20),
		*models.NewClass(10000, 49999, 20),
		*models.NewClass(50000, 1000000, 20),
	}

	manager := lobby.NewManager(logger, agonesClient, classes, rdb)
	ticketHandler := handlers.NewTicket(logger, rdb, agonesClient, manager)

	r := mux.NewRouter()
	r.Handle("/api/poker/ticket", mw.RequireAuth(ticketHandler.GetTicketWithParams)).Methods(http.MethodGet).Queries("class", "{class:[0-9]+}", "buyIn", "{buyIn:[0-9]+}")
	r.Handle("/api/poker/ticket/{id}", mw.RequireAuth(ticketHandler.GetTicketWithID)).Methods(http.MethodGet).Queries("class", "{class:[0-9]+}", "buyIn", "{buyIn:[0-9]+}")
	r.HandleFunc("/api/poker/pokerinfo", ticketHandler.PokerInfo).Methods(http.MethodGet)
	r.HandleFunc("/api/poker/classinfo", ticketHandler.ClassInfo).Methods(http.MethodGet)
	r.Handle("/metrics", promhttp.Handler())

	metricsMiddleware := metricsMW.New(metricsMW.Config{
		Recorder: prometheus.NewRecorder(prometheus.Config{}),
		Service:  "PokerMatchMaker",
	})

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
		Debug:            false,
	})

	n := negroni.New(negroni.NewLogger(), negroni.NewRecovery(), metricsNegroni.Handler("", metricsMiddleware), cors)
	n.UseHandler(r)

	// Start Application
	server := &http.Server{
		Addr:    fmt.Sprintf(":%v", viper.GetInt(config.HTTPPort)),
		Handler: n,
	}

	utils.StartGracefully(logger, server, viper.GetDuration(config.GracefulShutdownTimeout))
}

func ping(client *redis.Client) error {
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		return err
	}
	fmt.Println(pong)

	return nil
}
