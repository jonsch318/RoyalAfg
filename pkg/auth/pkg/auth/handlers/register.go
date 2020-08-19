package handlers

import (
	"context"
	"net/http"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/auth/pkg/auth/security"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/protos"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/shared/pkg/models"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/shared/pkg/responses"

	"github.com/spf13/viper"
)

// Register registers a new user
// swagger:route POST /account/register authentication registerUser
//
//	Register a new user
//
// This will register a new user with the provided details
//
//	Consumes:
//	-	application/json
//
//	Produces:
//	-	application/json
//
//	Schemes: http, https
//
// 	Responses:
// 		default: ErrorResponse
//		400: ErrorResponse
//		422: ValidationErrorResponse
//		500: ErrorResponse
//		200: UserResponse
//
func (h *User) Register(rw http.ResponseWriter, r *http.Request) {
	h.l.Info("Register route called")

	// Set content type header to json
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("X-Content-Type-Options", "nosniff")

	dto := &RegisterUser{}
	err := FromJSON(dto, r.Body)
	if err != nil {
		h.l.Error("Decoding JSON", "error", err)
		JSONError(rw, &responses.ErrorResponse{Error: "user could not be decoded"}, http.StatusBadRequest)
		return
	}

	h.l.Debug("Decoded user")

	user, err := dto.ToObject()

	if err != nil {
		h.l.Error(err)
		JSONError(rw, &responses.ErrorResponse{Error: "User could not be created"}, http.StatusInternalServerError)
		return
	}

	if err := user.Validate(); err != nil {
		h.l.Error("Validation", "error", err)
		JSONError(rw, &responses.ValidationError{Errors: err}, http.StatusUnprocessableEntity)
		return
	}

	m := protos.ToMessageUser(user)
	_, err = h.userService.SaveUser(context.Background(), m)

	if err != nil {
		h.l.Error(err)
		JSONError(rw, &responses.ErrorResponse{Error: err.Error()}, http.StatusInternalServerError)
		return
	}

	h.l.Debug("User saved")

	token, err := generateBearerToken(user)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		h.l.Error(err)
		return
	}

	cookie := generateCookie(token, dto.RememberMe)

	http.SetCookie(rw, cookie)
	ToJSON(NewUserDTO(user), rw)
}

// RegisterUser defines the dto for the user account registration
type RegisterUser struct {
	Username   string `json:"username"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	FullName   string `json:"fullName"`
	Birthdate  int64  `json:"birthdate"`
	RememberMe bool   `json:"rememberme"`
}

// ToObject converts the RegisterUser dto to the internal user object
func (dto RegisterUser) ToObject() (*models.User, error) {

	user := models.NewUser(dto.Username, dto.Email, dto.FullName, dto.Birthdate)

	hash, err := security.HashPassword(dto.Password, viper.GetString("User.Pepper"))

	if err != nil {
		return nil, err
	}

	user.Hash = hash
	return user, nil
}
