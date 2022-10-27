package main

import (
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/config"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/log"
	"github.com/JohnnyS318/RoyalAfgInGo/services/auth/pkg"
	"github.com/JohnnyS318/RoyalAfgInGo/services/auth/pkg/serviceconfig"
)

func main() {
	logger := log.RegisterService()
	defer log.CleanLogger()

	config.ReadStandardConfig("slot", logger)
	serviceconfig.ConfigureDefaults()

	pkg.Start(logger)
}
