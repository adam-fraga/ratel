/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package middlewareCmd

import (
	"github.com/adam-fraga/ratel/handlers"
	"github.com/spf13/cobra"
)

// listMiddlewareCmd represents the listMiddleware command
var listMiddlewareCmd = &cobra.Command{
	Use:   "list",
	Short: "List all middlewares",
	Long:  "List all middlewares available in the project.",

	Run: func(cmd *cobra.Command, args []string) {
		handlers.ListMiddlewares()
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listMiddlewareCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listMiddlewareCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
