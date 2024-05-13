/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package dbCmd

import (
	"fmt"

	dbMigrationCmd "github.com/adam-fraga/ratel/cmd/dbCmd/migration"
	"github.com/spf13/cobra"
)

// dbCmd represents the db command
var DbCmd = &cobra.Command{
	Use:   "db",
	Short: "Database commands",
	Long:  `Database commands provide a way to interact with the database system of the web framework.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("db called")
	},
}

func addDbSubCommands() {
	// DbCmd.AddCommand(createDbCmd)
	// DbCmd.AddCommand(listDbCmd)
	DbCmd.AddCommand(dbMigrationCmd.MigrationCmd)
}

func init() {
	addDbSubCommands()
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dbCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dbCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
