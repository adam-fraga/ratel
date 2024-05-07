package commands

import (
	"fmt"
	"github.com/adam-fraga/ratel/handlers"
	"github.com/spf13/cobra"
)

// RootCommands returns the root commands of the application
func ProjectCommands() []*cobra.Command {
	return []*cobra.Command{
		{
			Use:   "project",
			Short: "init a new project",
			Long:  "init a new project",
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println("PROJECTS COMMANDS")
			},
		},
	}
}

func ProjectSubCommands() []*cobra.Command {
	return []*cobra.Command{
		{
			Use:         "create",
			Short:       "Create a new project",
			Long:        "Create and initialize structure for a new project with the framework",
			Annotations: map[string]string{"category": "project"},
			Run: func(cmd *cobra.Command, args []string) {
				handlers.InitProject()
			},
		},
		{
			Use:       "list-dependencies",
			Short:     "list all projects dependencies",
			Long:      "list all projects dependencies",
			ValidArgs: []string{"name"},
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println("LIST DEPENDENCIES")
			},
		},
	}

}
