package commands

import (
	"fmt"
	"github.com/adam-fraga/ratel/handlers"
	"github.com/spf13/cobra"
)

// RootCommands returns the root commands of the application
func MiddlewareCommands() []*cobra.Command {
	return []*cobra.Command{
		{
			Use:   "middleware",
			Short: "middleware commands",
			Long:  "middleware commands",
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println("MIDDLEWARE COMMANDS")
			},
		},
	}
}

func MiddlewareSubCommands() []*cobra.Command {
	return []*cobra.Command{
		{
			Use:         "create",
			Short:       "Create a new middleware",
			Long:        "Create and initialize structure for a new middleware with the framework",
			Annotations: map[string]string{"category": "project"},
			Run: func(cmd *cobra.Command, args []string) {
				handlers.CreateGenericMiddleware()
			},
		},
		{
			Use:       "create-auth",
			Short:     "Create a new auth middleware",
			Long:      "Create and initialize structure for a new auth middleware with the framework",
			ValidArgs: []string{"name"},
			Run: func(cmd *cobra.Command, args []string) {
				if len(args) == 0 {
					fmt.Println("Please provide a name for the middleware")
					return
				}
				handlers.CreateAuthMiddleware(args[0])
			},
		},
	}

}
