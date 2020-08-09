/*
Copyright © 2020 Jonas Schneider jonas.max.schneider@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/account"
	"github.com/spf13/cobra"
)

// accountCmd represents the account command
var accountCmd = &cobra.Command{
	Use:   "account",
	Short: "Start the account service of the RoyalAfg project",
	Long:  `The account service is all about the user. It includes registration, signin, authentication`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting the Application... \n Logging will be initialized in a bit.")
		account.Start()
	},
}

func init() {
	startCmd.AddCommand(accountCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// accountCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// accountCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
