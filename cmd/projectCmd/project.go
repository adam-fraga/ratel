/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package projectCmd

import (
	"fmt"

	"github.com/adam-fraga/ratel/cmd"
	"github.com/spf13/cobra"
)

// projectCmd represents the project command
var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Project related commands for the web framework",
	Long:  `Project commands provide a way to interact with the project system of the web framework.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("PROJECT COMMAND CALLED")
	},
}

func AddProjectCommand(command *cobra.Command) {
	projectCmd.AddCommand(createProjectCmd)
	projectCmd.AddCommand(listProjectDependenciesCmd)
}

func init() {
	cmd.RootCmd.AddCommand(projectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// projectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// projectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
