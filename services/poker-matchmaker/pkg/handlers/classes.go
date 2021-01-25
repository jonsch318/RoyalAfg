package handlers

import (
	"encoding/json"
	"net/http"
)

//ClassInfo
func (h *Ticket) ClassInfo(rw http.ResponseWriter, r *http.Request) {
	classes := h.manager.GetRegisteredClasses()
	_ = json.NewEncoder(rw).Encode(classes)
}
