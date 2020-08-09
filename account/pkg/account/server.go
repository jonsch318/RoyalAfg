package account

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/JohnnyS318/RoyalAfgInGo/account//pkg/account/handlers"
	"github.com/JohnnyS318/RoyalAfgInGo/account/pkg/account/database"
	"github.com/JohnnyS318/RoyalAfgInGo/shared/internal/log"
	sharedMiddleware "github.com/JohnnyS318/RoyalAfgInGo/shared/pkg/middleware"
	"github.com/Kamva/mgm/v3"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Start() {
	logger := log.NewLogger()

	logger.Warn("Application started. Router will be configured next")
	defer logger.Warn("Application shut down")

	r := mux.NewRouter()

	err := mgm.SetDefaultConfig(&mgm.Config{CtxTimeout: 15 * time.Second}, "RoyalAfgInGo", options.Client().ApplyURI("mongodb://127.0.0.1:27017/"))
	_, client, _, err := mgm.DefaultConfigs()

	if err != nil {
		panic(err)
	}
	defer client.Disconnect(mgm.Ctx())

	if err != nil {
		panic(err)
	}

	userDb := database.NewUserDatabase(logger)

	// Register Middleware
	loggerHandler := sharedMiddleware.NewLoggerHandler(logger)
	r.Use(loggerHandler.LogRouteWithIP)

	// Handlers
	userHandler := handlers.NewUserHandler(logger, userDb)

	// Get Subrouters
	postRouter := r.Methods(http.MethodPost).Subrouter()
	getRouter := r.Methods(http.MethodGet).Subrouter()

	postRouter.Use(loggerHandler.ContentTypeJSON)

	postRouter.HandleFunc("/account/register", userHandler.Register)
	postRouter.HandleFunc("/account/login", userHandler.Login)

	getRouter.HandleFunc("/account/verify", userHandler.VerifyLoggedIn)

	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	getRouter.Handle("/docs", sh)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	// SERVER SETUP
	port := 8080

	srv := &http.Server{
		Addr:         ":8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logger.Error(err)
		}
	}()

	logger.Warn("Listening on Port", port)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	srv.Shutdown(ctx)

	logger.Warn("Application Shutting down!")
}
