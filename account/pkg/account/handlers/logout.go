package handlers

import (
	"net/http"
	"time"
)

// Logout removes the cookie and therefore logs the user out
// swagger:route POST /account/logout authentication
//
// Logout of account
//
// This will remove the identity cookie
//
//	Consumes:
//
//	Produces:
//	-	application/json
//
//	Schemes: http, https
//
//	Responses:
//	default: ErrorResponse
//	401: ErrorResponse
//	200: NoContentResponse
//
func (h *User) Logout(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("X-Content-Type-Options", "nosniff")

	cookie := generateCookie("", false)

	cookie.Expires = time.Unix(0, 0)

	http.SetCookie(rw, cookie)

	h.l.Debugw("logged out user", "sub", r.Context().Value(KeyUserId{}).(string))

	rw.WriteHeader(http.StatusOK)
	ToJSON(&noContentResponse{}, rw)
}
