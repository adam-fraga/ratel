package routeCmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var createRouterCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new route",
	Long:  `Create a new route in the project, note that route should be associated with a router`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("CREATE ROUTE COMMAND CALLED")
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
