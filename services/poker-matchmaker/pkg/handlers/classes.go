package handlers

import (
	"encoding/json"
	"net/http"
)

//ClassInfo
func (h *Ticket) ClassInfo(rw http.ResponseWriter, r *http.Request) {
	//Get registered poker classes from the poker manager
	classes := h.manager.GetRegisteredClasses()

	//Send response
	_ = json.NewEncoder(rw).Encode(classes)
}
