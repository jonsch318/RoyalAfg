package config

import (
	"time"

	"github.com/spf13/viper"
)

// Port on which the server listens
const Port = "HttpServer.Port"

// GracefulTimeout sets the timout to wait for active connectons after that forcing them to close
const GracefulTimeout = "HttpServer.GracefulTimeout"

const DatabaseName = "Database.Name"
const DatabaseUrl = "Database.Url"
const DatabaseTimeout = "Database.Timeout"

const CookieName = "HttpServer.Cookie.Name"
const CookieExpires = "HttpServer.Cookie.Expires"

const JwtSigningKey = "HttpServer.Jwt.SigningKey"
const JwtIssuer = "HttpServer.Jwt.Issuer"
const JwtExpiresAt = "HttpServer.Jwt.ExpiresAt"

// ConfigureDefaults configures the viper default configuration values
func ConfigureDefaults() {

	// Database
	viper.SetDefault(DatabaseName, "github.com/JohnnyS318/RoyalAfgInGoInGo")
	viper.SetDefault(DatabaseUrl, "mongodb//localhost:27017/")
	viper.SetDefault(DatabaseTimeout, time.Second*20)

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
