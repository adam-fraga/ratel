/*
Copyright © 2024 Admtechlabs adam.fraga@admtechlabs.com
Package dbMigrationCmd provide a way to interact with the migration system of the ratel web framework.
*/

package dbMigrationCmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// migrationCmd represents the migration command
var MigrationCmd = &cobra.Command{
	Use:   "migration",
	Short: "Manage database migrations",
	Long: `The migration command provides tools to manage database schema changes, 
including creating, applying, and rolling back migrations to keep the database in sync with the application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("MIGRATION COMMAND CALLED")
	},
}

func addMigrationSubcommand() {
	MigrationCmd.AddCommand(runMigrationCreateCmd)
	MigrationCmd.AddCommand(runMigrationUpCmd)
	MigrationCmd.AddCommand(runMigrationDownCmd)
	MigrationCmd.AddCommand(runMigrationStatusCmd)
	MigrationCmd.AddCommand(runMigrationVersionCmd)
}

func init() {
	addMigrationSubcommand()

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// migrationCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// migrationCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
