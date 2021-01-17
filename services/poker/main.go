package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	coresdk "agones.dev/agones/pkg/sdk"
	sdk "agones.dev/agones/sdks/go"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"github.com/urfave/negroni"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/config"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/log"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/bank"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/gameServer"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/handlers"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/lobby"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/serviceConfig"
)

//main method is the entry point of the game server
func main() {
	logger := log.RegisterService()
	defer log.CleanLogger()

	config.ReadStandardConfig("royalafg-poker", logger)

	serviceConfig.SetDefaults()

	//connect to rabbitmq to send user commands

	b := bank.NewBank()







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



	lobbyInstance := lobby.NewLobby(b, s)
	err = s.WatchGameServer(func (gs *coresdk.GameServer){
		err := SetLobby(b, lobbyInstance, gs, s)
		if err != nil {
			logger.Fatalf("some information was not passed: %s", err)
		}
	})
	if err != nil {
		logger.Fatalf("Error during sdk annotation subscription: %s", err)
	}

	//game
	gameHandler := handlers.NewGame(lobbyInstance, s)

	//gorilla router instance
	r := mux.NewRouter()
	r.HandleFunc("/join", gameHandler.Join).Methods(http.MethodPost)

	recoverMW := negroni.NewRecovery()
	recoverMW.PanicHandlerFunc = func(information *negroni.PanicInformation){
		s.Shutdown()
	}
	n := negroni.New(negroni.NewLogger(), negroni.NewRecovery())
	n.UseHandler(r)

	port := viper.GetString(config.Port)
	logger.Warn(http.ListenAndServe(fmt.Sprintf(":%s", port), n).Error())
}

func SetLobby(b *bank.Bank, lobbyInstance *lobby.Lobby, gs *coresdk.GameServer, sdk *sdk.SDK) error {
	labels := gs.GetObjectMeta().GetLabels()
	min, err := GetFromLabels("min-buy-in", labels)
	if err != nil {
		return err
	}

	var max int
	max, err = GetFromLabels("max-buy-in", labels)
	if err != nil {
		return err
	}

	var blind int
	blind, err = GetFromLabels("blind", labels)
	if err != nil {
		return err
	}

	lobbyId, ok := labels["lobbyId"]
	if !ok {
		return errors.New("can not get the required information for the key lobbyId", )
	}
	b.RegisterLobby(lobbyId)

	lobbyInstance.RegisterLobbyValue(min, max, blind)
	return nil
}

func GetFromLabels(key string, labels map[string]string) (int, error) {
	valString, ok := labels[key]

	if !ok {
		return 0, fmt.Errorf("can not get the required information for the key [%s]", key)
	}

	val, err := strconv.Atoi(valString)
	if err != nil {
		return 0, err
	}

	return val, nil
}
