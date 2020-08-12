package user

import (
	"net/http"

	"github.com/JohnnyS318/RoyalAfgInGo/shared/pkg/log"
	sharedMiddleware "github.com/JohnnyS318/RoyalAfgInGo/shared/pkg/middleware"
	"github.com/JohnnyS318/RoyalAfgInGo/shared/pkg/utils"
	"github.com/JohnnyS318/RoyalAfgInGo/user/pkg/user/config"
	"github.com/Kamva/mgm"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Start() {
	logger := log.NewLogger()
	logger.Warn("User service now running")

	defer logger.Warn("User service shut down")
	defer logger.Desugar().Sync()

	// Mongodb configuration
	cfg := &mgm.Config{CtxTimeout: viper.GetDuration(config.DatabaseTimeout)}
	err := mgm.SetDefaultConfig(cfg, viper.GetString(config.DatabaseName), options.Client().ApplyURI(viper.GetString(config.DatabaseUrl)))
	if err != nil {
		logger.Fatalw("Connection to mongo failed", "error", err)
	}
	_, client, _, err := mgm.DefaultConfigs()
	if err != nil {
		logger.Fatalw("Could not get the mongo client", "error", err)
	}
	defer utils.DisconnectClient(logger, client)

	r := mux.NewRouter()

	loggerHandler := sharedMiddleware.NewLoggerHandler(logger)
	stdChain := alice.New(loggerHandler.LogRouteWithIP)

	gr := r.Methods(http.MethodGet).Subrouter()

	gr.Handle("/hello", stdChain.ThenFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(200)
	}))

	server := &http.Server{
		Addr:         ":" + viper.GetString(config.Port),
		WriteTimeout: viper.GetDuration(config.WriteTimeout),
		ReadTimeout:  viper.GetDuration(config.ReadTimeout),
		IdleTimeout:  viper.GetDuration(config.IdleTimeout),
		Handler:      r,
	}

	utils.StartGracefully(logger, server, viper.GetDuration(config.GracefulTimeout))
}
