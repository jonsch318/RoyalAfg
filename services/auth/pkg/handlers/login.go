package handlers

import (
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/responses"
	"github.com/JohnnyS318/RoyalAfgInGo/services/auth/pkg/dto"
	"github.com/JohnnyS318/RoyalAfgInGo/services/auth/pkg/services"
	"github.com/go-ozzo/ozzo-validation"
	"net/http"
)

// Validate validates the LoginDto dto to conform to the api's expectation
func Validate(dto *dto.LoginDto) error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.Username, validation.Required, validation.Length(4, 100)),
		validation.Field(&dto.Password, validation.Required, validation.Length(4, 100)),
	)
}

// Login authenticates the user and generates a cookie remember this .
// swagger:route POST /account/login authentication loginUser
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
//	Schemes: http, https
//
// 	Responses:
//	default: ErrorResponse
//	400: ErrorResponse
//	422: ValidationErrorResponse
//	404: ErrorResponse
//	401: ErrorResponse
//	500: ErrorResponse
//	200: UserResponse
//
func (h *Auth) Login(rw http.ResponseWriter, r *http.Request) {

	// Set content type header to json
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("X-Content-Type-Options", "nosniff")

	loginDto := &dto.LoginDto{}

	cType := r.Header.Get("Content-Type")

	h.l.Debug("Content Type", cType)

	switch cType {
	case "application/x-www-form-urlencoded":
		h.l.Debug("content type form urlencoded")

		err := FromFormURLEncodedRequest(loginDto, r)

		if err != nil {
			h.l.Errorw("could not decode login dto", "error", err)
			responses.JSONError(rw, &responses.ErrorResponse{Error: "wrong format decoding failed"}, http.StatusBadRequest)
			return
		}

	case "application/json":
		h.l.Debug("content type json")
		err := FromJSON(loginDto, r.Body)
		if err != nil {
			h.l.Errorw("could not decode login dto", "error", err)
			responses.JSONError(rw, &responses.ErrorResponse{Error: "wrong format decoding failed"}, http.StatusBadRequest)
			return
		}
	}

	// validate dto
	err := Validate(loginDto)
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
	cookie := services.GenerateCookie(token, loginDto.RememberMe)

	// send id cookie with jwt
	http.SetCookie(rw, cookie)

	// Set status code to OK
	rw.WriteHeader(http.StatusOK)

	// send user
	err = ToJSON(dto.NewUserDTO(user), rw)
	if err != nil {
		h.l.Errorw("json serialization", "error", err)
		responses.JSONError(rw, &responses.ErrorResponse{Error: "Something went wrong"}, http.StatusInternalServerError)
		return
	}
}

// VerifyLoggedIn verifies and validates the cookie and it's jwt token. returns 401 if you are not signed in and 200 if everything is valid-
// swagger:route GET /account/verify authentication account verifyLoggedIn
//
// Verify that the user is logged in
//
// This will return either status code 401 Unauthorized if user is not signed in and 200 when the login token is valid
//
//	Consumes:
//
// 	Produces:
//	-	application/json
//
//	Schemes: http, https
//
//	Security:
//
// 	Responses:
//	default: ErrorResponse
//	401: ErrorResponse
//	200: NoContentResponse
func (h *Auth) VerifyLoggedIn(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("X-Content-Type-Options", "nosniff")

	authenticated, err := h.Auth.VerifyAuthentication()

	if !authenticated || err != nil {
		h.l.Errorw("A error during login verification", "error", err)
		rw.WriteHeader(http.StatusUnauthorized)
	}

	rw.WriteHeader(http.StatusOK)
	err = ToJSON(&noContentResponse{}, rw)
	if err != nil {
		h.l.Errorw("json serialization", "error", err)
		responses.JSONError(rw, &responses.ErrorResponse{Error: "Something went wrong"}, http.StatusInternalServerError)
		return
	}
}

// noContentResponse is a empty object.
type noContentResponse struct {
}
