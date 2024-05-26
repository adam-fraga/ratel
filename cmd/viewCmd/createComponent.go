/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package viewCmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// createViewCmd represents the createView command
var createComponentCmd = &cobra.Command{
	Use:   "create component",
	Short: "Create a new view component with go templ (.templ)",
	Long:  `Create a new view component with go templ (.templ) in the component folder.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("CREATE VIEW COMPONENT CALLED")
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
