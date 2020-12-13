package mw

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/responses"

	"github.com/dgrijalva/jwt-go"
	"go.uber.org/zap"
)

type KeyJWTClaims struct{}
type KeyUserId struct{}

type AuthMWHandler struct {
	l   *zap.SugaredLogger
	key string
}

func NewAuthMWHandler(logger *zap.SugaredLogger, key string) *AuthMWHandler {
	return &AuthMWHandler{
		l:   logger,
		key: key,
	}
}

func (h *AuthMWHandler) AuthMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		claims, err := h.RequireAuthTokenHandler(rw, r)

		if err != nil {
			return
		}

		idCtx := context.WithValue(r.Context(), KeyUserId{}, claims["sub"])
		ctx := context.WithValue(idCtx, KeyJWTClaims{}, claims)

		next.ServeHTTP(rw, r.WithContext(ctx))

	})
}

func (h *AuthMWHandler) RequireAuthTokenHandler(rw http.ResponseWriter, r *http.Request) (jwt.MapClaims, error) {
	cookie, err := r.Cookie("identity")

	if err != nil {
		responses.JSONError(rw, &responses.ErrorResponse{Error: "You are not logged in"}, http.StatusUnauthorized)
		return nil, err
	}

	claims, err := ValidateJwt(cookie.Value, h.key)

	if err != nil {
		responses.JSONError(rw, &responses.ErrorResponse{Error: "You're login is not valid. We will sign you out"}, http.StatusUnauthorized)
		return nil, err
	}

	return claims, nil
}

func ValidateJwt(bearer, key string) (jwt.MapClaims, error) {

	tokenString := strings.Split(bearer, "Bearer ")[1]

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
		return []byte(key), nil
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
