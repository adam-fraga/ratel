/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package viewCmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// dbCmd represents the db command
var deleteViewCmd = &cobra.Command{
	Use:   "delere",
	Short: "Delete a view",
	Long:  `Delete a view from the project.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("DELETE VIEW COMMAND CALLED")
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dbCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dbCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
