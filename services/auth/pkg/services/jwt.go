package services

import (
	"fmt"
	"net/http"
	"time"

	"github.com/form3tech-oss/jwt-go"
	"github.com/google/uuid"
	"github.com/spf13/viper"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/config"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/models"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/mw"
)

func GenerateCookie(token string, persistent bool) *http.Cookie {
	cookie := &http.Cookie{
		Name:     viper.GetString(config.SessionCookieName),
		Value:    token,
		HttpOnly: true,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
	}

	if !persistent {
		cookie.Expires = getExpiration(config.SessionCookieExpiration)
	}

	return cookie
}

func GetJwt(user *models.User) (string, error) {
	signingKey := []byte(viper.GetString(config.JWTSigningKey))

	claims := jwt.MapClaims{
		"sub":      user.ID,
		"iss":      viper.GetString(config.JWTIssuer),
		"aud":      []string{"royalafg.games", "localhost:3000"},
		"exp":      time.Now().Add(viper.GetDuration(config.JWTExpiresAt)),
		"jti":      uuid.New().String(),
		"username": user.Username,
		"name":     user.FullName,
	}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := at.SignedString(signingKey)

	if err != nil {
		return "", err
	}

	return token, nil
}

func GenerateBearerToken(user *models.User) (string, error) {
	token, err := GetJwt(user)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Bearer %v", token), nil
}

func ExtendToken(val string) (*jwt.Token, string, error) {
	token, err := jwt.Parse(val, mw.GetKeyGetter(viper.GetString(config.JWTSigningKey)))

	if err != nil || !token.Valid {
		return nil, "", err
	}

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = getExpiration(config.JWTExpiresAt)

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := at.SignedString([]byte(viper.GetString(config.JWTSigningKey)))
	return token, tokenString, err
}

func getExpiration(key string) time.Time {
	return time.Now().Add(viper.GetDuration(config.JWTExpiresAt))
}
