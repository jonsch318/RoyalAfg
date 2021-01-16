package services

import (
	"fmt"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/models"
	"github.com/JohnnyS318/RoyalAfgInGo/services/auth/config"
	"github.com/form3tech-oss/jwt-go"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

func GenerateCookie(token string, persistent bool) *http.Cookie {
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

func GetJwt(user *models.User) (string, error) {
	signingKey := []byte(viper.GetString(config.JwtSigningKey))
	claims := jwt.StandardClaims{
		Subject:   user.ID.Hex(),
		Issuer:    viper.GetString(config.JwtIssuer),
		Audience:  []string{"github.com/JohnnyS318/RoyalAfgInGo.games"},
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

func GenerateBearerToken(user *models.User) (string, error) {
	token, err := GetJwt(user)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Bearer %v", token), nil
}
