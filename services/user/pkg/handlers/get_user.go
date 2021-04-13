package handlers

import (
	"net/http"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/dtos"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/mw"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/responses"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/utils"
)

type GetUserResponse struct {
	User *dtos.User `json:"user"`

}

// ErrorResponse is a generic error response
// swagger:response ErrorResponse
type errorResponseWrapper struct {
	// The error
	// in: body
	Body responses.ErrorResponse
}

// getUserWrapper returns the current user information.
// swagger:response UserResponse
type userResponseWrapper struct {
	// The user.
	// in: body
	Body GetUserResponse
}

// GetUser returns the authenticated user object
// swagger:route GET /api/user account getUser
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
	claims := mw.FromUserTokenContext(r.Context().Value("user"))

	user, err := h.db.FindById(claims.ID)

	if err != nil {
		h.l.Errorw("User not found", "error", err, "id", claims.ID)
		responses.JSONError(rw, &responses.ErrorResponse{Error: "User not found"}, http.StatusNotFound)
		return
	}

	rw.WriteHeader(200)
	_ = utils.ToJSON(&GetUserResponse{User: dtos.NewUser(user)}, rw)
	return
}
