package main

import (
	"github.com/jonsch318/royalafg/pkg/config"
	"github.com/jonsch318/royalafg/pkg/log"
	"github.com/jonsch318/royalafg/services/bank/pkg"
	"github.com/jonsch318/royalafg/services/user/pkg/serviceconfig"
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
