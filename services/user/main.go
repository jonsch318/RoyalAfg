package main

import (
	"github.com/jonsch318/royalafg/pkg/config"
	"github.com/jonsch318/royalafg/pkg/log"

	"github.com/jonsch318/royalafg/services/user/pkg"
	"github.com/jonsch318/royalafg/services/user/pkg/serviceconfig"
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
