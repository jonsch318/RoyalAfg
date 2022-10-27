package handlers

import (
	"net/http"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/mw"
)

// POST request
// Spin: draws a new number and returns the result and prove.
func (s *SlotServer) Spin(rw http.ResponseWriter, r *http.Request) {

	// TODO: CSRF Validattion

	if err := mw.ValidateCSRF(r); err != nil {
		s.l.Errorw("Could not validate csrf token.", "error", err)
		return
	}

	// TODO: Check Permissions and login
	claims := mw.FromUserTokenContext(r.Context().Value("user"))

	

	// TODO: Read request body

	// TODO: Connect to Bank and transact 1€

	// TODO: Draw a new number

	// TODO: Calculate win

	// TODO: Save the result in the buffer/db

	// TODO: Connect to Bank and transact the win

	// TODO: Return the result and prove
}
