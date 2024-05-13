/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package middlewareCmd

import (
	"fmt"

	"github.com/adam-fraga/ratel/handlers"
	"github.com/spf13/cobra"
)

// createAuthMiddlewareCmd represents the createAuthMiddleware command
var createAuthMiddlewareCmd = &cobra.Command{
	Use:         "create-auth",
	Short:       "Create a new auth middleware",
	Long:        "Create and initialize structure for a new auth middleware with the framework",
	Annotations: map[string]string{"category": "project"},
	ValidArgs:   []string{"name"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please provide a name for the middleware")
			return
		}
		handlers.CreateAuthMiddleware(args[0])
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createAuthMiddlewareCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createAuthMiddlewareCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
