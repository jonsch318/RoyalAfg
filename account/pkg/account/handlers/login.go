package handlers

import (
	"net/http"

	"github.com/JohnnyS318/RoyalAfgInGo/account/pkg/account/models"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// LoginUser defines the object for the api login request
type LoginUser struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	RememberMe bool   `json:"rememberme"`
}

// Validate validates the LoginUser dto to conform to the api's expectation
func (dto LoginUser) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.Username, validation.Required, validation.Length(4, 100)),
		validation.Field(&dto.Password, validation.Required, validation.Length(4, 100)),
	)
}

// Login authenticates the user and generates a cookie remember this .
func (h *User) Login(rw http.ResponseWriter, r *http.Request) {

	// Set content type header to json
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("X-Content-Type-Options", "nosniff")

	// until logger middleware logging in each handler
	h.l.Debug("Login called")

	// decode json dto
	var dto LoginUser
	err := FromJSON(dto, r.Body)

	if err != nil {
		h.l.Error("could not decode login dto")
		JSONError(rw, &ErrorResponse{Error: "wrong format decoding failed"}, http.StatusBadGateway)
		return
	}

	// validate dto
	err = dto.Validate()
	if err != nil {
		h.l.Error("Validating login dto", "error", err)
		JSONError(rw, &ValidationError{Errors: err}, http.StatusUnprocessableEntity)
		return
	}

	// check wether email or userame was used

	isEmail := validation.Validate(dto.Username, is.EmailFormat)
	if isEmail != nil {
		h.l.Debug("Login using email")
	}

	var user *models.User

	// validate user in db
	if isEmail != nil {
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

	// send user
	ToJSON(user, rw)

	// Set status code to OK
	rw.WriteHeader(http.StatusOK)
}

// VerifyLoggedIn verifies and validates the cookie and it's jwt token. returns 401 if you are not signed in and 200 if everything is valid-
func (h *User) VerifyLoggedIn(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("X-Content-Type-Options", "nosniff")

	cookie, err := r.Cookie("identity")

	if err != nil {
		JSONError(rw, &ErrorResponse{Error: "identity cookie not found"}, http.StatusUnauthorized)
		return
	}

	_, err = validateJwt(cookie.Value)

	if err != nil {
		JSONError(rw, &ErrorResponse{Error: "jwt not valid"}, http.StatusUnauthorized)
		return
	}

	ToJSON(&noContentResponse{}, rw)
	rw.WriteHeader(http.StatusOK)
}

// noContentReponse is a empty object.
type noContentResponse struct {
}
