/*
Copyright © 2024 Adm FRG adam.fraga@live.fr
Package dbCmd provide a way to interact with the database system of the ratel web framework.
*/
package dbCmd

import (
	dbMigrationCmd "github.com/adam-fraga/ratel/cmd/dbCmd/migration"
	ut "github.com/adam-fraga/ratel/utils"
	"github.com/spf13/cobra"
)

// dbCmd represents the db command
var DbCmd = &cobra.Command{
	Use:   "db",
	Short: "Database commands",
	Long: `The "db" command encompasses a suite of subcommands tailored for managing database-related tasks within the project. 
  It serves as the central interface for executing operations related to database management and configuration.
  This command provides a comprehensive set of functionalities for working with databases, 
  including tasks such as database setup, migration, seeding, querying, and administration.
  With the "db" command, you can easily perform various database operations, such as creating database schemas,
  migrating database structures, seeding initial data, executing database queries, and managing database connections.
  By consolidating all database-related functionalities under a single entry point, 
  the "db" command streamlines the workflow and enhances productivity for developers working on web applications 
  that require database integration.`,

	Run: func(cmd *cobra.Command, args []string) {
		ut.RunCommandWithOutput("ratel", "db --help")
	},
}

func addDbSubCommands() {
	DbCmd.AddCommand(createDbContainerCmd)
	DbCmd.AddCommand(initCmd)
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
