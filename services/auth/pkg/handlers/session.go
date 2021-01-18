package handlers

import (
	"net/http"
	"time"

	"github.com/spf13/viper"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/config"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/errors"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/mw"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/responses"
	"github.com/JohnnyS318/RoyalAfgInGo/services/auth/pkg/dto"
	"github.com/JohnnyS318/RoyalAfgInGo/services/auth/pkg/services"
)

// Session verifies the session and extends the jwt token if valid.
// swagger:route GET /api/auth/session session
//
// Session verifies the session and extends the jwt token if valid.
//
// After verification the extended jwt will be passed as a cookie and the user id and username will be returned
//	Consumes:
//
//	Produces:
//	-	application/json
//
//	Schemes: http, https
//
//	Responses:
//	default: ErrorResponse
//	401: ErrorResponse
//	200: UserResponse
func (h *Auth) Session(rw http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(viper.GetString(config.IdentityCookieKey))
	if err != nil || cookie.Value == "" {
		responses.Unauthorized(rw)
		return
	}

	parsed, token, err := services.ExtendToken(cookie.Value)

	if err != nil {
		switch _ := err.(type) {
		case *errors.InvalidTokenError:
			cookie = services.GenerateCookie("", false)
			cookie.Expires = time.Unix(0,0)
			http.SetCookie(rw, cookie)
			responses.Error(rw, "token could not be parsed. This removes the session", http.StatusUnauthorized)
			http.Error(rw,"token could not be parsed", http.StatusUnauthorized)
			return
		default:
			responses.Unauthorized(rw)
			return
		}
	}

	user := mw.FromUserTokenContext(parsed)

	cookie = services.GenerateCookie(token, false)

	http.SetCookie(rw, cookie)
	rw.WriteHeader(http.StatusOK)
	err = ToJSON(dto.UserDTO{
		ID:        user.ID,
		Username:  user.Username,
	}, rw)

	if err != nil {
		h.l.Errorw("json serialization", "error", err)
		responses.Error(rw, "Something went wrong", http.StatusInternalServerError)
		return
	}
}