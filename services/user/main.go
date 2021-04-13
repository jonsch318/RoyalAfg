package main

import (
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/config"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/log"

	"github.com/JohnnyS318/RoyalAfgInGo/services/user/pkg"
	"github.com/JohnnyS318/RoyalAfgInGo/services/user/pkg/serviceconfig"
)

func main() {
	//Register logger
	logger := log.RegisterService()
	defer log.CleanLogger()

	//Configure
	config.ReadStandardConfig("user", logger)
	serviceconfig.ConfigureDefaults()

	//Start
	pkg.Start(logger)
}
