package main

import (
	"fmt"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/config"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/handlers"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/lobbies"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func main() {

	config.SetDefaults()

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
	playerHandler := handlers.NewLobbyHandler(lobbyManager)

	// each lobby creates a new game and waits for at least 3 players

	// /join {PlayerId, PlayerUsername} searches through the lobby map or creates one if none are found and the maximum lobby count is not succeeded.
	r.HandleFunc("/join", playerHandler.Join)
	r.HandleFunc("/options", playerHandler.LobbyOptions)

	// join upgrades connection to websocket

	// game starts and communicates only via the websocket connection

	port := viper.GetInt(config.Port)
	log.Printf("Serve on Port %v", port)

	log.Printf(http.ListenAndServe(":"+fmt.Sprint(port), r).Error())

}
