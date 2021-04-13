package main

import (
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/config"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/log"
	"github.com/JohnnyS318/RoyalAfgInGo/services/auth/pkg"
	"github.com/JohnnyS318/RoyalAfgInGo/services/auth/pkg/serviceconfig"
)

func main() {
	//Register logger
	logger := log.RegisterService()
	defer log.CleanLogger()

	//configuration
	config.ReadStandardConfig("auth", logger)
	serviceconfig.ConfigureDefaults()

	//start
	pkg.Start(logger)
}
