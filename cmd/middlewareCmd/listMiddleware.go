/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package middlewareCmd

import (
	"errors"
	h "github.com/adam-fraga/ratel/handlers/middleware"
	er "github.com/adam-fraga/ratel/internal/errors"
	ut "github.com/adam-fraga/ratel/utils"
	"github.com/spf13/cobra"
)

// listMiddlewareCmd represents the listMiddleware command
var listMiddlewareCmd = &cobra.Command{
	Use:   "list",
	Short: "List all middlewares",
	Long:  "List all middlewares available in the project.",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			if err := h.List(); err != nil {
				var middlewareError *er.MiddlewareError
				if errors.As(err, &middlewareError) {
					ut.PrintErrorMsg("Error listing the middlewares: " + middlewareError.Msg)
				}
			}
		} else {
			if err := ut.RunCommand("ratel", true, "middleware list --help"); err != nil {
				ut.PrintErrorMsg("Error running the command: " + err.Error())
			}
		}
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
