package pkg

import (
	"fmt"
	"net"
	"net/http"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"

	"github.com/jonsch318/royalafg/pkg/mw"
	"github.com/jonsch318/royalafg/pkg/protos"
	"github.com/jonsch318/royalafg/pkg/utils"
	"github.com/jonsch318/royalafg/services/user/pkg/handlers"

	"github.com/gorilla/mux"
	"github.com/slok/go-http-metrics/metrics/prometheus"
	"github.com/urfave/negroni"
	"go.uber.org/zap"

	"github.com/jonsch318/royalafg/pkg/config"
	"github.com/jonsch318/royalafg/services/user/pkg/database"
	"github.com/jonsch318/royalafg/services/user/pkg/metrics"
	"github.com/jonsch318/royalafg/services/user/pkg/servers"
	"github.com/jonsch318/royalafg/services/user/pkg/serviceconfig"

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

	//###################### Redis ######################
	//Connect to Redis and setup cache
	red := redis.NewClient(&redis.Options{
		Addr:     viper.GetString(config.RedisAddress),
		Username: viper.GetString(config.RedisUsername),
		Password: viper.GetString(config.RedisPassword),
	})

	defer red.Close()

	userCache := cache.New(&cache.Options{
		Redis:      red,
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})
	logger.Debugf("Redis configured to %v", viper.GetString(config.RedisAddress))

	//###################### MONGO DB ######################

	// Mongodb configuration
	cfg := &mgm.Config{CtxTimeout: viper.GetDuration(serviceconfig.DatabaseTimeout)}
	err := mgm.SetDefaultConfig(cfg, viper.GetString(serviceconfig.DatabaseName), options.Client().ApplyURI(viper.GetString(serviceconfig.DatabaseUrl)))
	if err != nil {
		logger.Fatalw("Could not set the mongodb config", "error", err)
	}

	//Connect to mongo
	_, client, _, err := mgm.DefaultConfigs()
	if err != nil {
		logger.Fatalw("Connection to mongo failed", "error", err)
	}
	defer utils.DisconnectClient(logger, client)
	logger.Debugf("Database connection established to [%v] with database name [%v]", viper.GetString(serviceconfig.DatabaseUrl), viper.GetString(serviceconfig.DatabaseName))

	userDatabase := database.NewUserDatabase(logger, userCache)

	//###################### STATUS REDIS DB ###########

	//TODO: maybe a ring would be better
	statusRed := redis.NewClient(&redis.Options{
		Addr:     viper.GetString(config.StatusRedisAddress),
		Username: viper.GetString(config.StatusRedisUsername),
		Password: viper.GetString(config.StatusRedisPassword),
	})

	defer statusRed.Close()

	statusDB := database.NewOnlineStatusDatabase(logger, statusRed)

	//###################### GRPC ######################
	//Configure GRPC server
	userServer := servers.NewUserServer(logger, userDatabase, statusDB, metrics.New())
	gs := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpc_zap.UnaryServerInterceptor(logger.Desugar()),
				grpc_recovery.UnaryServerInterceptor(),
			),
		),
	)
	protos.RegisterUserServiceServer(gs, userServer)
	reflection.Register(gs)
	l, err := net.Listen("tcp4", fmt.Sprintf(":%d", viper.Get(serviceconfig.Port)))
	if err != nil {
		logger.Fatalw("Unable to create listener", "error", err)
	}

	//Start the GRPC server
	go utils.StartGrpcGracefully(logger, gs, l)

	//###################### HTTP ######################
	//HTTP Handlers
	userHandler := handlers.NewUserHandler(logger, userDatabase, statusDB)

	//Setup Routes
	r := mux.NewRouter()
	r.Handle("/api/user", mw.RequireAuth(userHandler.GetUser)).Methods(http.MethodGet)
	r.Handle("/api/user", mw.RequireAuth(userHandler.UpdateUser)).Methods(http.MethodPut)
	r.Handle("/metrics", promhttp.Handler())

	//RegisterMiddleware
	metricsMiddleware := metricsMW.New(metricsMW.Config{
		Recorder: prometheus.NewRecorder(prometheus.Config{}),
		Service:  "authHTTP",
	})
	n := negroni.New(mw.NewLogger(logger.Desugar()), negroni.NewRecovery(), metricsNegroni.Handler("", metricsMiddleware))
	n.UseHandler(r)

	//Configure HTTP server
	srv := &http.Server{
		Addr:         ":" + viper.GetString(config.HTTPPort),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      n,
	}

	//Start HTTP server
	utils.StartGracefully(logger, srv, viper.GetDuration(config.GracefulShutdownTimeout))
}
