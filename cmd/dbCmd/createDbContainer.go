/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package dbCmd

import (
	"errors"

	"github.com/adam-fraga/ratel/handlers/db"
	er "github.com/adam-fraga/ratel/internal/errors"
	ut "github.com/adam-fraga/ratel/utils"
	"github.com/spf13/cobra"
)

// dbCmd represents the db command
var createDbContainerCmd = &cobra.Command{
	Use:   "create-dev-database",
	Short: "Create a database container for the project using Docker and your choice of database provider",
	Long: `The "create-dev-database" command helps you set up a database container for your project using Docker. You have the flexibility to choose from various database providers, including MongoDB, Redis, PostgreSQL, and SQLite (for testing purposes, stored as a local file).

To specify the database provider, use one of the following flags:
  --mongo: Create a MongoDB database container.
  --redis: Create a Redis database container.
  --postgres: Create a PostgreSQL database container.
  --sqlite: Use SQLite for testing purposes (no Docker container required).

For MongoDB, Redis, and PostgreSQL, the command orchestrates the creation of a Docker container, 
ensuring seamless integration into your development environment. 
You'll need to fill in the appropriate values in the .env file to configure the connection details.

If you opt for SQLite, the command sets up a local file-based database. 
No Docker container is spun up for SQLite, as it operates directly on the local file system.

This command streamlines the process of provisioning a database container tailored to your project's needs,
whether for local development or testing.`,

	Run: func(cmd *cobra.Command, args []string) {
		provider, err := cmd.Flags().GetString("provider")
		if err != nil {
			ut.PrintErrorMsg(err.Error())
		}
		if provider != "" {
			if err := db.InitDbDevelopmentContainer(provider); err != nil {
				var dbError *er.DBError
				if errors.As(err, &dbError) {
					ut.PrintErrorMsg("Failed initializing DB container, error " + dbError.Msg)
				}

			}
		} else {
			if err := ut.RunCommand("ratel", true, "db create-dev-database --help"); err != nil {
				ut.PrintErrorMsg("Error running the command: " + err.Error())
			}
		}
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	createDbContainerCmd.Flags().StringP("provider", "p", "", "Choose the database provider to create the container (postgres, mariadb, mongo)")
	createDbContainerCmd.MarkFlagRequired("provider")
}
