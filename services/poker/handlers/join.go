package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/events"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/models"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/utils"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *Game) Join(rw http.ResponseWriter, r *http.Request) {

	//Check if lobby is configured
	if h.lby == nil {
		log.Printf("lobby is nil, because the game servers annotations where not changed yet")
	}

	//Check if lobby is configured
	if h.lby == nil {
		log.Printf("lobby is nil, because the game servers annotations where not changed yet")
	}

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
		_ = utils.SendToChanTimeout(playerConn.Out, models.NewEvent("VALIDATION_FAILED", "The joining event was not as the server expected"))
		_ = conn.Close()
		http.Error(rw, "VALIDATION_FAILED. The joining event was not as the server expected", http.StatusBadRequest)
		return
	}

	player := models.NewPlayer(joinEvent.Username, joinEvent.ID, joinEvent.BuyIn, playerConn.In, playerConn.Out, playerConn.Close)

	h.lby.Join(player)

}
