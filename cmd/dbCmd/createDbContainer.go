/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package dbCmd

import (
	"fmt"

	"github.com/adam-fraga/ratel/handlers/db"
	"github.com/adam-fraga/ratel/internal/errors"
	ut "github.com/adam-fraga/ratel/utils"
	"github.com/spf13/cobra"
)

// dbCmd represents the db command
var createDbContainerCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a database container",
	Long:  `Create a database container for the project using Docker and PostgreSQL`,
	Run: func(cmd *cobra.Command, args []string) {
		provider, err := cmd.Flags().GetString("provider")
		if err != nil {
			var error = &errors.DevError{
				Type:       "Error",
				Origin:     "createDbContainerCmd",
				FileOrigin: "handlers/dbCmd/createDbContainer.go",
				Msg:        err.Error() + fmt.Sprintf("Error getting the provider flag")}

			ut.PrintErrorMsg(error.Msg)
		}

		db.InitDbDevelopmentContainer(provider)
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
