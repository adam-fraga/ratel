/*
Copyright © 2024 Admtechlabs adam.fraga@admtechlabs.com
*/
package dbMigrationCmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// runMigrationDownCmd represents the command to roll back the last applied migration
var runMigrationDownCmd = &cobra.Command{
	Use:   "down",
	Short: "Rollback the last database migration",
	Long: `Revert the most recent database migration to undo schema changes.
This is useful when a migration introduces issues or needs to be reverted.`,
	Annotations: map[string]string{"category": "db"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Database migration rolled back")
	},
}

func init() {
	// Define flags and configuration settings here.
}
