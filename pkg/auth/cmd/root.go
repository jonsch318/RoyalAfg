package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/auth/pkg/auth"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "auth",
	Short: "Start the auth service of the github.com/JohnnyS318/RoyalAfgInGo project",
	Long:  `The auth service is all about the user. It includes registration, signin, authentication`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting the Application... \n Logging will be initialized in a bit.")
		auth.Start()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.github.com/JohnnyS318/RoyalAfgInGoInGo.d/auth_service.yaml)")

	rootCmd.Flags().Int("port", 8080, "Defines the port on which the github.com/JohnnyS318/RoyalAfgInGo server will listen for request")
	//viper.BindPFlag("Port", rootCmd.Flags().Lookup("port"))

	rootCmd.Flags().Duration("gracefulTimeout", time.Second*15, "The duration for which the server waits for existing connections to finish")
	viper.BindPFlag("GracefulTimeout", rootCmd.Flags().Lookup("gracefulTimeout"))

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name "RoyalAfgInGo.d" (without extension).
		viper.AddConfigPath(home + "/.RoyalAfgInGo.d/")
		viper.AddConfigPath(".")
		viper.AddConfigPath("./pkg/auth/")
		viper.AddConfigPath("./.RoyalAfgInGo.d/")
		viper.SetConfigName("auth_service")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
