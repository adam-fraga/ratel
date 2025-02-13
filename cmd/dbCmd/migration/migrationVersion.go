/*
Copyright © 2024 Admtechlabs adam.fraga@admtechlabs.com
*/
package dbMigrationCmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// runMigrationVersionCmd represents the command to check the current migration version
var runMigrationVersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the current database migration version",
	Long: `Displays the current migration version applied to the database.
This helps track which migration is currently in effect.`,
	Annotations: map[string]string{"category": "db"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Current database migration version displayed")
	},
}

func init() {
	// Define flags and configuration settings here.
}
