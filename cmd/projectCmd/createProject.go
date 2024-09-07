/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package projectCmd

import (
	"github.com/adam-fraga/ratel/handlers/project"
	ut "github.com/adam-fraga/ratel/utils"
	"github.com/spf13/cobra"
)

// createProjectCmd represents the createProject command
var createProjectCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new project for your Go teempl project with Ratel framework",
	Long: `
The "create" command allows you to generate a new project for your Go team project using the Ratel framework.
This command takes the project name as an argument and generates a project structure based on a predefined architecture.

The generated project structure includes the following directories and files:

- LICENSE: A file containing the license information for the project.
- Makefile: A Makefile for building and managing the project.
- README.md: A README file providing information about the project.
- config/config.yml: Configuration file for project settings.
- db/db.go: Database package containing database setup and connection logic.
- docs: Directory for project documentation.
- error/error.go: Error handling package.
- handler/indexHandler.go: Handler package for HTTP request handling.
- middlewares/auth-middleware.go: Middleware package for authentication.
- model/user.go: Model package containing data models for the project.
- package.json: npm package configuration file.
- ratel: Ratel framework directory.
- src/css: CSS stylesheets directory.
- src/script: TypeScript scripts directory.
- static: Static assets directory.
- test/index.test.go: Test file for project testing.
- tsconfig.json: TypeScript configuration file.
- views: Views directory containing templates for the project.

Within the "views" directory, there are subdirectories and template files organized as follows:
- components: Directory for reusable components.
- layouts: Directory containing layout templates.
- metadatas: Directory for metadata templates.
- pages: Directory for page templates.
- partials: Directory for partial templates.
- templates: Directory for general templates.

Users have the flexibility to override npm dependencies as needed to customize the project 
according to their requirements.

This command simplifies the process of creating a new project with the Ratel framework,
providing a structured architecture to kickstart your development process.
`,

	Annotations: map[string]string{"category": "project"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			project.CreateProject(args[0])
		} else {
			if err := ut.RunCommand("ratel", true, "project create --help"); err != nil {
				ut.PrintErrorMsg("Error running the command: " + err.Error())
			}
		}
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
