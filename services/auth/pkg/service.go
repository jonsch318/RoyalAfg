package pkg

import (
	"net/http"
	"time"

	"github.com/slok/go-http-metrics/metrics/prometheus"
	metricsMW"github.com/slok/go-http-metrics/middleware"

	"github.com/JohnnyS318/RoyalAfgInGo/services/auth/pkg/handlers"
	"github.com/JohnnyS318/RoyalAfgInGo/services/auth/pkg/services/authentication"
	"github.com/JohnnyS318/RoyalAfgInGo/services/auth/pkg/services/user"

	jwtMW "github.com/auth0/go-jwt-middleware"
	"github.com/form3tech-oss/jwt-go"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/urfave/negroni"

	"google.golang.org/grpc"

	metricsNegroni "github.com/slok/go-http-metrics/middleware/negroni"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/log"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/mw"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/protos"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/utils"
	"github.com/JohnnyS318/RoyalAfgInGo/services/auth/config"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
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

	logger.Infof("Auth service url %v trying to connect", viper.GetString(config.UserServiceUrl))
	conn, err := grpc.Dial(viper.GetString(config.UserServiceUrl), grpc.WithInsecure())
	if err != nil {
		logger.Fatalw("Connection could not be established", "error", err, "target", viper.GetString(config.UserServiceUrl))
	}
	state := conn.GetState()
	logger.Infow("Calling state", "state", state.String())

	defer conn.Close()

	userServiceClient := protos.NewUserServiceClient(conn)


	//Middleware config
	jwtMiddleware := jwtMW.New(jwtMW.Options{
		Extractor:           jwtMW.FromFirst(mw.ExtractFromCookie, jwtMW.FromAuthHeader),
		ValidationKeyGetter: mw.GetKeyGetter(viper.GetString(config.JwtSigningKey)),
		SigningMethod:       jwt.SigningMethodHS256,
		Debug:               true,
	})

	metricsMiddleware := metricsMW.New(metricsMW.Config{
		Recorder:               prometheus.NewRecorder(prometheus.Config{	}),
		Service:                "authHTTP",
	})

	//services
	userRepo := user.NewUserService(userServiceClient)
	authService := authentication.NewService(userRepo)

	// Handlers
	authHandler := handlers.NewAuth(logger, authService)

	// Get Subrouters
	postRouter := r.Methods(http.MethodPost).Subrouter()
	getRouter := r.Methods(http.MethodGet).Subrouter()

	postRouter.HandleFunc("/api/auth/register", authHandler.Register)
	postRouter.HandleFunc("/api/auth/login", authHandler.Login)

	//Required Authenticated Request
	postRouter.Handle("/api/auth/logout", mw.RequireAuth(authHandler.Logout))
	getRouter.Handle("/api/auth/verify", mw.RequireAuth(authHandler.VerifyLoggedIn))
	getRouter.HandleFunc("/api/auth/session", authHandler.Session)

	//Exposes metrics to prometheus
	getRouter.Handle("/metrics", promhttp.Handler())

	n := negroni.New(negroni.NewLogger(), negroni.NewRecovery(), metricsNegroni.Handler("", metricsMiddleware))
	n.UseHandler(r)

	logger.Debug("Setup Routes")

	// SERVER SETUP
	port := viper.GetString(config.Port)

	srv := &http.Server{
		Addr:         ":" + port,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      n,
	}

	utils.StartGracefully(logger, srv, viper.GetDuration(config.GracefulTimeout))
}
