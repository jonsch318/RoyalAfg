package pkg

import (
	"fmt"
	"log"
	"net/http"
	"time"

	jwtMW "github.com/auth0/go-jwt-middleware"
	"github.com/form3tech-oss/jwt-go"
	goes "github.com/jetbasrawi/go.geteventstore"
	"github.com/slok/go-http-metrics/metrics/prometheus"
	metricsMW "github.com/slok/go-http-metrics/middleware"
	metricsNegroni "github.com/slok/go-http-metrics/middleware/negroni"
	"github.com/spf13/viper"
	"github.com/urfave/negroni"

	"github.com/gorilla/mux"
	ycq "github.com/jetbasrawi/go.cqrs"
	"go.uber.org/zap"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/mw"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/utils"
	"github.com/JohnnyS318/RoyalAfgInGo/services/auth/config"
	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/pkg/commands"
	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/pkg/dtos"
	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/pkg/events"
	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/pkg/handlers"
	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/pkg/repositories"
)

func Start(logger *zap.SugaredLogger) {

	//eventstore client config
	eventStore, err := goes.NewClient(nil, "http://localhost:2113")

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


	//Read Model declarations
	accountBalanceQuery := dtos.NewAccountBalanceQuery()
	accountHistoryQuery := dtos.NewAccountHistoryQuery()

	eventBus.AddHandler(accountBalanceQuery, &events.AccountCreated{}, &events.Deposited{}, &events.Withdrawn{})
	eventBus.AddHandler(accountHistoryQuery, &events.AccountCreated{}, &events.Deposited{}, &events.Withdrawn{})

	repo, err := repositories.NewAccount(eventStore, eventBus)

	if err != nil {
		log.Fatalf("account repo err %v", err)
	}

	accountCommandHandler := commands.NewAccountCommandHandlers(repo)
	dispatcher := ycq.NewInMemoryDispatcher()


	err = dispatcher.RegisterHandler(accountCommandHandler, &commands.CreateBankAccount{}, &commands.Deposit{}, &commands.Withdraw{})

	if err != nil {
		log.Fatal(err)
	}

	//TODO: bank message consume.

	//TODO: grpc bank client.


	jwtMiddleware := jwtMW.New(jwtMW.Options{
		Extractor:           jwtMW.FromFirst(mw.ExtractFromCookie, jwtMW.FromAuthHeader),
		ValidationKeyGetter: mw.GetKeyGetter(viper.GetString(config.JwtSigningKey)),
		SigningMethod:       jwt.SigningMethodHS256,
		Debug:               true,
	})

	metricsMiddleware := metricsMW.New(metricsMW.Config{
		Recorder: prometheus.NewRecorder(prometheus.Config{}),
		Service: "bankHTTP",
	})


	accountHandler := handlers.NewAccountHandler(dispatcher, eventBus, accountBalanceQuery, accountHistoryQuery)


	r := mux.NewRouter()
	gr := r.Methods(http.MethodGet).Subrouter()
	pr := r.Methods(http.MethodPost).Subrouter()

	gr.Handle("/api/bank/balance", requireAuth(jwtMiddleware, accountHandler.QueryBalance)).Queries("userId", "")
	gr.Handle("/api/bank/history", requireAuth(jwtMiddleware, accountHandler.QueryHistory)).Queries("userId", "", "i", "{i:[0-9]+}")
	pr.HandleFunc("/api/bank/create", accountHandler.Create)
	pr.Handle("/api/bank/deposit", requireAuth(jwtMiddleware, accountHandler.Deposit))
	pr.Handle("/api/bank/withdraw", requireAuth(jwtMiddleware, accountHandler.Withdraw))


	n := negroni.New(negroni.NewLogger(), negroni.NewRecovery(), metricsNegroni.Handler("", metricsMiddleware))
	n.UseHandler(r)

	port := 8080
	server := &http.Server{
		Addr: fmt.Sprintf(":%v",port),
		Handler: n,
	}

	utils.StartGracefully(logger, server, time.Second * 10)
}

func requireAuth(mw *jwtMW.JWTMiddleware, f func(http.ResponseWriter, *http.Request)) http.Handler {
	nAuth := negroni.New(negroni.HandlerFunc(mw.HandlerWithNext))
	nAuth.UseHandlerFunc(f)
	return nAuth
}