package pkg

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	ycq "github.com/jetbasrawi/go.cqrs"
	"github.com/justinas/alice"
	"go.uber.org/zap"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/mw"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/utils"
	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/pkg/commands"
	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/pkg/dtos"
	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/pkg/events"
	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/pkg/handlers"
	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/pkg/repositories"
)

func Start(logger *zap.SugaredLogger) {


	eventBus := ycq.NewInternalEventBus()

	//Read Model declarations
	accountBalanceQuery := dtos.NewAccountBalanceQuery()

	eventBus.AddHandler(accountBalanceQuery, &events.AccountCreated{}, &events.Deposited{}, &events.Withdrawn{})


	repo := repositories.NewInMemoryAccount(eventBus)

	accountCommandHandler := commands.NewAccountCommandHandlers(repo)

	dispatcher := ycq.NewInMemoryDispatcher()


	err := dispatcher.RegisterHandler(accountCommandHandler, &commands.CreateBankAccount{}, &commands.Deposit{}, &commands.Withdraw{})

	if err != nil {
		log.Fatal(err)
	}

	loggerHandler := mw.NewLoggerHandler(logger)
	stdChain := alice.New(loggerHandler.LogRoute)

	accountHander := handlers.NewAccountHandler(dispatcher, eventBus, accountBalanceQuery)

	r := mux.NewRouter()

	gr := r.Methods(http.MethodGet).Subrouter()
	pr := r.Methods(http.MethodPost).Subrouter()

	gr.Handle("/api/bank/balance", stdChain.ThenFunc(accountHander.QueryBalance)).Queries("userId", "")
	pr.Handle("/api/bank/create", stdChain.ThenFunc(accountHander.Create))
	pr.Handle("/api/bank/deposit", stdChain.ThenFunc(accountHander.Deposit))
	pr.Handle("/api/bank/withdraw", stdChain.ThenFunc(accountHander.Withdraw))

	port := 8080
	server := &http.Server{
		Addr: fmt.Sprintf(":%v",port),
		Handler: r,
	}

	utils.StartGracefully(logger, server, time.Second * 10)
}
