package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/jonsch318/royalafg/services/poker-matchmaker/pkg/lobby"
)

func (h *Game) Health(rw http.ResponseWriter, r *http.Request) {
	players := len(h.lby.Players)

	json.NewEncoder(rw).Encode(lobby.HealthPingResponse{
		Count:   players,
		Running: true,
		LobbyID: h.lby.LobbyID,
	})
}
