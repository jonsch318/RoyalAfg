package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/account/models"
)

func (h *User) Register(rw http.ResponseWriter, r *http.Request) {
	h.l.Info("Register route called")

	var dto RegisterUser
	err := dto.FromJSON(r.Body)
	if err != nil {
		h.l.Error(err)
		http.Error(rw, "The resource could not be decoded", 400)
		return
	}

	h.l.Debug("Decoded user")

	if err := dto.Validate(); err != nil {
		h.l.Error(err)
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(err)
		return
	}

	user, err := models.NewUser(dto.Username, dto.Password, dto.Email)

	if err != nil {
		http.Error(rw, "Something went wrong", http.StatusInternalServerError)
		h.l.Error(err)
		return
	}

	h.l.Debug("User validated")

	if err = h.db.CreateUser(user); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		h.l.Error(err)
		return
	}

	h.l.Debug("User saved")

	token, err := getJwt(user)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		h.l.Error(err)
		return
	}

	cookie := http.Cookie{
		Name:     "identity",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24 * 7),
		HttpOnly: true,
		Path:     "/",
	}

	http.SetCookie(rw, &cookie)
	user.ToJSON(rw)
}

func getJwt(user *models.User) (string, error) {
	signingKey := []byte("testing_key")
	claims := jwt.StandardClaims{
		Subject:   user.ID.Hex(),
		Issuer:    "royalafg.games",
		Audience:  "royalafg.games",
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(),
	}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := at.SignedString(signingKey)

	if err != nil {
		return "", err
	}

	return token, nil
}

type RegisterUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// FromJSON reads from the given io reader and decodes it if possible to a RegisterUser dto, else it returns an error.
func (dto *RegisterUser) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(dto)
}

// Validate validates if the RegisterUser dto matches all the user requirements
func (dto RegisterUser) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.Password, validation.Required, validation.Length(4, 100)),
		validation.Field(&dto.Username, validation.Required, validation.Length(4, 100)),
		validation.Field(&dto.Email, is.Email),
	)
}
