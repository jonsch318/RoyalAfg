package mw

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/form3tech-oss/jwt-go"
	"go.uber.org/zap"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/responses"
)

const IdentityCookieKey = "Identity"

type KeyJWTClaims struct{}
type KeyUserId struct{}

type UnauthorizedError struct {
	Err error
}
type InvalidTokenError struct {
	Err error
}

func (err UnauthorizedError) Error() string {
	return "user not signed in"
}

func (err InvalidTokenError) Error() string {
	return "identity token is invalid"
}

type AuthMWHandler struct {
	l   *zap.SugaredLogger
	key string
}

//NewAuthMWHandler creates a new AuthMWR Handler
func NewAuthMWHandler(logger *zap.SugaredLogger, key string) *AuthMWHandler {
	return &AuthMWHandler{
		l:   logger,
		key: key,
	}
}

//AuthMWR is the auth middleware for a required authenticated request.
func (h *AuthMWHandler) AuthMWR(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		claims, err := h.RequireAuthTokenHandler(rw, r)

		if err != nil {
			switch e := err.(type) {
			case *UnauthorizedError:
				responses.JSONError(rw, &responses.ErrorResponse{Error: err.Error()}, http.StatusUnauthorized)
				return
			case *InvalidTokenError:
				h.l.Errorw("invalid token", "error", e.Err)
				responses.JSONError(rw, &responses.ErrorResponse{Error: err.Error()}, http.StatusUnauthorized)
				return
			}
		}

		idCtx := context.WithValue(r.Context(), KeyUserId{}, claims["sub"])
		ctx := context.WithValue(idCtx, KeyJWTClaims{}, claims)

		next.ServeHTTP(rw, r.WithContext(ctx))

	})
}

//AuthMWO is the auth middleware for a optional authenticated request.
func (h *AuthMWHandler) AuthMWO(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		claims, err := h.RequireAuthTokenHandler(rw, r)

		if err != nil {
			switch e := err.(type) {
			case *InvalidTokenError:
				h.l.Errorw("invalid token", "error", e.Err)
				return
			}
		}
		if err == nil {
			//User is signed in so we can pass the read user values.
			idCtx := context.WithValue(r.Context(), KeyUserId{}, claims["sub"])
			ctx := context.WithValue(idCtx, KeyJWTClaims{}, claims)
			next.ServeHTTP(rw, r.WithContext(ctx))
			return
		}

		next.ServeHTTP(rw, r)
	})
}

func (h *AuthMWHandler) RequireAuthTokenHandler(rw http.ResponseWriter, r *http.Request) (jwt.MapClaims, error) {
	cookie, err := r.Cookie("identity")

	if err != nil {
		return nil, &UnauthorizedError{Err: err}
	}

	var claims jwt.MapClaims
	claims, err = ValidateJwt(cookie.Value, h.key)

	if err != nil {
		return nil, &InvalidTokenError{Err: err}
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

//ExtractFromCookie extracts a jwt from the identtiy Cookie
func ExtractFromCookie(r *http.Request) (string, error) {
	cookie, err := r.Cookie(IdentityCookieKey)
	if err != nil {
		return "", &UnauthorizedError{Err: err}
	}

	return cookie.Value, nil
}

func GetKeyGetter(key string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(key), nil
	}
}
