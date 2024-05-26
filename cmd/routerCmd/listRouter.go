package routerCmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listRouterCmd = &cobra.Command{
	Use:   "list",
	Short: "List all routers",
	Long:  `List all routers created in the project.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("LIST ROUTER COMMAND CALLED")
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