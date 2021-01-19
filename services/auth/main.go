package main

import (
	gConfig "github.com/JohnnyS318/RoyalAfgInGo/pkg/config"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/log"
	"github.com/JohnnyS318/RoyalAfgInGo/services/auth/config"
	"github.com/JohnnyS318/RoyalAfgInGo/services/auth/pkg"
)

func main() {

	logger := log.RegisterService()
	defer log.CleanLogger()

	gConfig.ReadStandardConfig("auth", logger)
	config.ConfigureDefaults()

	pkg.Start(logger)
}
