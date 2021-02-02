package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/form3tech-oss/jwt-go"
	"github.com/spf13/viper"

	pokerModels "github.com/JohnnyS318/RoyalAfgInGo/pkg/poker/models"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/events"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/models"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/serviceConfig"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/utils"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		//TODO: CheckOrigin function with configured accepted origins
		return true
	},
}

func (h *Game) Join(rw http.ResponseWriter, r *http.Request) {

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

	//Get all relevant info from jwt token (signed jwt because we do not validate anything here)

	joinEvent, err := events.ToJoinEvent(raw)
	if err != nil {
		log.Printf("joinEvent was invalid %v", err)
		_ = utils.SendToChanTimeout(playerConn.Out, models.NewEvent("VALIDATION_FAILED", "The joining event was not as the server expected"))
		_ = conn.Close()
		http.Error(rw, "VALIDATION_FAILED. The joining event was not as the server expected", http.StatusBadRequest)
		return
	}

	token, err := jwt.Parse(joinEvent.Token, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(viper.GetString(serviceConfig.MatchMakerJWTKey)), nil
	})

	var tokenDec *pokerModels.Token
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid && err == nil {
		log.Printf("Token values: %s, %s, %s", claims["username"], claims["id"], claims["buyIn"])
		tokenDec = pokerModels.FromToken(claims)
	} else {
		log.Printf("joinEvent was invalid %v", err)
		_ = utils.SendToChanTimeout(playerConn.Out, models.NewEvent("VALIDATION_FAILED", "The joining event was not as the server expected"))
		_ = conn.Close()
		http.Error(rw, "VALIDATION_FAILED. The joining event was not as the server expected", http.StatusBadRequest)
		return
	}

	player := models.NewPlayer(tokenDec.Username, tokenDec.Id, tokenDec.BuyIn, playerConn.In, playerConn.Out, playerConn.Close)

	h.lby.Join(player)

	if len(h.lby.Players) == 0 {
		h.stopShutdown <- true
	}
}
