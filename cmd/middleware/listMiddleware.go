/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package middlewareCmd

import (
	"github.com/adam-fraga/ratel/cmd"
	"github.com/adam-fraga/ratel/handlers"
	"github.com/spf13/cobra"
)

// listMiddlewareCmd represents the listMiddleware command
var listMiddlewareCmd = &cobra.Command{
	Use:   "listMiddleware",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		handlers.ListMiddlewares()
	},
}

func init() {
	cmd.RootCmd.AddCommand(listMiddlewareCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listMiddlewareCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listMiddlewareCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
