package handlers

import (
	"net/http"
	"time"

	"royalafg/pkg/auth/pkg/auth/config"
	"royalafg/pkg/shared/pkg/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

func generateCookie(token string, persistent bool) *http.Cookie {
	cookie := &http.Cookie{
		Name:     viper.GetString(config.CookieName),
		Value:    token,
		HttpOnly: true,
		Path:     "/",
	}

	if persistent {
		cookie.Expires = viper.GetTime(config.CookieExpires)
	}

	return cookie
}

func getJwt(user *models.User) (string, error) {
	signingKey := []byte(viper.GetString(config.JwtSigningKey))
	claims := jwt.StandardClaims{
		Subject:   user.ID.Hex(),
		Issuer:    viper.GetString(config.JwtIssuer),
		Audience:  "royalafg.games",
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: viper.GetInt64(config.JwtExpiresAt),
	}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := at.SignedString(signingKey)

	if err != nil {
		return "", err
	}

	return token, nil
}
