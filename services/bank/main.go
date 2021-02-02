package main

import (
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/config"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/log"
	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/pkg"
	"github.com/JohnnyS318/RoyalAfgInGo/services/user/pkg/serviceconfig"
)

func main() {
	logger := log.RegisterService()
	defer log.CleanLogger()

	config.ReadStandardConfig("bank", logger)
	serviceconfig.ConfigureDefaults()

	pkg.Start(logger)
}
