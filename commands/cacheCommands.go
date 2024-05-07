package commands

import (
	"fmt"
	"github.com/adam-fraga/ratel/handlers"
	"github.com/spf13/cobra"
)

// RootCommands returns the root commands of the application
func CacheCommands() []*cobra.Command {
	return []*cobra.Command{
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

func CacheSubCommands() []*cobra.Command {
	return []*cobra.Command{
		{
			Use:         "init",
			Short:       "init a new cache",
			Long:        "Create and initialize structure for a new middleware with the framework",
			Annotations: map[string]string{"category": "project"},
			ValidArgs:   []string{"name"},
			Run: func(cmd *cobra.Command, args []string) {
				if len(args) == 0 {
					fmt.Println("Please provide a name for the middleware")
					return
				}
				handlers.InitCache(args[0])
			},
		},
		{
			Use:         "run",
			Short:       "run the cache",
			Long:        "run the cache",
			Annotations: map[string]string{"category": "project"},
			Run: func(cmd *cobra.Command, args []string) {
				handlers.RunCacheContainer()
			},
		},
	}
}
