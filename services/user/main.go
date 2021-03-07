package main

import (
	gConfig "github.com/JohnnyS318/RoyalAfgInGo/pkg/config"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/log"

	"github.com/JohnnyS318/RoyalAfgInGo/services/user/pkg"
	"github.com/JohnnyS318/RoyalAfgInGo/services/user/pkg/serviceconfig"
)

func main() {
	logger := log.RegisterService()
	defer log.CleanLogger()

	gConfig.ReadStandardConfig("user", logger)
	serviceconfig.ConfigureDefaults()

	pkg.Start(logger)
}
