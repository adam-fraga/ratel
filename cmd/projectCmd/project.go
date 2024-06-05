/*
Copyright © 2024 Adm FRG adam.fraga@live.fr
Package projectCmd provide a way to interact with the project system of the ratel web framework.
*/

package projectCmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// projectCmd represents the project command
var ProjectCmd = &cobra.Command{
	Use:   "project",
	Short: "Project related commands for the web framework",
	Long:  `Project commands provide a way to interact with the project system of the web framework.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("PROJECT COMMAND CALLED")
	},
}

func AddProjectSubCommands() {
	ProjectCmd.AddCommand(createProjectCmd)
	ProjectCmd.AddCommand(listProjectDependenciesCmd)
}

func init() {
	AddProjectSubCommands()

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// projectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// projectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
