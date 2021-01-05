package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	sdk "agones.dev/agones/sdks/go"
	coresdk "agones.dev/agones/pkg/sdk"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"

	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/config"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/gameServer"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/handlers"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/lobby"
)

//main method is the entry point of the game server
func main() {
	// configure configuration defaults (ports etc) with viper.
	config.SetDefaults()

	//Register stop signal
	go gameServer.DoSignal()

	//Creating agones sdk instance to communicate with the game server orchestrator
	log.Println("Creating SDK instance")
	s, err := sdk.NewSDK()
	if err != nil {
		log.Fatalf("Error during sdk connection, %v", err)
	}

	//Health ping to agones.
	log.Println("Health Ping to agones server management")
	stop := make(chan struct{})
	go gameServer.DoHealthPing(s, stop)

	var lobbyInstance *lobby.Lobby
	err = s.WatchGameServer(func (gs *coresdk.GameServer){
		err := SetLobby(lobbyInstance, gs, s)
		if err != nil {
			log.Fatalf("some information was not passed: %s", err)
		}
	})
	if err != nil {
		log.Fatalf("Error during sdk annotation subscription: %s", err)
	}

	//game
	gameHandler := handlers.NewGame(lobbyInstance, s)

	//gorilla router instance
	r := mux.NewRouter()
	gr := r.Methods(http.MethodPost).Subrouter()

	gr.HandleFunc("/join", func(rw http.ResponseWriter, r *http.Request){
		log.Printf("/join called from ip  %v", r.RemoteAddr)
		gameHandler.Join(rw, r)
	})


	port := viper.GetString(config.Port)
	log.Printf(http.ListenAndServe(fmt.Sprintf(":%s", port), r).Error())
}

func SetLobby(lobbyInstance *lobby.Lobby, gs *coresdk.GameServer, sdk *sdk.SDK) error {
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

	lobbyID, ok := labels["lobbyID"]
	if !ok {
		return errors.New("the lobbyId of the game server is required")
	}

	lobbyInstance = lobby.NewLobby(min, max, blind, lobbyID, sdk)
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