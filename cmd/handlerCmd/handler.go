/*
Copyright © 2024 Adm FRG adam.fraga@live.fr
Package handlerCmd provide a way to interact with the handler system of the ratel web framework.
*/

package handlerCmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// cacheCmd represents the cache command
var HandlerCmd = &cobra.Command{
	Use:   "handler",
	Short: "handler commands",
	Long: `Handler commands provide a way to interact with the handler system
  of the web framework. Yo ucan think of handlers as the controllers of the
  web framework.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("HANDLER COMMAND CALLED")
	},
}

func addHandlerSubCommand() {
	HandlerCmd.AddCommand(createHandlerCmd)
	HandlerCmd.AddCommand(listHandlerCmd)
}

func init() {
	addHandlerSubCommand()
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cacheCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cacheCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
