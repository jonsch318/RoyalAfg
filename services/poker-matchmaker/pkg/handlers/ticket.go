package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//TicketResponse is the successful response of a ticket request
type TicketResponse struct {
	Address string `json:"address"`
}

//GetTicketWithParams requests a ticket with lobby params
func (h *Ticket) GetTicketWithParams(rw http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()

	class, err := strconv.Atoi(vals.Get("id"))

	if err != nil {
		http.Error(rw, "Either a valid class or a lobby class has to be given", http.StatusBadRequest)
	}

	addr := ""
	addr, err = h.manager.RequestTicket(class)

	if err != nil {
		h.logger.Errorw("Error during ticket request", "error", err)
		http.Error(rw, "something went wrong during a lobby search", http.StatusInternalServerError)
	}

	json.NewEncoder(rw).Encode(&TicketResponse{Address: addr})

}

//GetTicketWithID requests a ticket with lobby id
func (h *Ticket) GetTicketWithID(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, ok := vars["id"]
	if !ok || id == "" {
		http.Error(rw, "Either a valid class or a lobby class has to be given", http.StatusBadRequest)
	}

	addr, err := h.manager.Connect(id)

	if err != nil {
		h.logger.Errorw("error during connection", "error", err)
		http.Error(rw, "a lobby iwth the given id is not found", http.StatusNotFound)
	}

	json.NewEncoder(rw).Encode(&TicketResponse{Address: addr})

}
