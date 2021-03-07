package pkg

import (
	"fmt"
	"net/http"
	"time"

	"github.com/slok/go-http-metrics/metrics/prometheus"
	metricsMW "github.com/slok/go-http-metrics/middleware"
	"go.uber.org/zap"

	gConfig "github.com/JohnnyS318/RoyalAfgInGo/pkg/config"
	"github.com/JohnnyS318/RoyalAfgInGo/services/auth/pkg/handlers"
	"github.com/JohnnyS318/RoyalAfgInGo/services/auth/pkg/services/authentication"
	"github.com/JohnnyS318/RoyalAfgInGo/services/auth/pkg/services/user"
	"github.com/JohnnyS318/RoyalAfgInGo/services/auth/pkg/rabbit"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/urfave/negroni"

	"google.golang.org/grpc"

	metricsNegroni "github.com/slok/go-http-metrics/middleware/negroni"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/mw"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/protos"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/utils"
	"github.com/JohnnyS318/RoyalAfgInGo/services/auth/config"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"

	"github.com/rs/cors"
)

// Start starts the account service
func Start(logger *zap.SugaredLogger) {

	viper.SetEnvPrefix("auth")
	viper.BindEnv(gConfig.RabbitMQUsername)
	viper.BindEnv(gConfig.RabbitMQPassword)

	rabbitUrl := fmt.Sprintf("amqp://%s:%s@%s", viper.GetString(gConfig.RabbitMQUsername),viper.GetString(gConfig.RabbitMQPassword), viper.GetString(gConfig.RabbitMQUrl))
	rabbitConn, err := rabbit.NewRabbitMessageBroker(logger, rabbitUrl)

	if err != nil {
		logger.Fatalw("Error during rabbit connection", "error", err)
	}

	// Grpc Setup set (use grpc.WithInsecure() explicitly or

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


	//services
	userRepo := user.NewUserService(userServiceClient)
	authService := authentication.NewService(userRepo)

	// Handlers
	authHandler := handlers.NewAuth(logger, authService, rabbitConn)

	r := mux.NewRouter()
	// Get Subrouters
	r.HandleFunc("/api/auth/register", authHandler.Register).Methods(http.MethodPost)
	r.HandleFunc("/api/auth/login", authHandler.Login).Methods(http.MethodPost)

	//Required Authenticated Request
	r.Handle("/api/auth/logout", mw.RequireAuth(authHandler.Logout)).Methods(http.MethodPost)
	r.Handle("/api/auth/verify", mw.RequireAuth(authHandler.VerifyLoggedIn)).Methods(http.MethodGet)
	r.HandleFunc("/api/auth/session", authHandler.Session).Methods(http.MethodGet)

	//Exposes metrics to prometheus
	r.Handle("/metrics", promhttp.Handler())

	metricsMiddleware := metricsMW.New(metricsMW.Config{
		Recorder: prometheus.NewRecorder(prometheus.Config{}),
		Service:  "authHTTP",
	})
	n := negroni.New(mw.NewLogger(logger.Desugar()), negroni.NewRecovery(), metricsNegroni.Handler("", metricsMiddleware))

	if viper.GetBool(gConfig.CorsEnabled) {
		cors := cors.New(cors.Options{
			AllowedOrigins:   []string{"http://localhost:3000"},
			AllowCredentials: true,
			Debug:            true,
		})
		n.Use(cors)
	}
	n.UseHandler(r)

	logger.Debug("Setup Routes")

	// SERVER SETUP
	port := viper.GetString(gConfig.HTTPPort)

	logger.Warnf("HTTP Port set to %v", port)

	srv := &http.Server{
		Addr:         ":" + port,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      n,
	}

	utils.StartGracefully(logger, srv, viper.GetDuration(gConfig.GracefulShutdownTimeout))

	rabbitConn.Close()
}
