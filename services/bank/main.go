package main

import (
	"flag"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/log"
	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/pkg"
)

func main() {
	logger := log.RegisterService()
	defer log.CleanLogger()

	//ConfigureViper(logger)

	pkg.Start(logger)

}


func ConfigureViper(logger *zap.SugaredLogger) {
	configFile := ""
	flag.StringVar(&configFile, "serviceConfig", "", "serviceConfig file (default is $HOME/.github.com/JohnnyS318/RoyalAfgInGoInGo.d/bank_service.yaml)")

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
		viper.SetConfigFile("/etc/royalafg-bank/serviceConfig.yaml")
	}

	viper.SetEnvPrefix("royalafg")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		logger.Fatalw("Error during serviceConfig file parsing", "error", err)
	}
	logger.Infow("Parsed config file", "file", viper.ConfigFileUsed())

}
