package main

import (
	coresdk "agones.dev/agones/pkg/sdk"
	sdk "agones.dev/agones/sdks/go"
	"fmt"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/config"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/gameServer"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/handlers"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/lobbies"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func main() {
	go gameServer.DoSignal()


	config.SetDefaults()


	log.Println("Creating SDK instance")
	s, err := sdk.NewSDK()
	if err != nil {
		log.Fatalf("Error during sdk connection, %v", err)
	}

	log.Println("Health Ping to agones server management")
	stop := make(chan struct{})
	go gameServer.DoHealthPing(s, stop)

	passthrough := true

	port := viper.GetString(config.Port)

	if passthrough{
		var gs *coresdk.GameServer
		gs, err = s.GameServer()
		if err != nil {
			log.Fatalf("Could not get gameserver port")
		}
		p := strconv.FormatInt(int64(gs.Status.Ports[0].Port), 10)
		port = p
	}

	// Setup
	r := mux.NewRouter()

	// Setup Lobby map

	classes := viper.Get(config.BuyInOptions).([][]int)
	if classes == nil {
		log.Fatal("No buy in classses given")
		return
	}
	lobbyManager := lobbies.NewManager(10, classes)

	// Setup Handlers
	playerHandler := handlers.NewLobbyHandler(lobbyManager, s)

	// each lobby creates a new game and waits for at least 3 players

	// /join {PlayerId, PlayerUsername} searches through the lobby map or creates one if none are found and the maximum lobby count is not succeeded.
	r.HandleFunc("/join", playerHandler.Join)
	r.HandleFunc("/options", playerHandler.LobbyOptions)

	// join upgrades connection to websocket

	// game starts and communicates only via the websocket connection
	log.Printf("Serve on Port %v", port)


	err = s.Ready()
	if err != nil {
		log.Fatalf("Failed to mark Server as Ready")
	}


	log.Printf(http.ListenAndServe(":"+fmt.Sprint(port), r).Error())

}
