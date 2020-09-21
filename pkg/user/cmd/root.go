package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/fsnotify/fsnotify"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/user/pkg/user"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "user",
	Short: "Start the user service of the RoyalAfg project",
	Long:  `The user service is all about the user. It includes registration, signin, authentication`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting the user service... \n Logging will be initialized in a bit.")
		user.Start()
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

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/RoyalAfgInGo.d/user_service.yaml)")

	rootCmd.Flags().Int("port", 8080, "Defines the port on which the RoyalAfg user service will listen for http api request")
	viper.BindPFlag("Port", rootCmd.Flags().Lookup("port"))

	rootCmd.Flags().Duration("gracyfulTimeout", time.Second*20, "The duration for which the server waits for existing connections to finish")
	viper.BindPFlag("GracefulTimeout", rootCmd.Flags().Lookup("gracefulTimeout"))

	rootCmd.Flags().String("db-url", "", "Url to connect to the mongo database")
	viper.BindPFlag("Database.Url", rootCmd.Flags().Lookup("db-url"))

	rootCmd.Flags().String("user-pepper", "", "Pepper to include after the hash")
	viper.BindPFlag("User.Pepper", rootCmd.Flags().Lookup("user-pepper"))

	rootCmd.Flags().String("jwt-key", "", "Key to sign jwts for authorization")
	viper.BindPFlag("Jwt.SigningKey", rootCmd.Flags().Lookup("jwt-key"))

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
		viper.SetConfigName("user_service")
		viper.SetConfigFile("/etc/royalafg-user/config.yaml")
	}

	viper.SetEnvPrefix("ryluser")
	viper.BindEnv("mongodburl")
	viper.BindEnv("port")
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed %v", e.Name)
		viper.ReadInConfig()
	})
}
