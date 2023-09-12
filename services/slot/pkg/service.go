package pkg

import (
	"net/http"

	"github.com/Kamva/mgm"
	"github.com/gorilla/mux"
	"github.com/jonsch318/royalafg/pkg/bank"
	"github.com/jonsch318/royalafg/pkg/config"
	"github.com/jonsch318/royalafg/pkg/models"
	"github.com/jonsch318/royalafg/pkg/mw"
	"github.com/jonsch318/royalafg/pkg/protos"
	"github.com/jonsch318/royalafg/pkg/utils"
	localconfig "github.com/jonsch318/royalafg/services/slot/pkg/config"
	"github.com/jonsch318/royalafg/services/slot/pkg/crypto"
	"github.com/jonsch318/royalafg/services/slot/pkg/database"
	"github.com/jonsch318/royalafg/services/slot/pkg/handlers"
	"github.com/jonsch318/royalafg/services/slot/pkg/logic"
	"github.com/spf13/viper"
	"github.com/urfave/negroni"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func Start(logger *zap.SugaredLogger) {

	viper.SetEnvPrefix("slot")

	buffer, gameDatabase := configGameBuffers(logger)

	// ############### Bank Connection ################

	bankService, err := bank.NewRabbitBankConnection(viper.GetString(localconfig.BankURL))

	if err != nil {
		logger.Fatalw("Could not connect to the bank", "error", err)
		return
	}

	// ############### User Service Connection ################

	conn, err := grpc.Dial(viper.GetString(localconfig.UserServiceURL), grpc.WithInsecure())

	defer conn.Close()

	if err != nil {
		logger.Fatalw("Could not connect to user service", "error", err)
	}

	userServiceClient := protos.NewUserServiceClient(conn)

	// ############### Crypto keys and rng ################

	// Read crypto keys
	privateKey, publicKey, err := crypto.ReadECDSAKeys(viper.GetString(localconfig.PublicKeyPath), viper.GetString(localconfig.PrivateKeyPath))

	if err != nil {
		logger.Fatalw("Error reading crypto keys", "error", err)
	}

	// Create the crypto logic
	rng := crypto.NewVRFNumberGenerator(privateKey, publicKey)

	// ############### Game Provider ################

	gameProvider := logic.NewGameProvider(buffer, gameDatabase, rng)

	slotHandler := handlers.NewSlotServer(logger, gameProvider, userServiceClient, bankService)

	r := mux.NewRouter()

	r.HandleFunc("/api/games/slot/spin", slotHandler.Spin).Methods("POST")

	n := negroni.New(negroni.NewRecovery(), mw.NewLogger(logger.Desugar()))

	n.UseHandler(r)

	port := viper.GetString(config.HTTPPort)
	logger.Warnf("HTTP Port set to %v", port)
	srv := &http.Server{
		Addr:              ":" + port,
		WriteTimeout:      viper.GetDuration(config.WriteTimeout),
		ReadHeaderTimeout: viper.GetDuration(config.ReadTimeout),
		IdleTimeout:       viper.GetDuration(config.IdleTimeout),
		Handler:           n,
	}

	utils.StartGracefully(logger, srv, viper.GetDuration(config.GracefulShutdownTimeout))
}

func configGameBuffers(logger *zap.SugaredLogger) (*database.GameBuffer, *database.GameDatabase) {

	cfg := &mgm.Config{CtxTimeout: viper.GetDuration(localconfig.DatabaseTimeout)}
	err := mgm.SetDefaultConfig(cfg, viper.GetString(localconfig.DatabaseName), options.Client().ApplyURI(viper.GetString(localconfig.DatabaseUrl)))

	if err != nil {
		logger.Fatalw("Could not read the mongodb config", "error", err)
	}

	_, client, _, err := mgm.DefaultConfigs()
	if err != nil {
		logger.Fatalw("Connection to mongodb failed", "error", err)
	}

	defer utils.DisconnectClient(logger, client)
	logger.Debugf("Database connection established to [%v] with database name [%v]", viper.GetString(localconfig.DatabaseUrl), viper.GetString(localconfig.DatabaseName))

	gameDatabase, err := database.NewGameDatabase(logger)

	if err != nil {
		logger.Fatalw("Could not create game database", "error", err)
	}

	// ############### Game Databank buffer ################

	buffer := database.NewGameBuffer(func(games []*models.SlotGame) error {
		err := gameDatabase.SaveGameBuffer(games)
		if err != nil {
			logger.Errorw("Could not save game buffer", "error", err)
		}
		return err
	})

	return buffer, gameDatabase
}
