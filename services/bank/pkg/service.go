package pkg

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	goes "github.com/jetbasrawi/go.geteventstore"
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
	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/pkg/commands"
	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/pkg/events"
	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/pkg/handlers"
	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/pkg/projections"
	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/pkg/rabbit"
	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/pkg/repositories"
	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/pkg/serviceconfig"
)

func Start(logger *zap.SugaredLogger) {

	viper.SetEnvPrefix("bank")
	viper.BindEnv(config.RabbitMQUsername)
	viper.BindEnv(config.RabbitMQPassword)

	//eventstore client config
	eventStore, err := goes.NewClient(nil, viper.GetString(serviceconfig.EventstoreDbUrl))

	if err != nil {
		log.Fatalf("eventstore err %v", err)
	}

	//TODO: mongo

	//This service uses an internal event bus that is only accessible in only this instance of the bank service.
	//this is acceptable because both the events (our single source of truth) are persisted in the eventstore db
	//and the read models in a mongodb database instance that is used by all instances.
	//So only a single instance receives the command and publishes the resulting event internally and handles the event internally.
	//the result is then saved into the shared db.
	eventBus := ycq.NewInternalEventBus()

	repo, err := repositories.NewAccount(eventStore, eventBus)
	if err != nil {
		logger.Fatalw("account repo err", "error", err)
	}

	//Read Model declarations
	accountBalanceQuery := projections.NewAccountBalanceQuery(repo)
	accountHistoryQuery := projections.NewAccountHistoryQuery(repo, eventStore)

	eventBus.AddHandler(accountBalanceQuery, &events.AccountCreated{}, &events.Deposited{}, &events.Withdrawn{})
	eventBus.AddHandler(accountHistoryQuery, &events.AccountCreated{}, &events.Deposited{}, &events.Withdrawn{})

	accountCommandHandler := commands.NewAccountCommandHandlers(repo)
	dispatcher := ycq.NewInMemoryDispatcher()

	err = dispatcher.RegisterHandler(accountCommandHandler, &commands.CreateBankAccount{}, &commands.Deposit{}, &commands.Withdraw{})

	if err != nil {
		logger.Fatalw("Could not register handlers", "error", err)
	}

	rabbitURL := fmt.Sprintf("amqp://%s:%s@%s", viper.GetString(config.RabbitMQUsername), viper.GetString(config.RabbitMQPassword), viper.GetString(config.RabbitMQUrl))
	rabbitConnections, err := rabbit.RegisterRabbitMqConsumers(logger, eventBus, dispatcher, rabbitURL)

	if err != nil {
		logger.Fatalw("Could not establish rabbitmq connection", "error", err)
	}

	accountHandler := handlers.NewAccountHandler(dispatcher, eventBus, accountBalanceQuery, accountHistoryQuery)
	r := mux.NewRouter()
	r.Handle("/api/bank/balance", mw.RequireAuth(accountHandler.QueryBalance)).Methods(http.MethodGet)
	r.Handle("/api/bank/history", mw.RequireAuth(accountHandler.QueryHistory)).Methods(http.MethodGet)
	r.Handle("/api/bank/deposit", mw.RequireAuth(accountHandler.Deposit)).Methods(http.MethodPost)
	r.Handle("/api/bank/withdraw", mw.RequireAuth(accountHandler.Withdraw)).Methods(http.MethodPost)

	r.HandleFunc("/api/bank/verifyAmount", accountHandler.VerifyAmount).Methods(http.MethodGet).Queries("userId", "", "amount", "{i:[0-9]+}")

	metricsMiddleware := metricsMW.New(metricsMW.Config{
		Recorder: prometheus.NewRecorder(prometheus.Config{}),
		Service:  "bankHTTP",
	})
	n := negroni.New(mw.NewLogger(logger.Desugar()), negroni.NewRecovery(), metricsNegroni.Handler("", metricsMiddleware))
	n.UseHandler(r)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%v", viper.GetString(config.HTTPPort)),
		Handler: n,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			logger.Fatalw("Http server could not listen and serve", "error", err)
		}
	}()
	logger.Warnf("Http server Listening on address %v", server.Addr)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	rabbitConnections.Close()

	ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration(config.GracefulShutdownTimeout))
	defer cancel()
	err = server.Shutdown(ctx)

	if err != nil {
		logger.Fatalw("Http server encountered an error while shutting down", "error", err)
	}

	logger.Warn("Http server shutting down")
}
