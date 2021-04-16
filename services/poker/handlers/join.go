package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/form3tech-oss/jwt-go"
	"github.com/spf13/viper"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/log"
	pokerModels "github.com/JohnnyS318/RoyalAfgInGo/pkg/poker/models"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/poker/ticket"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/events"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/models"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/serviceconfig"
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
		log.Logger.Errorf("lobby is nil, because the game servers annotations where not changed yet")
	}

	conn, err := upgrader.Upgrade(rw, r, nil)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadGateway)
	}

	//create player connection abstraction. Used to send and receive messages.
	playerConn := NewPlayerConn(conn)
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
		log.Logger.Warnw("joinEvent was invalid", "error", err)
		_ = utils.SendToChanTimeout(playerConn.Out, models.NewEvent("VALIDATION_FAILED", "The joining event was not as the server expected"))
		playerConn.CloseConnection(false)
		http.Error(rw, "VALIDATION_FAILED. The joining event was not as the server expected", http.StatusBadRequest)
		return
	}

	token, err := jwt.Parse(joinEvent.Token, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(viper.GetString(serviceconfig.MatchMakerJWTKey)), nil
	})

	if err != nil {
		log.Logger.Warnw("joinEvent was invalid", "error", err)
		_ = utils.SendToChanTimeout(playerConn.Out, models.NewEvent("VALIDATION_FAILED", "The joining event was not as the server expected"))
		playerConn.CloseConnection(false)
		http.Error(rw, "VALIDATION_FAILED. The joining event was not as the server expected", http.StatusBadRequest)
		return
	}

	var tokenDec *pokerModels.Token
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid && err == nil {
		log.Logger.Debugf("Token values: %s, %s, %s", claims["username"], claims["id"], claims["buyIn"])
		tokenDec = ticket.FromToken(claims)
	} else {
		log.Logger.Warnw("joinEvent was invalid", "error", err)
		_ = utils.SendToChanTimeout(playerConn.Out, models.NewEvent("VALIDATION_FAILED", "The joining event was not as the server expected"))
		playerConn.CloseConnection(false)
		http.Error(rw, "VALIDATION_FAILED. The joining event was not as the server expected", http.StatusBadRequest)
		return
	}

	player := models.NewPlayer(tokenDec.Username, tokenDec.Id, tokenDec.BuyIn, playerConn.In, playerConn.Out, playerConn.Close)

	log.Logger.Debugf("joining player to lobby")
	err = h.lby.Join(player)

	if err != nil {
		log.Logger.Warnw("join was unsuccessful", "error", err)
		_ = utils.SendToChanTimeout(playerConn.Out, models.NewEvent("VALIDATION_FAILED", "error during joining process"))
		playerConn.CloseConnection(false)
		http.Error(rw, "VALIDATION_FAILED. Error during joining process", http.StatusBadRequest)
		return
	}

	if h.lby.Count() == 0 {
		log.Logger.Debugf("stopping shutdown")
		h.stopShutdown <- true
	}
}
