package pkg

import (
	"fmt"
	"net/http"

	"github.com/EventStore/EventStore-Client-Go/esdb"
	"github.com/slok/go-http-metrics/metrics/prometheus"
	metricsMW "github.com/slok/go-http-metrics/middleware"
	metricsNegroni "github.com/slok/go-http-metrics/middleware/negroni"
	"github.com/spf13/viper"
	"github.com/urfave/negroni"

	"github.com/gorilla/mux"
	ycq "github.com/jetbasrawi/go.cqrs"
	"go.uber.org/zap"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/config"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/mw"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/utils"
	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/pkg/commands"
	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/pkg/events"
	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/pkg/handlers"
	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/pkg/projections"
	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/pkg/rabbit"
	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/pkg/repositories"
	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/pkg/serviceconfig"
)

func Start(logger *zap.SugaredLogger) {

	//Bind to environment variables
	viper.SetEnvPrefix("bank")
	_ = viper.BindEnv(config.RabbitMQUsername)
	_ = viper.BindEnv(config.RabbitMQPassword)

	//################ EventStore ################

	eventStore, err := configEventStore()
	if err != nil {
		logger.Fatalw("Could not connect to eventstore", "error", err)
	}

	//TODO: mongo

	//This service uses an internal event bus that is only accessible in only this instance of the bank service.
	//this is acceptable because both the events (our single source of truth) are persisted in the eventstore db.
	//So only a single instance receives the command and publishes the resulting event internally and handles the event internally.
	//the result is then saved into the shared db.
	eventBus := ycq.NewInternalEventBus()

	//Repositories
	repo, err := repositories.NewAccount(eventStore, eventBus)
	if err != nil {
		logger.Fatalw("account repo err", "error", err)
	}

	//Setup Read Model
	accountBalanceQuery := projections.NewAccountBalanceQuery(repo)
	accountHistoryQuery := projections.NewAccountHistoryQuery(repo, eventStore)

	//Register Read Models with event bus
	eventBus.AddHandler(accountBalanceQuery, &events.AccountCreated{}, &events.Deposited{}, &events.Withdrawn{})
	eventBus.AddHandler(accountHistoryQuery, &events.AccountCreated{}, &events.Deposited{}, &events.Withdrawn{})

	//Command Handlers
	accountCommandHandler := commands.NewAccountCommandHandlers(repo)
	dispatcher := ycq.NewInMemoryDispatcher()
	err = dispatcher.RegisterHandler(accountCommandHandler, &commands.CreateBankAccount{}, &commands.Deposit{}, &commands.Withdraw{})
	if err != nil {
		logger.Fatalw("Could not register handlers", "error", err)
	}

	//Configure RabbitMq Message Broker
	rabbitURL := fmt.Sprintf("amqp://%s:%s@%s", viper.GetString(config.RabbitMQUsername), viper.GetString(config.RabbitMQPassword), viper.GetString(config.RabbitMQUrl))
	rabbitConnections, err := rabbit.RegisterRabbitMqConsumers(logger, eventBus, dispatcher, rabbitURL)
	if err != nil {
		logger.Fatalw("Could not establish rabbitmq connection", "error", err)
	}
	defer rabbitConnections.Close()

	//Create HTTP Handlers
	accountHandler := handlers.NewAccountHandler(dispatcher, eventBus, accountBalanceQuery, accountHistoryQuery)

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
