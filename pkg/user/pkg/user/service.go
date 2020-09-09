package user

import (
	"fmt"
	"net"
	"os"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/protos"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/shared/pkg/log"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/shared/pkg/utils"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/user/pkg/user/config"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/user/pkg/user/database"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/user/pkg/user/servers"

	"github.com/Kamva/mgm"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Start starts the user service
func Start() {
	logger := log.NewLogger()
	logger.Warn("User service now running")

	defer logger.Warn("User service shut down")
	defer logger.Desugar().Sync()

	// Mongodb configuration
	cfg := &mgm.Config{CtxTimeout: viper.GetDuration(config.DatabaseTimeout)}
	err := mgm.SetDefaultConfig(cfg, viper.GetString(config.DatabaseName), options.Client().ApplyURI(viper.GetString(config.DatabaseUrl)))
	if err != nil {
		logger.Errorf("Connection error to url %v see!", viper.GetString(config.DatabaseUrl))
		logger.Fatalw("Connection to mongo failed", "error", err)
	}
	_, client, _, err := mgm.DefaultConfigs()
	if err != nil {
		logger.Fatalw("Could not get the mongo client", "error", err)
	}
	defer utils.DisconnectClient(logger, client)

	userDatabase := database.NewUserDatabase(logger)

	// grpc server config
	gs := grpc.NewServer()

	userServer := servers.NewUserServer(logger, userDatabase)

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
}
