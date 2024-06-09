/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package dbCmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// dbCmd represents the db command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Run the database schema initialization command",
	Long:  `Run the database schema initialization command to create the tables in the database.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")
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
