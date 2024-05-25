package routeCmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listRouteCmd = &cobra.Command{
	Use:   "list",
	Short: "List all routes",
	Long:  `List all routes created in the project.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("LIST ROUTE COMMAND CALLED")
	},
}

func init() {}
