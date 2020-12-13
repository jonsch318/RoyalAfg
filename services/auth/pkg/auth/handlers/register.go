package handlers

import (
	"context"
	"net/http"

	"github.com/JohnnyS318/RoyalAfgInGo/services/auth/pkg/auth/security"
	"github.com/JohnnyS318/RoyalAfgInGo/services/protos"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/models"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/responses"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

	user, err := dto.ToObject()

	h.l.Debug("Decoded newly registered user")

	if err != nil {
		h.l.Errorw("Error during dto eval and hashing", "error", err)
		JSONError(rw, &responses.ErrorResponse{Error: "User could not be created"}, http.StatusInternalServerError)
		return
	}

	if err := user.Validate(); err != nil {
		h.l.Error("Validation failed", "error", err)
		JSONError(rw, &responses.ValidationError{Errors: err}, http.StatusUnprocessableEntity)
		return
	}

	msg := protos.ToMessageUser(user)

	m, err := h.userService.SaveUser(context.Background(), msg)

	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			h.l.Errorw("Error from grpc conn", "error", err, "status", st)
		} else {
			h.l.Errorw("Error from grpc conn", "error", err)
		}

		switch st.Code() {
		case codes.InvalidArgument:
			h.l.Errorw("Validation", "error", st.Err())
			JSONError(rw, &responses.ValidationError{Errors: st.Message()}, http.StatusUnprocessableEntity)
			return
		case codes.Internal:
			h.l.Errorw("UserService Call Internal", "error", st.Err())
			JSONError(rw, &responses.ErrorResponse{Error: st.Message()}, http.StatusInternalServerError)
			return
		case codes.AlreadyExists:
			h.l.Errorw("Validation", "error", st.Err())
			JSONError(rw, &responses.ValidationError{Errors: st.Message()}, http.StatusUnprocessableEntity)
			return
		}

		JSONError(rw, &responses.ErrorResponse{Error: err.Error()}, http.StatusInternalServerError)
		return
	}

	user = protos.FromMessageUserExact(m)

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

	hash, err := security.HashPassword(dto.Password, "")

	if err != nil {
		return nil, err
	}

	user.Hash = hash
	return user, nil
}
