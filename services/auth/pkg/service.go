package pkg

import (
	"github.com/JohnnyS318/RoyalAfgInGo/services/auth/pkg/services/authentication"
	"github.com/JohnnyS318/RoyalAfgInGo/services/auth/pkg/services/user"
	"net/http"
	"time"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/log"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/mw"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/protos"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/utils"
	"github.com/JohnnyS318/RoyalAfgInGo/services/auth/config"
	"github.com/JohnnyS318/RoyalAfgInGo/services/auth/pkg/handlers"
	"google.golang.org/grpc"

	"github.com/Kamva/mgm/v3"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

// Start starts the account service
func Start() {
	logger := log.NewLogger()
	logger.Warn("Application started. Router will be configured next")
	defer logger.Warn("Application shut down")
	defer logger.Desugar().Sync()

	config.ConfigureDefaults()

	r := mux.NewRouter()

	// Grpc Setup

	logger.Infof("User service url %v trying to connect", viper.GetString(config.UserServiceUrl))
	connectAddrs := viper.GetString(config.UserServiceUrl)
	conn, err := grpc.Dial(connectAddrs, grpc.WithInsecure())
	if err != nil {
		logger.Fatalw("Connection could not be established", "error", err, "target", connectAddrs)
	}
	state := conn.GetState()
	logger.Infow("Calling state", "state", state.String())

	defer conn.Close()

	userServiceClient := protos.NewUserServiceClient(conn)

	//userDb := database.NewUserDatabase(logger)

	// Register Middleware
	loggerHandler := mw.NewLoggerHandler(logger)
	stdChain := alice.New(loggerHandler.LogRoute)

	//services
	userRepo := user.NewUserService(userServiceClient)
	authService := authentication.NewService(userRepo)

	// Handlers
	userHandler := handlers.NewUserHandler(logger, authService)
	authMWHandler := mw.NewAuthMWHandler(logger, viper.GetString(config.JwtSigningKey))

	// Get Subrouters
	postRouter := r.Methods(http.MethodPost).Subrouter()
	getRouter := r.Methods(http.MethodGet).Subrouter()

	prChain := stdChain.Append(loggerHandler.ContentTypeJSON)

	postRouter.Handle("/account/register", prChain.ThenFunc(userHandler.Register))
	postRouter.Handle("/account/login", prChain.ThenFunc(userHandler.Login))
	postRouter.Handle("/account/logout", prChain.Append(authMWHandler.AuthMW).ThenFunc(userHandler.Logout))

	getRouter.Handle("/account/verify", stdChain.Append(authMWHandler.AuthMW).ThenFunc(userHandler.VerifyLoggedIn))

	logger.Debug("Setup Routes")

	// SERVER SETUP
	port := viper.GetString(config.Port)

	srv := &http.Server{
		Addr:         ":" + port,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	utils.StartGracefully(logger, srv, viper.GetDuration(config.GracefulTimeout))
}

func disconnectClient(logger *zap.SugaredLogger, client *mongo.Client) {
	err := client.Disconnect(mgm.Ctx())
	if err != nil {
		logger.Error("MongoDB disconnect", "error", err)
	}
	logger.Warn("Mongodb disconnected")
}
