package routeCmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteRouteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a route",
	Long:  `Delete a route`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("DELETE ROUTE COMMAND CALLED")
	},
}

func init() {}
