package pkg

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/mw"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/protos"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/utils"
	"github.com/JohnnyS318/RoyalAfgInGo/services/user/pkg/handlers"

	"github.com/gorilla/mux"
	"github.com/slok/go-http-metrics/metrics/prometheus"
	"github.com/urfave/negroni"
	"go.uber.org/zap"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/config"
	"github.com/JohnnyS318/RoyalAfgInGo/services/user/pkg/database"
	"github.com/JohnnyS318/RoyalAfgInGo/services/user/pkg/metrics"
	"github.com/JohnnyS318/RoyalAfgInGo/services/user/pkg/servers"
	"github.com/JohnnyS318/RoyalAfgInGo/services/user/pkg/serviceconfig"

	"github.com/Kamva/mgm"
	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	metricsMW "github.com/slok/go-http-metrics/middleware"
	metricsNegroni "github.com/slok/go-http-metrics/middleware/negroni"
)

// Start starts the user service
func Start(logger *zap.SugaredLogger) {

	metr := metrics.New()

	// Mongodb configuration
	cfg := &mgm.Config{CtxTimeout: viper.GetDuration(serviceconfig.DatabaseTimeout)}
	err := mgm.SetDefaultConfig(cfg, viper.GetString(serviceconfig.DatabaseName), options.Client().ApplyURI(viper.GetString(serviceconfig.DatabaseUrl)))
	if err != nil {
		logger.Errorf("Connection error to url %v see!", viper.GetString(serviceconfig.DatabaseUrl))
		logger.Fatalw("Connection to mongo failed", "error", err)
	}
	logger.Debugf("Database connection established to [%v] with database name [%v]", viper.GetString(serviceconfig.DatabaseUrl), viper.GetString(serviceconfig.DatabaseName))

	logger.Debugf("with database name [%v]", viper.GetString(serviceconfig.DatabaseName))

	_, client, _, err := mgm.DefaultConfigs()
	if err != nil {
		logger.Fatalw("Could not get the mongo client", "error", err)
	}
	defer utils.DisconnectClient(logger, client)

	redis := redis.NewClient(&redis.Options{
		Addr:               viper.GetString(config.RedisAddress),
		Username:           viper.GetString(config.RedisUsername),
		Password:           viper.GetString(config.RedisPassword),
	})

	userCache := cache.New(&cache.Options{
		Redis:        redis,
		LocalCache:   cache.NewTinyLFU(1000, time.Minute),
	})

	userDatabase := database.NewUserDatabase(logger, userCache)

	// grpc server config
	gs := grpc.NewServer()




	userServer := servers.NewUserServer(logger, userDatabase, metr)

	protos.RegisterUserServiceServer(gs, userServer)

	reflection.Register(gs)

	l, err := net.Listen("tcp4", fmt.Sprintf(":%d", viper.Get(serviceconfig.Port)))
	if err != nil {
		logger.Fatalw("Unable to create listener", "error", err)
	}

	// Start the grpc server
	utils.StartGrpcGracefully(logger, gs, l)

	userHandler := handlers.NewUserHandler(logger, userDatabase)

	r := mux.NewRouter()
	r.Handle("/api/user", mw.RequireAuth(userHandler.GetUser)).Methods(http.MethodGet)
	r.Handle("/api/user", mw.RequireAuth(userHandler.UpdateUser)).Methods(http.MethodPut)
	r.Handle("/metrics", promhttp.Handler())

	metricsMiddleware := metricsMW.New(metricsMW.Config{
		Recorder: prometheus.NewRecorder(prometheus.Config{}),
		Service:  "authHTTP",
	})
	n := negroni.New(mw.NewLogger(logger.Desugar()), negroni.NewRecovery(), metricsNegroni.Handler("", metricsMiddleware))
	n.UseHandler(r)

	srv := &http.Server{
		Addr:         ":" + viper.GetString(config.HTTPPort),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      n,
	}
	utils.StartGracefully(logger, srv, viper.GetDuration(config.GracefulShutdownTimeout))
}
