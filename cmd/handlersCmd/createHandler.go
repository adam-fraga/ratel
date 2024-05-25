/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package handlerCmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// createViewCmd represents the createView command
var createHandlerCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new handler for the web framework",
	Long:  `Create a new handler for the web framework. Handlers are the controllers of the web framework.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("CREATE HANDLER CALLED")
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createViewCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createViewCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
