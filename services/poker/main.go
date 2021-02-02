package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	coresdk "agones.dev/agones/pkg/sdk"
	sdk "agones.dev/agones/sdks/go"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"github.com/urfave/negroni"
	"golang.org/x/net/context"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/config"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/log"
	pokerModels "github.com/JohnnyS318/RoyalAfgInGo/pkg/poker/models"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/utils"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/bank"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/gameServer"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/handlers"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/lobby"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/rabbit"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/serviceConfig"
)

//main method is the entry point of the game server
func main() {
	logger := log.RegisterService()
	defer log.CleanLogger()

	config.ReadStandardConfig("poker", logger)

	serviceConfig.SetDefaults()

	viper.SetEnvPrefix("poker")
	viper.BindEnv(config.RabbitMQUsername)
	viper.BindEnv(config.RabbitMQPassword)

	//connect to rabbitmq to send user commands
	rabbitURL := fmt.Sprintf("amqp://%s:%s@%s", viper.GetString(config.RabbitMQUsername), viper.GetString(config.RabbitMQPassword), viper.GetString(config.RabbitMQUrl))
	rabbitConn, err := rabbit.NewRabbitMessageBroker(logger, rabbitURL)

	if err != nil {
		logger.Fatalw("Could not connect to service bus", "error", err)
	}

	b := bank.NewBank(rabbitConn)

	//Register stop signal
	go gameServer.DoSignal()

	//Creating agones sdk instance to communicate with the game server orchestrator
	logger.Info("Creating SDK instance")
	s, err := sdk.NewSDK()
	if err != nil {
		logger.Fatalf("Error during sdk connection, %v", err)
	}

	//Health ping to agones.
	logger.Info("Health Ping to agones server management")
	stop := make(chan struct{})
	go gameServer.DoHealthPing(s, stop)

	serviceConfig.SetDefaults()

	shutDownStop := make(chan interface{})

	lobbyInstance := lobby.NewLobby(b, s)
	err = s.WatchGameServer(func(gs *coresdk.GameServer) {
		err := SetLobby(b, lobbyInstance, gs, s)
		if err == nil {
			logger.Warnw("Lobby configured", "id", lobbyInstance.LobbyID)
			if lobbyInstance.TotalPlayerCount() <= 0 {
				go StartShutdownTimer(shutDownStop, s)
			}
		}

	})
	if err != nil {
		logger.Fatalf("Error during sdk annotation subscription: %s", err)
	}

	//game
	gameHandler := handlers.NewGame(lobbyInstance, s, shutDownStop)

	//gorilla router instance
	r := mux.NewRouter()
	r.HandleFunc("/api/poker/join", gameHandler.Join).Methods(http.MethodGet)
	r.HandleFunc("/api/poker/health", gameHandler.Health).Methods(http.MethodGet)

	recoverMW := negroni.NewRecovery()
	recoverMW.PanicHandlerFunc = func(information *negroni.PanicInformation) {
		s.Shutdown()
	}
	n := negroni.New(negroni.NewLogger(), negroni.NewRecovery())
	n.UseHandler(r)

	s.Ready()

	logger.Info("SDK Ready called")

	port := viper.GetString(config.HTTPPort)
	utils.StartGracefully(logger, &http.Server{
		Addr:    fmt.Sprintf(":%v", port),
		Handler: n,
	}, viper.GetDuration(config.GracefulShutdownTimeout))

	rabbitConn.Close()
}

func SetLobby(b *bank.Bank, lobbyInstance *lobby.Lobby, gs *coresdk.GameServer, sdk *sdk.SDK) error {
	labels := gs.GetObjectMeta().GetLabels()
	min, err := GetFromLabels("min-buy-in", labels)
	if err != nil {
		return err
	}

	max, err := GetFromLabels("max-buy-in", labels)
	if err != nil {
		return err
	}

	blind, err := GetFromLabels("blind", labels)
	if err != nil {
		return err
	}

	index, err := GetFromLabels("class-index", labels)
	if err != nil {
		return err
	}

	lobbyId, ok := labels["lobbyId"]
	if !ok {
		return errors.New("can not get the required information for the key lobbyId")
	}
	b.RegisterLobby(lobbyId)



	lobbyInstance.RegisterLobbyValue(&pokerModels.Class{
		Min:   min,
		Max:   max,
		Blind: blind,
	}, index, lobbyId)
	return nil
}

func GetFromLabels(key string, labels map[string]string) (int, error) {
	valString, ok := labels[key]

	if !ok {
		return 0, errors.New("key needed")
	}

	val, err := strconv.Atoi(valString)
	if err != nil {
		return 0, err
	}

	return val, nil
}

func StartShutdownTimer(stop chan interface{}, s *sdk.SDK) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(10*time.Minute))
	select {
	case _ = <-stop:
		cancel()
		//Cancel shutdown
	case <-ctx.Done():
		s.Shutdown()
	}
}
