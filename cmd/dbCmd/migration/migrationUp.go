/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package dbMigrationCmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// runMigrationCmd represents the runMigration command
var runMigrationUpCmd = &cobra.Command{
	Use:         "up [version]",
	Short:       "Run database migration",
	Long:        `Run database migration to update the database schema.`,
	Annotations: map[string]string{"category": "db"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("runMigration called")
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runMigrationCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runMigrationCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
