/*
 Copyright © 2024 Adm FRG adam.fraga@live.fr
 Package routerCmd provide a way to interact with the router system of the ratel web framework.
*/

package routerCmd

import (
	"fmt"

	"github.com/adam-fraga/ratel/cmd/routerCmd/routeCmd"
	"github.com/spf13/cobra"
)

var RouterCmd = &cobra.Command{
	Use:   "router",
	Short: "Router commands for the project",
	Long: `Router commands for the project, you can handle the routes of the project here
  you can create several router that contains several routes for your project all of them
  well structured and organized.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ROUTER COMMAND CALLED")
	},
}

func AddrouterSubCmd() {
	RouterCmd.AddCommand(createRouterCmd)
	RouterCmd.AddCommand(listRouterCmd)
	RouterCmd.AddCommand(routeCmd.RouteCmd)
}

func init() {
	AddrouterSubCmd()

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// projectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// projectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
