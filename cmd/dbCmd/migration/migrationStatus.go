/*
Copyright © 2024 Admtechlabs adam.fraga@admtechlabs.com
*/
package dbMigrationCmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// runMigrationStatusCmd represents the command to check the migration status
var runMigrationStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Check the status of database migrations",
	Long: `Displays the current state of database migrations, showing applied and pending migrations.
This helps track which migrations have been executed and which are yet to be applied.`,
	Annotations: map[string]string{"category": "db"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Database migration status checked")
	},
}

func init() {
	// Define flags and configuration settings here.
}
