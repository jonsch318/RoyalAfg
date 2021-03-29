package pkg

import (
	"fmt"
	"net/http"

	"github.com/slok/go-http-metrics/metrics/prometheus"
	metricsMW "github.com/slok/go-http-metrics/middleware"
	"go.uber.org/zap"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/config"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/log"
	"github.com/JohnnyS318/RoyalAfgInGo/services/auth/pkg/handlers"
	"github.com/JohnnyS318/RoyalAfgInGo/services/auth/pkg/rabbit"
	"github.com/JohnnyS318/RoyalAfgInGo/services/auth/pkg/serviceconfig"

	"github.com/JohnnyS318/RoyalAfgInGo/services/auth/pkg/services/authentication"
	"github.com/JohnnyS318/RoyalAfgInGo/services/auth/pkg/services/user"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/urfave/negroni"

	metricsNegroni "github.com/slok/go-http-metrics/middleware/negroni"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/mw"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/utils"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"

	"github.com/rs/cors"
)

// Start starts the account service
func Start(logger *zap.SugaredLogger) {
	//Bind to environment variables
	viper.SetEnvPrefix("auth")
	_ = viper.BindEnv(config.RabbitMQUsername)
	_ = viper.BindEnv(config.RabbitMQPassword)

	//Connect to rabbitmq
	rabbitUrl := fmt.Sprintf("amqp://%s:%s@%s", viper.GetString(config.RabbitMQUsername), viper.GetString(config.RabbitMQPassword), viper.GetString(config.RabbitMQUrl))
	rabbitConn, err := rabbit.NewRabbitMessageBroker(logger, rabbitUrl)
	if err != nil {
		logger.Fatalw("Error during rabbit connection", "error", err)
	}
	defer rabbitConn.Close()

	//Connect to the user service
	userRepo, err := user.NewUser()
	if err != nil {
		log.Logger.Fatalw("Connection could not be established", "error", err, "target", viper.GetString(serviceconfig.UserServiceUrl))
	}
	defer userRepo.Close()

	authService := authentication.NewAuthentication(userRepo)

	// Create handlers
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

	//Register Middleware
	metricsMiddleware := metricsMW.New(metricsMW.Config{
		Recorder: prometheus.NewRecorder(prometheus.Config{}),
		Service:  "authHTTP",
	})
	n := negroni.New(mw.NewLogger(logger.Desugar()), negroni.NewRecovery(), metricsNegroni.Handler("", metricsMiddleware))

	//enable cors if wanted
	if viper.GetBool(config.CorsEnabled) {
		c := cors.New(cors.Options{
			AllowedOrigins:   []string{"http://localhost:3000"},
			AllowCredentials: true,
			Debug:            true,
		})
		n.Use(c)
	}
	n.UseHandler(r)

	logger.Debug("Routes setup successfully")

	//HTTP server setup
	port := viper.GetString(config.HTTPPort)
	logger.Warnf("HTTP Port set to %v", port)
	srv := &http.Server{
		Addr:         ":" + port,
		WriteTimeout: viper.GetDuration(config.WriteTimeout),
		ReadHeaderTimeout:  viper.GetDuration(config.ReadTimeout),
		IdleTimeout:  viper.GetDuration(config.IdleTimeout),
		Handler:      n,
	}

	//Start HTTP server
	utils.StartGracefully(logger, srv, viper.GetDuration(config.GracefulShutdownTimeout))
}
