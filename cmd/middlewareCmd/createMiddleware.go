/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package middlewareCmd

import (
	"os"
	s "strings"

	h "github.com/adam-fraga/ratel/handlers/middleware"
	ut "github.com/adam-fraga/ratel/utils"
	"github.com/spf13/cobra"
)

// createMiddlewareCmd represents the createMiddleware command
var createMiddlewareCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new middleware",
	Long:  `Create a new middleware in the middlewares folder.`,
	Run: func(cmd *cobra.Command, args []string) {

		var m h.Middleware

		if len(args) == 0 {
			ut.PrintErrorMsg("You must provide a name for the middleware")
			os.Exit(1)
		} else if len(args) > 0 && len(args) < 50 {
			for i, arg := range args {
				args[i] = s.ToLower(arg)
				args[i] = s.ReplaceAll(args[i], "-", "_")
			}
			if err := m.Create(args); err != nil {
				ut.PrintErrorMsg(err.Error())
			}
		} else {
			ut.PrintErrorMsg("You cannot create more than 50 middleware at once.")
			os.Exit(1)
		}
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createMiddlewareCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createMiddlewareCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
