package handlers

import (
	"net/http"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/shared/pkg/mw"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/shared/pkg/responses"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/shared/pkg/utils"
)

// GetUser returns the authenticated user object
// swagger:route GET /account account getUser
//	return the authenticated user based on the api key
//	Consumes:
//
//	Produces:
//	-	application/json
//	Schemes: http, https
//
//	Security:
//		api_key:
//
//	Responses:
//	default: ErrorResponse
//	401: ErrorResponse
// 	404: ErrorResponse
//	500: ErrorResponse
//	200: UserResponse
//
//
func (h *UserHandler) GetUser(rw http.ResponseWriter, r *http.Request) {
	id, ok := r.Context().Value(mw.KeyUserId{}).(string)

	if ok && id != "" {
		h.l.Errorw("User Id after authmv empty", "error", "User Id after authmv empty")
		responses.JSONError(rw, &responses.ErrorResponse{Error: "User could not be recognized or is not authenticated"}, http.StatusUnauthorized)
		return
	}

	user, err := h.db.FindById(id)

	if err != nil {
		h.l.Errorw("User not found", "error", err, "id", id)
		responses.JSONError(rw, &responses.ErrorResponse{Error: "User not found"}, http.StatusNotFound)
		return
	}

	rw.WriteHeader(200)
	err = utils.ToJSON(user, rw)

	if err != nil {
		h.l.Errorw("User could get decoded", "error", err)
		responses.JSONError(rw, &responses.ErrorResponse{Error: "Something went wrong while decoding the user"}, http.StatusInternalServerError)
		return
	}

}
