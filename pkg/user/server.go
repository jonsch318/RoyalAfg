package user

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/JohnnyS318/RoyalAfgInGo/internal/log"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/user/database"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/user/handlers"
	"github.com/Kamva/mgm/v3"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Start() {
	logger := log.NewLogger()

	logger.Warn("Application started. Router will be configured next")

	r := mux.NewRouter()

	err := mgm.SetDefaultConfig(nil, "RoyalAfgInGo", options.Client().ApplyURI("mongodb://127.0.0.1:27017/"))

	if err != nil {
		panic(err)
	}

	userDb := database.NewUserDatabase(logger)

	postRouter := r.Methods(http.MethodPost).Subrouter()

	userHandler := handlers.NewUserHandler(logger, userDb)
	postRouter.HandleFunc("/account/register", userHandler.Register)
	postRouter.HandleFunc("/account/login", userHandler.Login)

	// SERVER SETUP

	port := 8080

	srv := &http.Server{
		Addr: ":8080",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		logger.Warn("Listening on Port", port)
		if err := srv.ListenAndServe(); err != nil {
			logger.Error(err)
		}
	}()

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)

	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)

	logger.Warn("Application Shutting down!")
	os.Exit(0)

}
