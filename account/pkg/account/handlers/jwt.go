package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/JohnnyS318/RoyalAfgInGo/account/pkg/account/models"
	"github.com/dgrijalva/jwt-go"
)

func generateCookie(token string, persistent bool) *http.Cookie {
	cookie := &http.Cookie{
		Name:     "identity",
		Value:    token,
		HttpOnly: true,
		Path:     "/",
	}

	if persistent {
		cookie.Expires = time.Now().Add(time.Hour * 24 * 7)
	}

	return cookie
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

func validateJwt(tokenString string) (jwt.MapClaims, error) {
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte("testing_key"), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		err = claims.Valid()
		if err != nil {
			return nil, err
		}

		return claims, nil
	}
	return nil, fmt.Errorf("The token validation failed")
}
