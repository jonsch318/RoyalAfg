package handlers

import (
	"net/http"
	"time"

	"github.com/jonsch318/royalafg/pkg/auth"
	"github.com/jonsch318/royalafg/pkg/responses"
	"github.com/jonsch318/royalafg/pkg/utils"

	"github.com/jonsch318/royalafg/pkg/mw"
)

// Logout removes the cookie and therefore logs the user out
// swagger:route POST /account/logout account logoutUser
//
// # Logout of account
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
func (h *Auth) Logout(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("X-Content-Type-Options", "nosniff")
	user := mw.FromUserTokenContext(r.Context().Value("user"))
	cookie := auth.GenerateCookie("", false)

	err := h.Auth.Logout(user.ID)

	if err != nil {
		h.l.Errorw("error during logout process", "error", err)
		responses.JSONError(rw, &responses.ErrorResponse{Error: "error during logout process"}, http.StatusInternalServerError)
		return
	}

	cookie.Expires = time.Unix(0, 0)

	http.SetCookie(rw, cookie)

	h.l.Debugw("logged out user", "id", user.ID)

	rw.WriteHeader(http.StatusOK)
	err = utils.ToJSON(&noContentResponse{}, rw)
	if err != nil {
		h.l.Errorw("json serialization", "error", err)
		responses.JSONError(rw, &responses.ErrorResponse{Error: "Something went wrong"}, http.StatusInternalServerError)
		return
	}
}
