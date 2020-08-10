package handlers

import (
	"context"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

type KeyUserId struct{}
type KeyJWTClaims struct{}

func (h *User) AuthMW(next http.Handler) http.Handler {
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

func (h *User) RequireAuthTokenHandler(rw http.ResponseWriter, r *http.Request) (jwt.MapClaims, error) {
	cookie, err := r.Cookie("identity")

	if err != nil {
		JSONError(rw, &ErrorResponse{Error: "You are not logged in"}, http.StatusUnauthorized)
		return nil, err
	}

	claims, err := validateJwt(cookie.Value)

	if err != nil {
		JSONError(rw, &ErrorResponse{Error: "You're login is not valid. We will sign you out"}, http.StatusUnauthorized)
		return nil, err
	}

	return claims, nil
}
