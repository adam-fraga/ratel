package commands

import (
	"fmt"
	"github.com/adam-fraga/ratel/handlers"
	"github.com/spf13/cobra"
)

func DbCommands() []*cobra.Command {
	return []*cobra.Command{
		{
			Use:   "db",
			Short: "project commands",
			Long:  "project commands",
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println("DB COMMANDS")
			},
		},
	}
}

func dbSubCommands() []*cobra.Command {
	return []*cobra.Command{
		{
			Use:         "container",
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
		},
	}
}
