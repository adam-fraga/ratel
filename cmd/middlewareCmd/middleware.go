/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package middlewareCmd

import (
	"fmt"

	// "github.com/adam-fraga/ratel/handlers"
	"github.com/spf13/cobra"
)

// middlewareCmd represents the middleware command
var MiddlewareCmd = &cobra.Command{
	Use:   "middleware",
	Short: "Middleware commands for the project",
	Long:  `Middleware commands for the project.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("MIDDLEWARE COMMAND CALLED")
	},
}

func addMiddlewareSubCommands() {
	MiddlewareCmd.AddCommand(listMiddlewareCmd)
	MiddlewareCmd.AddCommand(createMiddlewareCmd)
}

func init() {
	addMiddlewareSubCommands()

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// middlewareCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// middlewareCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
