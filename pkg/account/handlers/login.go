package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/account/models"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type LoginUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (dto *LoginUser) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(dto)
}

func (dto LoginUser) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.Username, validation.Required, validation.Length(4, 100)),
		validation.Field(&dto.Password, validation.Required, validation.Length(4, 100)),
	)
}

func (h *User) Login(rw http.ResponseWriter, r *http.Request) {
	// until logger middleware logging in each handler
	h.l.Debug("Login called")

	// decode json dto
	var dto LoginUser
	err := dto.FromJSON(r.Body)

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

	// send user

	// send id cookie with jwt
}
