/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package middlewareCmd

import (
	h "github.com/adam-fraga/ratel/handlers/middleware"
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
				ut.PrintErrorMsg("Error listing the middlewares: " + err.Error())
			}
		} else {
			if err := ut.RunCommandWithOutput("ratel", "middleware list --help"); err != nil {
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
