package handlers

import (
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/events"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/models"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/utils"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *Lobby) Join(rw http.ResponseWriter, r *http.Request) {

	log.Printf("/join called")

	conn, err := upgrader.Upgrade(rw, r, nil)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadGateway)
	}

	playerConn := NewPlayerConn(conn)
	playerConn.OnClose = func(err error) {
		log.Printf("Closing err %v", err)
		return
	}

	go playerConn.reader()
	go playerConn.writer()

	raw, err := utils.WaitUntilEventD(playerConn.In, 2*time.Minute)

	if err != nil {
		http.Error(rw, "Join Timeout. The joining process has timed out.", http.StatusBadRequest)
		return
	}
	joinEvent, err := events.ToJoinEvent(raw)

	if err != nil {
		log.Printf("joinEvent was invalid %v", err)
		utils.SendToChanTimeout(playerConn.Out, models.NewEvent("VALIDATION_FAILED", "The joining event was not as the server expected"))
		conn.Close()
		http.Error(rw, "VALIDATION_FAILED. The joining event was not as the server expected", http.StatusBadRequest)
		return
	}

	player := models.NewPlayer(joinEvent.Username, joinEvent.ID, joinEvent.BuyIn, playerConn.In, playerConn.Out, playerConn.Close)

	_, err = h.Lobbies.DistributePlayer(player, joinEvent)

	if err != nil {
		log.Printf("Error: %v", err)
		playerConn.conn.Close()
		return
	}

}
