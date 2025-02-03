/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package middlewareCmd

import (
	"errors"
	"os"
	s "strings"

	h "github.com/adam-fraga/ratel/handlers/middleware"
	er "github.com/adam-fraga/ratel/internal/errors"
	ut "github.com/adam-fraga/ratel/utils"
	"github.com/spf13/cobra"
)

// createMiddlewareCmd represents the createMiddleware command
var createMiddlewareCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new middleware file in the middleware folder.",
	Long: `The "ratel middleware create" command is an essential part of the toolset provided by our web framework.
It simplifies the process of creating new middleware by generating new files in the middleware directory.
You can create up to 20 middleware components at a time.`,

	Run: func(cmd *cobra.Command, args []string) {

		var m h.Middleware

		if len(args) == 0 {
			ut.PrintErrorMsg("You must provide a name for the middleware")
			os.Exit(1)
		} else if len(args) > 0 && len(args) < 10 {
			for i, arg := range args {
				args[i] = s.ToLower(arg)
				args[i] = s.ReplaceAll(args[i], "-", "_")
			}
			if err := m.Create(args); err != nil {
				var middlewareError *er.MiddlewareError
				if errors.As(err, &middlewareError) {
					ut.PrintErrorMsg("Failed creating Middleware, error " + middlewareError.Msg)
				}
			}
		} else {
			if err := ut.RunCommand("ratel", true, "middleware create --help"); err != nil {
				ut.PrintErrorMsg("Error running the command: " + err.Error())
			}
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
