/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package projectCmd

import (
	"errors"
	"fmt"

	"github.com/adam-fraga/ratel/handlers/project"
	er "github.com/adam-fraga/ratel/internal/errors"
	ut "github.com/adam-fraga/ratel/utils"
	"github.com/spf13/cobra"
)

// createProjectCmd represents the createProject command
var createProjectCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new project for your Go templ project with Ratel framework",
	Long: `The create command is a powerful tool designed to streamline the process 
  of setting up a new project for your Go team using the Ratel framework. 
  By simply providing the project name as an argument, 
  this command automatically generates a fully structured project based on a well-defined,
  best-practice architecture.

The generated project includes the following directories and files:

  -LICENSE: A file containing the license information for the project.

  -Makefile: A Makefile for building and managing the project.

  -README.md: A README file providing an overview of the project.

  -config/config.yml: A configuration file for project settings.

  -db: A directory containing database setup and connection logic.

  -docs: A directory for project documentation.

  -errors: A directory for custom error handling.

  -handlers: A directory for HTTP request handlers.

  -middlewares: A directory for middleware logic.

  -package.json: An npm package configuration file.

  -src/css: A directory for Tailwind CSS stylesheets.

  -static: A directory for static assets (e.g., CSS and JS output from TypeScript and Tailwind).

  -test: A directory for project tests.

  -tsconfig.json: A TypeScript configuration file.

  -views: A directory containing templates for the project.

Within the Views directory, the following subdirectories and template files are organized:

  -components: A directory for reusable components.

  -layouts: A directory containing layout templates.

  -metadatas: A directory for metadata templates.

  -pages: A directory for page templates.

  -partials: A directory for partial templates.

  -templates: A directory for general templates.

  The create command simplifies the process of setting up a new project with the Ratel framework. By providing a well-structured architecture, it allows developers to quickly kickstart their development process and focus on building their application.
`,

	Annotations: map[string]string{"category": "project"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			var projectError *er.ProjectError
			if err := project.CreateProject(args[0]); err != nil {
				if errors.As(err, &projectError) {
					fmt.Println("Failed creating the project: ", projectError.Msg)
				}
			}

		} else {
			if err := ut.RunCommand("ratel", true, "project create --help"); err != nil {
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
