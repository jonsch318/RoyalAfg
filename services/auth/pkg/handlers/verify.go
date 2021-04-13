package handlers

import (
	"net/http"

	"github.com/form3tech-oss/jwt-go"
	"github.com/spf13/viper"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/config"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/dtos"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/mw"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/utils"
)


// VerifyLoggedIn verifies and validates the cookie and it's jwt token. returns 401 if you are not signed in and 200 if everything is valid-
// swagger:route GET /account/verify authentication account verifyLoggedIn
//
// Verify that the user is logged in
//
// This will return either status code 401 Unauthorized if user is not signed in and 200 when the login token is valid
//
//	Consumes:
//
// 	Produces:
//	-	application/json
//
//	Schemes: http, https
//
//	Security:
//
// 	Responses:
//	default: SessionInfo
//	200: NoContentResponse
func (h *Auth) VerifyLoggedIn(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("X-Content-Type-Options", "nosniff")

	//Get token string from cookie
	tokenRaw, err := mw.ExtractFromCookie(r)
	if err != nil {
		h.l.Infow("Session cookie not found", "error", err)
		_ = utils.ToJSON(&dtos.SessionInfo{
			Authenticated: false,
		}, rw)
		return
	}

	//Verify signature of the token
	token, err := jwt.Parse(tokenRaw, mw.GetKeyGetter(viper.GetString(config.JWTSigningKey)))
	if err != nil {
		h.l.Infow("Session token not signed", "error", err)
		_ = utils.ToJSON(&dtos.SessionInfo{
			Authenticated: false,
		}, rw)
		return
	}

	//Check authentication or authorization with other services. (currently a NoOp)
	authenticated := h.Auth.VerifyAuthentication(mw.FromUserTokenContext(token))
	if !authenticated {
		h.l.Errorw("A error during login verification", "error", err)
		_ = utils.ToJSON(&dtos.SessionInfo{
			Authenticated: false,
		}, rw)
	}

	rw.WriteHeader(http.StatusOK)
	_ = utils.ToJSON(&dtos.SessionInfo{
		Authenticated: true,
	}, rw)
}
