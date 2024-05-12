/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package middlewareCmd

import (
	"github.com/adam-fraga/ratel/cmd"
	"github.com/adam-fraga/ratel/handlers"
	"github.com/spf13/cobra"
)

// createMiddlewareCmd represents the createMiddleware command
var createMiddlewareCmd = &cobra.Command{
	Use:         "create",
	Short:       "Create a new middleware",
	Long:        "Create and initialize structure for a new middleware with the framework",
	Annotations: map[string]string{"category": "project"},
	Run: func(cmd *cobra.Command, args []string) {
		handlers.CreateGenericMiddleware()
	},
}

func init() {
	cmd.RootCmd.AddCommand(createMiddlewareCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createMiddlewareCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createMiddlewareCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
