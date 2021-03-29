package main

import (
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/config"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/log"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker-matchmaker/pkg"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker-matchmaker/pkg/serviceconfig"
)

func main() {

	//Register logger
	logger := log.RegisterService()
	defer log.CleanLogger()

	//Configure
	config.ReadStandardConfig("poker-matchmaker", logger)
	serviceconfig.RegisterDefaults()

	//Start
	pkg.Start(logger)
}

