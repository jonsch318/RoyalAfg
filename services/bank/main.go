package main

import (
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/config"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/log"
	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/pkg"
	"github.com/JohnnyS318/RoyalAfgInGo/services/user/pkg/serviceconfig"
)

func main() {
	//Register Logger
	logger := log.RegisterService()
	defer log.CleanLogger()

	//Configure
	config.ReadStandardConfig("bank", logger)
	serviceconfig.ConfigureDefaults()

	//Start
	pkg.Start(logger)
}
