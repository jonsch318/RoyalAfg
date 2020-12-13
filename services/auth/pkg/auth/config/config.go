package config

import (
	"time"

	"github.com/spf13/viper"
)

// Port on which the server listens
const Port = "port"

// GracefulTimeout sets the timout to wait for active connectons after that forcing them to close
const GracefulTimeout = "gracefultimeout"

const UserServiceUrl = "userservice"

const CookieName = "cookiename"
const CookieExpires = "cookieexpires"

const JwtSigningKey = "jwtsigningkey"
const JwtIssuer = "jwtissuer"
const JwtExpiresAt = "jwtexpiresat"

// ConfigureDefaults configures the viper default configuration values
func ConfigureDefaults() {

	viper.SetDefault(UserServiceUrl, "royalafg-user.royalafg.svc.cluster.local:8080")

	// HTTPServer settings
	viper.SetDefault(GracefulTimeout, time.Second*20)
	viper.SetDefault(Port, 8080)

	// Cookie Settings
	viper.SetDefault(CookieName, "identity")
	viper.SetDefault(CookieExpires, time.Now().Add(time.Hour*24*7))

	// Jwt Settings
	viper.SetDefault(JwtIssuer, "github.com/JohnnyS318/RoyalAfgInGo.games")
	viper.SetDefault(JwtExpiresAt, time.Now().Add(time.Hour*24*7).Unix())
}
