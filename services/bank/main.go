package main

import (
	"flag"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/log"
	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/pkg"
)

func main() {
	logger := log.RegisterService()
	defer log.CleanLogger()

	configFile := ""
	flag.StringVar(&configFile, "config", "", "config file (default is $HOME/.github.com/JohnnyS318/RoyalAfgInGoInGo.d/bank_service.yaml)")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	err := viper.BindPFlags(pflag.CommandLine)

	if err != nil {
		logger.Fatalw("Error during Flag binding", "error", err)
	}

	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			logger.Fatalw("Error when finding home directory", "error", err)
		}

		viper.AddConfigPath(home + "/.RoyalAfgInGo.d/")
		viper.AddConfigPath(".")
		viper.SetConfigFile("/etc/royalafg-bank/config.yaml")
	}

	viper.SetEnvPrefix("royalafg")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		logger.Fatalw("Error during config file parsing", "error", err)
	}

	logger.Infow("Parsed config file", "file", viper.ConfigFileUsed())

	pkg.Start(logger)

}
