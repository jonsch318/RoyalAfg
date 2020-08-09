/*
Copyright © 2020 Jonas Schneider jonas.max.schnedier@gmail.com
*/
package cmd

import (
	"time"

	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

func init() {
	rootCmd.AddCommand(startCmd)

	startCmd.Flags().Int("port", 8080, "Defines the port on which the royalafg server will listen for request")
	viper.BindPFlag("port", startCmd.Flags().Lookup("port"))

	startCmd.Flags().Duration("gracyfulTimeout", time.Second*20, "The duration for which the server waits for existing connections to finish")
	viper.BindPFlag("GracefulTimeout", startCmd.Flags().Lookup("gracefulTimeout"))
}
