package account

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/JohnnyS318/RoyalAfgInGo/account/pkg/account/database"
	"github.com/JohnnyS318/RoyalAfgInGo/account/pkg/account/handlers"
	"github.com/JohnnyS318/RoyalAfgInGo/shared/pkg/log"
	sharedMiddleware "github.com/JohnnyS318/RoyalAfgInGo/shared/pkg/middleware"
	"github.com/Kamva/mgm/v3"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

// Start starts the account service
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

	defer logger.Desugar().Sync()

	defer disconnectClient(logger, client)

	if err != nil {
		panic(err)
	}

	logger.Warn("MongoDb connected")

	userDb := database.NewUserDatabase(logger)

	// Register Middleware
	loggerHandler := sharedMiddleware.NewLoggerHandler(logger)

	stdChain := alice.New(loggerHandler.LogRouteWithIP)

	// Handlers
	userHandler := handlers.NewUserHandler(logger, userDb)

	// Get Subrouters
	postRouter := r.Methods(http.MethodPost).Subrouter()
	getRouter := r.Methods(http.MethodGet).Subrouter()

	prChain := stdChain.Append(loggerHandler.ContentTypeJSON)

	postRouter.Handle("/account/register", prChain.ThenFunc(userHandler.Register))
	postRouter.Handle("/account/login", prChain.ThenFunc(userHandler.Login))
	postRouter.Handle("/account/logout", prChain.Append(userHandler.AuthMW).ThenFunc(userHandler.Logout))

	getRouter.Handle("/account/verify", stdChain.Append(userHandler.AuthMW).ThenFunc(userHandler.VerifyLoggedIn))

	logger.Debug("Setup Routes")

	opts := middleware.RedocOpts{SpecURL: "http://localhost:8080/swagger.yaml", Title: "RoyalAfg Auth API Documentation"}
	getRouter.Handle("/docs", stdChain.Then(middleware.Redoc(opts, nil)))
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	logger.Debug("Setup swagger docs")

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

func disconnectClient(logger *zap.SugaredLogger, client *mongo.Client) {
	err := client.Disconnect(mgm.Ctx())
	if err != nil {
		logger.Error("MongoDB disconnect", "error", err)
	}
	logger.Warn("Mongodb disconnected")
}
