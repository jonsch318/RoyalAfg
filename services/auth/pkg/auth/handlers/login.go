package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/models"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/responses"
	"github.com/JohnnyS318/RoyalAfgInGo/services/auth/pkg/auth/security"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/protos"
	"go.mongodb.org/mongo-driver/bson/primitive"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/spf13/viper"
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
			h.l.Errorw("could not decode login dto", "error", err)
			JSONError(rw, &responses.ErrorResponse{Error: "wrong format decoding failed"}, http.StatusBadRequest)
			return
		}

	case "application/json":
		h.l.Debug("content type json")
		err := FromJSON(dto, r.Body)
		if err != nil {
			h.l.Errorw("could not decode login dto", "error", err)
			JSONError(rw, &responses.ErrorResponse{Error: "wrong format decoding failed"}, http.StatusBadRequest)
			return
		}
	}

	// validate dto
	err := dto.Validate()
	if err != nil {
		h.l.Errorw("Validation login dto", "error", err)
		JSONError(rw, &responses.ValidationError{Errors: err}, http.StatusUnprocessableEntity)
		return
	}

	// check wether email or userame was used

	isEmail := validation.Validate(dto.Username, is.EmailFormat) == nil
	if isEmail {
		h.l.Debug("Login using email")
	}

	m, err := h.userService.GetUserByUsername(context.Background(), &protos.GetUser{Identifier: dto.Username})

	user := &models.User{
		Username:  m.Username,
		Email:     m.Email,
		Birthdate: m.Birthdate,
		FullName:  m.FullName,
		Hash:      m.Hash,
	}

	id, _ := primitive.ObjectIDFromHex(m.Id)

	user.ID = id

	user.CreatedAt = time.Unix(m.CreatedAt, 0)
	user.UpdatedAt = time.Unix(m.UpdatedAt, 0)

	if err != nil {
		h.l.Errorw("user with username or email not found", "error", dto.Username)
		JSONError(rw, &responses.ErrorResponse{Error: "user with username or email not found"}, http.StatusNotFound)
		return
	}

	// validate password
	if !security.ComparePassword(dto.Password, user.Hash, viper.GetString("User.Pepper")) {
		h.l.Errorw("password did not match", "error", dto.Username)
		JSONError(rw, &responses.ErrorResponse{Error: "credentials did not match"}, http.StatusUnauthorized)
		return
	}

	// execute other login schemes (later)
	// validate other schemes (later)

	// create jwt
	token, err := generateBearerToken(user)

	if err != nil {
		h.l.Errorw("jwt could not be created", "error", err)
		JSONError(rw, &responses.ErrorResponse{Error: "Something went wrong"}, http.StatusInternalServerError)
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
func (h *User) VerifyLoggedIn(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("X-Content-Type-Options", "nosniff")

	rw.WriteHeader(http.StatusOK)
	ToJSON(&noContentResponse{}, rw)
}

// noContentReponse is a empty object.
type noContentResponse struct {
}
