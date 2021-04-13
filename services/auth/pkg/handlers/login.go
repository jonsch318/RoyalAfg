package handlers

import (
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/auth"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/dtos"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/mw"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/responses"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/utils"
)

// Validate validates the LoginDto dto to conform to the api's expectation
func Validate(dto dtos.LoginDto) error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.Username, validation.Required, validation.Length(4, 100)),
		validation.Field(&dto.Password, validation.Required, validation.Length(4, 100)),
	)
}

// Login authenticates the user and generates a cookie remember this .
// swagger:route POST /api/auth/login authentication loginUser
//
// Login to a user account
//
// After matching credentials, this will generate a jwt and pass it on as a cookie
//
//	Consumes:
//	-	application/json
//	-	application/x-www-form-urlencoded
//
// 	Produces:
//	-	application/json
//
//	Security:
//	-	api_key
//
//	Schemes: http, https
//
// 	Responses:
//	default: ErrorResponse
//	400: ErrorResponse
//	422: ValidationErrorResponse
//	404: ErrorResponse
//	401: ErrorResponse
//	403: ErrorResponse
//	500: ErrorResponse
//	200: UserResponse
//
func (h *Auth) Login(rw http.ResponseWriter, r *http.Request) {

	h.l.Infof("Origin %v", r.RemoteAddr)
	// Set content type header to json
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("X-Content-Type-Options", "nosniff")

	if err := mw.ValidateCSRF(r); err != nil {
		h.l.Errorw("could not validate csrf token", "error", err)
		responses.JSONError(rw, &responses.ErrorResponse{Error: "wrong format decoding failed"}, http.StatusForbidden)
		return
	}

	loginDto := &dtos.LoginDto{}

	cType := r.Header.Get("Content-Type")

	h.l.Debug("Content Type", cType)

	switch cType {
	case "application/x-www-form-urlencoded":
		h.l.Debug("content type form urlencoded")

		err := utils.FromFormURLEncodedRequest(loginDto, r)

		if err != nil {
			h.l.Errorw("could not decode login dto", "error", err)
			responses.JSONError(rw, &responses.ErrorResponse{Error: "wrong format decoding failed"}, http.StatusBadRequest)
			return
		}

	case "application/json":
		h.l.Debug("content type json")
		err := utils.FromJSON(loginDto, r.Body)
		if err != nil {
			h.l.Errorw("could not decode login dto", "error", err)
			responses.JSONError(rw, &responses.ErrorResponse{Error: "wrong format decoding failed"}, http.StatusBadRequest)
			return
		}
	}

	h.l.Infof("Dto: %v", loginDto)

	// validate dto
	err := Validate(*loginDto)
	if err != nil {
		h.l.Errorw("Validation login dto", "error", err)
		responses.JSONError(rw, &responses.ValidationError{Errors: err}, http.StatusUnprocessableEntity)
		return
	}

	user, token, err := h.Auth.Login(loginDto.Username, loginDto.Password)

	if err != nil {
		h.l.Errorw("jwt could not be created", "error", err)
		responses.JSONError(rw, &responses.ErrorResponse{Error: "Something went wrong"}, http.StatusInternalServerError)
		return
	}

	// generate Cookie with jwt
	cookie := auth.GenerateCookie(token, loginDto.RememberMe)

	// send id cookie with jwt
	http.SetCookie(rw, cookie)

	// Set status code to OK
	rw.WriteHeader(http.StatusOK)

	// send user
	err = utils.ToJSON(dtos.NewUser(user), rw)
	if err != nil {
		h.l.Errorw("json serialization", "error", err)
		responses.JSONError(rw, &responses.ErrorResponse{Error: "Something went wrong"}, http.StatusInternalServerError)
		return
	}
}

// noContentResponse is a empty object.
type noContentResponse struct {
}
