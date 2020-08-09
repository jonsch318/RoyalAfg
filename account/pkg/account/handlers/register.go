package handlers

import (
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"

	"github.com/JohnnyS318/RoyalAfgInGo/account/pkg/account/models"
)

// Register registers a new user
func (h *User) Register(rw http.ResponseWriter, r *http.Request) {
	h.l.Info("Register route called")

	// Set content type header to json
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("X-Content-Type-Options", "nosniff")

	var dto RegisterUser
	err := FromJSON(dto, r.Body)
	if err != nil {
		h.l.Error(err)
		JSONError(rw, &ErrorResponse{Error: "user could not be decoded"}, http.StatusBadRequest)
		return
	}

	h.l.Debug("Decoded user")

	if err := dto.Validate(); err != nil {
		h.l.Error(err)
		JSONError(rw, &ValidationError{Errors: err}, http.StatusUnprocessableEntity)
		return
	}

	user, err := models.NewUser(dto.Username, dto.Password, dto.Email)

	if err != nil {
		h.l.Error(err)
		JSONError(rw, &ErrorResponse{Error: "User could not be created"}, http.StatusInternalServerError)
		return
	}

	h.l.Debug("User validated")

	if err = h.db.CreateUser(user); err != nil {
		h.l.Error(err)
		JSONError(rw, &ErrorResponse{Error: err.Error()}, http.StatusInternalServerError)
		return
	}

	h.l.Debug("User saved")

	token, err := getJwt(user)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		h.l.Error(err)
		return
	}

	cookie := generateCookie(token, dto.RememberMe)

	http.SetCookie(rw, cookie)
	user.ToJSON(rw)
}

// RegisterUser defines the dto for the user account registration
type RegisterUser struct {
	Username   string `json:"username"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	RememberMe bool   `json:"rememberme"`
}

// Validate validates if the RegisterUser dto matches all the user requirements
func (dto RegisterUser) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.Password, validation.Required, validation.Length(4, 100)),
		validation.Field(&dto.Username, validation.Required, validation.Length(4, 100)),
		validation.Field(&dto.Email, is.Email),
	)
}
