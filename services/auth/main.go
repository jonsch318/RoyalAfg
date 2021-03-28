package main

import (
	gConfig "github.com/JohnnyS318/RoyalAfgInGo/pkg/config"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/log"
	"github.com/JohnnyS318/RoyalAfgInGo/services/auth/config"
	"github.com/JohnnyS318/RoyalAfgInGo/services/auth/pkg"
)

func main() {
	//Register logger
	logger := log.RegisterService()
	defer log.CleanLogger()

	//configuration
	gConfig.ReadStandardConfig("auth", logger)
	config.ConfigureDefaults()

	//start
	pkg.Start(logger)
}
