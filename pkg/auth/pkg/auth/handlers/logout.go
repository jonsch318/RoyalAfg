package handlers

import (
	"net/http"
	"time"

	"royalafg/pkg/shared/pkg/mw"
)

// Logout removes the cookie and therefore logs the user out
// swagger:route POST /account/logout account logoutUser
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
//	Security:
//		api_key:
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

	h.l.Debugw("logged out user", "sub", r.Context().Value(mw.KeyUserId{}).(string))

	rw.WriteHeader(http.StatusOK)
	ToJSON(&noContentResponse{}, rw)
}
