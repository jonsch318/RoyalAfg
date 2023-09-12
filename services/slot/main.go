package main

import (
	"github.com/jonsch318/royalafg/pkg/config"
	"github.com/jonsch318/royalafg/pkg/log"
	"github.com/jonsch318/royalafg/services/auth/pkg"
	"github.com/jonsch318/royalafg/services/auth/pkg/serviceconfig"
)

func main() {
	logger := log.RegisterService()
	defer log.CleanLogger()

	config.ReadStandardConfig("slot", logger)
	serviceconfig.ConfigureDefaults()

	pkg.Start(logger)
}
