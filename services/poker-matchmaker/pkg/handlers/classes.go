package handlers

import (
	"encoding/json"
	"net/http"
)

// ClassInfo
func (h *Ticket) ClassInfo(rw http.ResponseWriter, r *http.Request) {
	//Get registered poker classes from the poker manager
	classes := h.manager.GetRegisteredClasses()

	//Send response
	err := json.NewEncoder(rw).Encode(classes)
	if err != nil {
		h.logger.Errorf("Could not encode response %v", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
}
