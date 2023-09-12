package main

import (
	"github.com/jonsch318/royalafg/pkg/config"
	"github.com/jonsch318/royalafg/pkg/log"
	"github.com/jonsch318/royalafg/services/auth/pkg"
	"github.com/jonsch318/royalafg/services/auth/pkg/serviceconfig"
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
