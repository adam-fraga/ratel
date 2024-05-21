/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package projectCmd

import (
	"fmt"
	"github.com/adam-fraga/ratel/handlers/project"
	"github.com/spf13/cobra"
)

// createProjectCmd represents the createProject command
var createProjectCmd = &cobra.Command{
	Use:         "create",
	Short:       "Create a new project",
	Long:        "Create and initialize structure for a new project with the framework",
	Annotations: map[string]string{"category": "project"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("CREATE PROJECT")
		project.InitProject("Osef")
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createProjectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createProjectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
