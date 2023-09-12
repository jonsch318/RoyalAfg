package handlers

import (
	"net/http"

	"github.com/jonsch318/royalafg/pkg/dtos"
	"github.com/jonsch318/royalafg/pkg/utils"
)

// PokerInfo returns 15 combined possible lobbies of all classes and the registered poker classes
func (h *Ticket) PokerInfo(rw http.ResponseWriter, r *http.Request) {

	//get lobbies
	lobbies := h.manager.GetRegisteredLobbies(15)
	//get classes
	classes := h.manager.GetRegisteredClasses()

	//Send results
	_ = utils.ToJSON(&dtos.PokerInfoResponse{
		Lobbies: lobbies,
		Classes: classes,
	}, rw)
}
