package bank

import (
	"fmt"
	"net/http"

	"github.com/EventStore/EventStore-Client-Go/esdb"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/config"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/mw"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/utils"
	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/aggregates"
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

	//################ EventStore ################

	eventStore, err := configEventStore()
	if err != nil {
		logger.Fatalw("Could not connect to eventstore", "error", err)
	}
	defer eventStore.Close()

	//################ RabbitMQ ################

	defer rabbitConnections.Close()

	//################ EventSourcing ################

	//Repositories
	repo, err := repositories.NewEventStoreRepository[aggregates.Account](eventStore)

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
