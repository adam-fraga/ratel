/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package dbMigrationCmd

import (
	"github.com/adam-fraga/ratel/cmd"
	"github.com/adam-fraga/ratel/handlers"
	"github.com/spf13/cobra"
)

// runMigrationCmd represents the runMigration command
var runMigrationCmd = &cobra.Command{
	Use:         "run-migration",
	Short:       "Run database migration",
	Long:        `Run database migration to update the database schema.`,
	Annotations: map[string]string{"category": "db"},
	Run: func(cmd *cobra.Command, args []string) {
		handlers.RunDbMigration()
	},
}

func init() {
	cmd.RootCmd.AddCommand(runMigrationCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runMigrationCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runMigrationCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
