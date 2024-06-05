/*
Copyright © 2024 Adm FRG adam.fraga@live.fr
Package routerCmd provide a way to interact with the router system of the ratel web framework.
*/

package routerCmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var createRouterCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new router",
	Long:  `Create a new router`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("CREATE ROUTER COMMAND CALLED")
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// middlewareCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// middlewareCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
