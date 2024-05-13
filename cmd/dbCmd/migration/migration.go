/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package dbMigrationCmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// migrationCmd represents the migration command
var MigrationCmd = &cobra.Command{
	Use:   "migration",
	Short: "Migration commands",
	Long:  `Migration commands provide a way to interact with the migration system of the web framework.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("MIGRATION COMMAND CALLED")
	},
}

func addMigrationSubcommand() {
	MigrationCmd.AddCommand(runMigrationCmd)
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
