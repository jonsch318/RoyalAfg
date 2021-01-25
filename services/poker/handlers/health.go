package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/JohnnyS318/RoyalAfgInGo/services/poker-matchmaker/pkg/lobby"
)

func (h *Game) Health(rw http.ResponseWriter, r *http.Request) {
	players := len(h.lby.Players)

	json.NewEncoder(rw).Encode(lobby.HealthPingResponse{
		Count:   players,
		Running: true,
	})
}
