package config

import (
	"flag"
	"fmt"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func ReadStandardConfig(serviceName string, logger *zap.SugaredLogger) {
	configFile := ""
	flag.StringVar(&configFile, "config", "", fmt.Sprintf("config file (default is /etc/royalafg-%v/config.yaml", serviceName))

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
		viper.SetConfigFile(fmt.Sprintf("/etc/royalafg-%v/config.yaml", serviceName))
	}

	viper.SetEnvPrefix("royalafg")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		logger.Fatalw("Error during config file parsing", "error", err)
	}

	logger.Infow("Parsed config file", "file", viper.ConfigFileUsed())

	RegisterDefaults()

}
