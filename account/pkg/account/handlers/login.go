package handlers

import (
	"net/http"

	"github.com/JohnnyS318/RoyalAfgInGo/account/pkg/account/models"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// LoginUser defines the object for the api login request
type LoginUser struct {
	Username   string `json:"username" schema:"username"`
	Password   string `json:"password" schema:"password"`
	RememberMe bool   `json:"rememberme" schema:"rememberme"`
}

// Validate validates the LoginUser dto to conform to the api's expectation
func (dto LoginUser) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.Username, validation.Required, validation.Length(4, 100)),
		validation.Field(&dto.Password, validation.Required, validation.Length(4, 100)),
	)
}

// Login authenticates the user and generates a cookie remember this .
// swagger:route POST /account/login authentication
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
func (h *User) Login(rw http.ResponseWriter, r *http.Request) {

	// Set content type header to json
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("X-Content-Type-Options", "nosniff")

	dto := &LoginUser{}

	cType := r.Header.Get("Content-Type")

	h.l.Debug("Content Type", cType)

	switch cType {
	case "application/x-www-form-urlencoded":
		h.l.Debug("content type form urlencoded")

		err := FromFormURLEncodedRequest(dto, r)

		if err != nil {
			h.l.Error("could not decode login dto")
			JSONError(rw, &ErrorResponse{Error: "wrong format decoding failed"}, http.StatusBadRequest)
			return
		}

	case "application/json":
		h.l.Debug("content type json")
		err := FromJSON(dto, r.Body)
		if err != nil {
			h.l.Error("could not decode login dto")
			JSONError(rw, &ErrorResponse{Error: "wrong format decoding failed"}, http.StatusBadRequest)
			return
		}
	}

	// validate dto
	err := dto.Validate()
	if err != nil {
		h.l.Errorf("Validating login dto %s %s", "error", err)
		h.l.Errorw("Validation login dto", "error", err)
		JSONError(rw, &ValidationError{Errors: err}, http.StatusUnprocessableEntity)
		return
	}

	// check wether email or userame was used

	isEmail := validation.Validate(dto.Username, is.EmailFormat) == nil
	if isEmail {
		h.l.Debug("Login using email")
	}

	var user *models.User

	// validate user in db
	if isEmail {
		// validation succeded. sign in using email
		user, err = h.db.FindByEmail(dto.Username)

	} else {
		// email format validation failed. Sign in using username
		user, err = h.db.FindByUsername(dto.Username)
	}

	if err != nil {
		h.l.Error("user with email not found", "error", dto.Username)
		JSONError(rw, &ErrorResponse{Error: "user with specified email not found"}, http.StatusNotFound)
		return
	}

	// validate password

	if !user.ComparePassword(dto.Password, "") {
		h.l.Error("password did not match", "error", dto.Username)
		JSONError(rw, &ErrorResponse{Error: "credentials did not match"}, http.StatusUnauthorized)
		return
	}

	// execute other login schemes (later)
	// validate other schemes (later)

	// create jwt
	token, err := getJwt(user)

	if err != nil {
		h.l.Error("jwt could not be created", "error", err)
		JSONError(rw, &ErrorResponse{Error: "Something went wrong"}, http.StatusInternalServerError)
		return
	}

	// generate Cookie with jwt
	cookie := generateCookie(token, dto.RememberMe)

	// send id cookie with jwt
	http.SetCookie(rw, cookie)

	// Set status code to OK
	rw.WriteHeader(http.StatusOK)

	// send user
	ToJSON(NewUserDTO(user), rw)
}

// VerifyLoggedIn verifies and validates the cookie and it's jwt token. returns 401 if you are not signed in and 200 if everything is valid-
// swagger:route GET /account/verify authentication
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
// 	Responses:
//	default: ErrorResponse
//	401: ErrorResponse
//	200: NoContentResponse
func (h *User) VerifyLoggedIn(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("X-Content-Type-Options", "nosniff")

	rw.WriteHeader(http.StatusOK)
	ToJSON(&noContentResponse{}, rw)
}

// noContentReponse is a empty object.
type noContentResponse struct {
}
