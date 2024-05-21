/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package viewCmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// createViewCmd represents the createView command
var createViewCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new view with go templ (.templ)",
	Long:  `Create a new view with go templ (.templ) in the views folder.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("createView called")
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
