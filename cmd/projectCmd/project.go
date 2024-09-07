/*
Copyright © 2024 Adm FRG adam.fraga@live.fr
Package projectCmd provide a way to interact with the project system of the ratel web framework.
*/

package projectCmd

import (
	ut "github.com/adam-fraga/ratel/utils"
	"github.com/spf13/cobra"
)

// projectCmd represents the project command
var ProjectCmd = &cobra.Command{
	Use:   "project",
	Short: "Project related commands for the web framework",
	Long: `
The "project" command encompasses a collection of subcommands tailored for managing various aspects of the web framework 
project. 
It serves as the central hub for executing operations related to project management and configuration.
This command consolidates all project-related functionalities, including tasks such as project creation, 
initialization, configuration, customization, and deployment. It provides a comprehensive set of tools to streamline the development and maintenance of web applications built on the framework.
With the "project" command, you can easily create new projects, initialize project structures, 
configure project settings, customize project components, and deploy projects to different environments. 
It serves as a versatile toolset for developers to efficiently manage and maintain their web framework projects.
By offering a centralized entry point for project-related commands, the "project" command simplifies 
the workflow and enhances productivity for developers working on web applications using the framework.
`,

	Run: func(cmd *cobra.Command, args []string) {
		ut.RunCommand("ratel", true, "project --help")
	},
}

func AddProjectSubCommands() {
	ProjectCmd.AddCommand(createProjectCmd)
  ProjectCmd.AddCommand(initProjectCmd)
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
