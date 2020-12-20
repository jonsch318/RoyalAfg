package user

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/log"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/protos"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/utils"

	"github.com/JohnnyS318/RoyalAfgInGo/services/user/pkg/user/config"
	"github.com/JohnnyS318/RoyalAfgInGo/services/user/pkg/user/database"
	"github.com/JohnnyS318/RoyalAfgInGo/services/user/pkg/user/metrics"
	"github.com/JohnnyS318/RoyalAfgInGo/services/user/pkg/user/servers"

	"github.com/Kamva/mgm"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Start starts the user service
func Start() {
	logger := log.RegisterService()
	defer log.CleanLogger()

	metrics := metrics.New()

	// Mongodb configuration
	cfg := &mgm.Config{CtxTimeout: viper.GetDuration(config.DatabaseTimeout)}
	err := mgm.SetDefaultConfig(cfg, viper.GetString(config.DatabaseName), options.Client().ApplyURI(viper.GetString(config.DatabaseUrl)))
	if err != nil {
		logger.Errorf("Connection error to url %v see!", viper.GetString(config.DatabaseUrl))
		logger.Fatalw("Connection to mongo failed", "error", err)
	}
	logger.Debugf("Database connection established to [%v] with database name [%v]", viper.GetString(config.DatabaseUrl), viper.GetString(config.DatabaseName))

	logger.Debugf("with database name [%v]", viper.GetString(config.DatabaseName))

	_, client, _, err := mgm.DefaultConfigs()
	if err != nil {
		logger.Fatalw("Could not get the mongo client", "error", err)
	}
	defer utils.DisconnectClient(logger, client)

	userDatabase := database.NewUserDatabase(logger)

	// grpc server config
	gs := grpc.NewServer()

	userServer := servers.NewUserServer(logger, userDatabase, metrics)

	protos.RegisterUserServiceServer(gs, userServer)

	reflection.Register(gs)

	// create a TCP socket for inbound server connections
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", viper.Get(config.Port)))
	if err != nil {
		logger.Errorw("Unable to create listener", "error", err)
		os.Exit(1)
	}

	// Start the grpc server
	utils.StartGrpcGracefully(logger, gs, l)

	// Prometheus metrics for data sourrounding cpu usage, request/min, etc...
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	srv := &http.Server{
		Addr:         ":" + viper.GetString(config.HttpPort),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      mux,
	}
	utils.StartGracefully(logger, srv, viper.GetDuration(config.GracefulTimeout))
}
