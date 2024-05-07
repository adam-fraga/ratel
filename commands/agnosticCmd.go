package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

/*
This function loop on the root commands and add sub commands to them by Looping on root commands and adding
sub commands to them ( you should pass the appropriate sub commands to the function) this function should be
called by looping on a list of sub commands and passing each sub command to it.
*/
func setSubCommands(rootCmds []*cobra.Command, subCmd *cobra.Command) {
	for _, cmd := range rootCmds {
		switch cmd.Use {
		case "db":
			if subCmd.Annotations["category"] == "db" {
				cmd.AddCommand(subCmd)
			}
		case "project":
			if subCmd.Annotations["category"] == "project" {
				cmd.AddCommand(subCmd)
			}
		case "middleware":
			if subCmd.Annotations["category"] == "middleware" {
				cmd.AddCommand(subCmd)
			}
		case "cache":
			if subCmd.Annotations["category"] == "cache" {
				cmd.AddCommand(subCmd)
			}
		}
	}
}

// RootCommands returns the root commands of the application
func RootCommands() []*cobra.Command {
	return []*cobra.Command{
		{
			Use:   "project",
			Short: "init a new project",
			Long:  "init a new project",
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println("PROJECTS COMMANDS")
			},
		},
		{
			Use:   "db",
			Short: "project commands",
			Long:  "project commands",
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println("DB COMMANDS")
			},
		},
		{
			Use:   "middleware",
			Short: "middleware commands",
			Long:  "middleware commands",
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println("MIDDLEWARE COMMANDS")
			},
		},
		{
			Use:   "cache",
			Short: "cache commands",
			Long:  "cache commands",
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println("CACHE COMMANDS")
			},
		},
	}
}

func subCommands() []*cobra.Command {
	return []*cobra.Command{
		{
			Use:         "create",
			Short:       "Create a new project",
			Long:        "Create and initialize structure for a new project with the framework",
			Annotations: map[string]string{"category": "project"},
			Run: func(cmd *cobra.Command, args []string) {
				initProject()
			},
		},
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
					RunDbContainer("sqlite")
				case "postgres":
					RunDbContainer("postgres")
				case "mysql":
					RunDbContainer("mysql")
				case "mongo":
					RunDbContainer("mongo")
				default:
					fmt.Println("Please provide a valid database type")
				}
			},
		},
		{
			Use:       "update",
			Short:     "update a project",
			Long:      "update a project",
			ValidArgs: []string{"name"},
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println("UPDATE COMMAND")
			},
		},
		{
			Use:       "list",
			Short:     "list all projects",
			Long:      "list all projects",
			ValidArgs: []string{"name"},
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println("LIST COMMAND")
			},
		},
	}

}
