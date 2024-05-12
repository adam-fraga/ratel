/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package dbCmd

import (
	"fmt"

	"github.com/adam-fraga/ratel/cmd"
	"github.com/spf13/cobra"
)

// runDbContainerCmd represents the runDbContainer command
var runDbContainerCmd = &cobra.Command{
	Use:         "run-container",
	Short:       "Create a new database container with docker",
	Long:        "delete a project",
	Annotations: map[string]string{"category": "db"},
	ValidArgs:   []string{"sqlite, postgres, mysql, mongo"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please provide a database type")
			return
		}
		switch args[0] {
		case "sqlite":
			handlers.RunDbContainer("sqlite")
		case "postgres":
			handlers.RunDbContainer("postgres")
		case "mysql":
			handlers.RunDbContainer("mysql")
		case "mongo":
			handlers.RunDbContainer("mongo")
		default:
			fmt.Println("Please provide a valid database type")
		}
	},
}

func init() {
	RootCmd.AddCommand(runDbContainerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runDbContainerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runDbContainerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
