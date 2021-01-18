package services

import (
	"fmt"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/errors"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/models"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/mw"
	"github.com/JohnnyS318/RoyalAfgInGo/services/auth/config"
	gConfig "github.com/JohnnyS318/RoyalAfgInGo/pkg/config"
	"github.com/form3tech-oss/jwt-go"
	"github.com/spf13/viper"
"github.com/google/uuid"
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

	if !persistent {
		cookie.Expires = viper.GetTime(config.CookieExpires)
	}

	return cookie
}

func GetJwt(user *models.User) (string, error) {
	signingKey := []byte(viper.GetString(gConfig.JWTSigningKey))

	claims := jwt.MapClaims{
		"sub": user.ID,
		"iss": viper.GetString(gConfig.JWTIssuer),
		"aud": []string{"royalafg.games", "localhost:3000"},
		"exp": time.Now().Add(viper.GetDuration(gConfig.JWTExpiresAt)),
		"jti": uuid.New().String(),
		"username": user.Username,
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
	token, err := jwt.Parse(val, mw.GetKeyGetter(viper.GetString(gConfig.JWTSigningKey)))

	if err != nil || !token.Valid {
		return nil, "", errors.InvalidTokenError{}
	}

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = getExpiration()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := at.SignedString([]byte(viper.GetString(gConfig.JWTSigningKey)))
	return token, tokenString, err
}

func getExpiration() time.Time {
	return time.Now().Add(viper.GetDuration(gConfig.JWTExpiresAt))
}