package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/dtos"
)

//LobbyInfo returns 15 combined possible Lobbies of all classes. Used for an Explore Lobbies functionality
func (h *Ticket) PokerInfo(rw http.ResponseWriter, r *http.Request)  {
	lobbies := h.manager.GetRegisteredLobbies(15)
	classes := h.manager.GetRegisteredClasses()

	_ = json.NewEncoder(rw).Encode(&dtos.PokerInfoResponse{Lobbies: lobbies, Classes: classes})
}
