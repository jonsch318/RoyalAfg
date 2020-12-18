package handlers

import (
	"encoding/json"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/responses"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/events"
	"net/http"
)

func (l *Lobby) LobbyOptions(rw http.ResponseWriter, r *http.Request) {

	rw.Header().Set("Access-Control-Allow-Origin", "*")

	//log.Printf("/options")

	encoder := json.NewEncoder(rw)

	classes := make([][]float32, len(l.Lobbies.BuyInClasses))

	for i := range classes {
		classes[i] = make([]float32, len(l.Lobbies.BuyInClasses[i]))
		for j := range classes[i] {
			classes[i][j] = float32(l.Lobbies.BuyInClasses[i][j]) / 100
		}
	}

	o := &events.JoinOptions{
		BuyInClasses: classes,
		Lobbies:      l.Lobbies.GetAllLobbies(),
	}
	err := encoder.Encode(o)
	if err != nil {
		responses.JSONError(rw, responses.ErrorResponse{Error: "Error during Encoding"}, http.StatusInternalServerError)
	}
}
