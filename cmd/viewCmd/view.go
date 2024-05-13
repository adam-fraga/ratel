/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package viewCmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// viewCmd represents the view command
var ViewCmd = &cobra.Command{
	Use:   "view",
	Short: "View commands for the project",
	Long:  `View commands for the project.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("VIEW COMMAND CALLED")
	},
}

func addViewSubCommands() {
	ViewCmd.AddCommand(createViewCmd)
	// ViewCmd.AddCommand(listViewCmd)
}

func init() {
	addViewSubCommands()

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// viewCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// viewCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
