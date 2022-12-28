package main

import (
	"fmt"
	"net/http"

	"github.com/EventStore/EventStore-Client-Go/esdb"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/bank/system"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/config"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/logging"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/mw"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/utils"
	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/aggregates"
	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/events"
	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/pkg/serviceconfig"
	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/repositories"
	"github.com/gorilla/mux"
	"github.com/slok/go-http-metrics/metrics/prometheus"
	"github.com/spf13/viper"
	"github.com/urfave/negroni"
)

func StartService() {
	//Bind to environment variables
	viper.SetEnvPrefix("bank")
	_ = viper.BindEnv(config.RabbitMQUsername)
	_ = viper.BindEnv(config.RabbitMQPassword)

	//init eventsourcing framework

	// Write models
	//- Add eventbus
	//- Add commandbus
	//- Add commandbus handlers

	// Read models
	//- Add eventbus handlers
	//- Add queryprocessor
	//- Add 

	//################ EventStore ################

	eventStore, err := configEventStore()
	if err != nil {
		logging.Logger.Fatalw("Could not connect to eventstore", "error", err)
	}
	defer eventStore.Close()

	//################ RabbitMQ ################

	rabbitConn, err := configRabbitMQ()

	defer rabbitConn.Close()

	//################ EventSourcing ################

	//Repositories
	eventParser := events.AccountEventParser{}
	factory := &aggregates.AccountFactory{}
	repo := repositories.NewEventStoreRepository[*aggregates.Account](eventStore, &eventParser, factory)

	//Eventbus
	eventBus := system.NewInternalEventBus()

	//################ GRPC ################

	//################ HTTP ################

	//Create HTTP Handlers
	//accountHandler := handlers.NewAccountHandler(dispatcher, eventBus, accountBalanceQuery, accountHistoryQuery)

	//Setup Routes
	r := mux.NewRouter()
	r.Handle("/api/bank/balance", mw.RequireAuth(accountHandler.QueryBalance)).Methods(http.MethodGet)
	r.Handle("/api/bank/history", mw.RequireAuth(accountHandler.QueryHistory)).Methods(http.MethodGet)
	r.Handle("/api/bank/deposit", mw.RequireAuth(accountHandler.Deposit)).Methods(http.MethodPost)
	r.Handle("/api/bank/withdraw", mw.RequireAuth(accountHandler.Withdraw)).Methods(http.MethodPost)

	//We could do this via a grpc connection but a http service already exists. It shows that normal
	r.HandleFunc("/api/bank/verifyAmount", accountHandler.VerifyAmount).Methods(http.MethodGet).Queries("userId", "", "amount", "{i:[0-9]+}")

	//Register Middleware
	metricsMiddleware := metricsMW.New(metricsMW.Config{
		Recorder: prometheus.NewRecorder(prometheus.Config{}),
		Service:  "bankHTTP",
	})
	n := negroni.New(mw.NewLogger(logger.Desugar()), negroni.NewRecovery(), metricsNegroni.Handler("", metricsMiddleware))
	n.UseHandler(r)

	//Create HTTP server
	server := &http.Server{
		Addr:              fmt.Sprintf(":%v", viper.GetString(config.HTTPPort)),
		Handler:           n,
		ReadHeaderTimeout: viper.GetDuration(config.ReadTimeout),
		WriteTimeout:      viper.GetDuration(config.WriteTimeout),
		IdleTimeout:       viper.GetDuration(config.IdleTimeout),
	}

	//Start
	utils.StartGracefully(logger, server, viper.GetDuration(config.GracefulShutdownTimeout))
}

func configEventStore() (*esdb.Client, error) {
	settings, err := esdb.ParseConnectionString(viper.GetString(serviceconfig.EventstoreDbUrl))

	if err != nil {
		return nil, err
	}

	settings.GossipTimeout = 5

	return esdb.NewClient(settings)
}

func configRabbitMQ() (*rabbit.RabbitMQBankClient, error) {
	rabbitURL := fmt.Sprintf("amqp://%s:%s@%s", viper.GetString(config.RabbitMQUsername), viper.GetString(config.RabbitMQPassword), viper.GetString(config.RabbitMQUrl))
	return rabbit.RegisterRabbitMqConsumers(logger, eventBus, dispatcher, rabbitURL)

}

func configureCQRS() (system.InternalCommandBus, system.InternalEventBus, error) {
	//CQRS
	eventBus := system.NewInternalEventBus()
	commandBus := system.NewInternalCommandBus()

	//Register Commands
	commandBus.Subscribe(&commands.CreateAccount{}, accountHandler.CreateAccount)
	dispatcher.RegisterHandler(&commands.CreateAccount{}, accountHandler.CreateAccount)
	dispatcher.RegisterHandler(&commands.Deposit{}, accountHandler.Deposit)
	dispatcher.RegisterHandler(&commands.Withdraw{}, accountHandler.Withdraw)

	return dispatcher, eventBus, nil
}
