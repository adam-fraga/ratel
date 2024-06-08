/*
Copyright © 2024 Adm FRG adam.fraga@live.fr
Package middlewareCmd provide a way to interact with the middleware system of the ratel web framework.
*/

package middlewareCmd

import (
	ut "github.com/adam-fraga/ratel/utils"
	"github.com/spf13/cobra"
)

// middlewareCmd represents the middleware command
var MiddlewareCmd = &cobra.Command{
	Use:   "middleware",
	Short: "Middleware commands for the project",
	Long: `The "middleware" command provides a set of subcommands designed for managing middleware within the project. It serves as the central hub for executing operations related to middleware management and configuration.
  This command offers a range of functionalities for working with middleware, including tasks such as creating new middleware,
  and isting existing middleware.
  With the "middleware" command, you can easily manage middleware functionalities to enhance the project's request 
  handling capabilities, such as authentication, logging, error handling, and request/response modification.
  By encapsulating all middleware-related functionalities under a single entry point, the "middleware" command 
  simplifies the workflow and enhances productivity for developers working on web applications using the framework.`,

	Run: func(cmd *cobra.Command, args []string) {
		ut.RunCommandWithOutput("ratel", "middleware --help")
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
