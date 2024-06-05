/*
Copyright © 2024 Adm FRG adam.fraga@live.fr
Package routeCmd provide a way to interact with the route system of the ratel web framework.
*/

package routeCmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var RouteCmd = &cobra.Command{
	Use:   "route",
	Short: "Route commands",
	Long:  `Route commands provide a way to interact with the route system of the web framework.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ROUTE COMMAND CALLED")
	},
}

func addrouteSubCommands() {
	RouteCmd.AddCommand(createRouteCmd)
	RouteCmd.AddCommand(listRouteCmd)
}

func init() {
	addrouteSubCommands()

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// migrationCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// migrationCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
