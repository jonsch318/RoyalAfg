package main

import (
	"github.com/jonsch318/royalafg/pkg/config"
	"github.com/jonsch318/royalafg/pkg/log"
	"github.com/jonsch318/royalafg/services/poker-matchmaker/pkg"
	"github.com/jonsch318/royalafg/services/poker-matchmaker/pkg/serviceconfig"
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
