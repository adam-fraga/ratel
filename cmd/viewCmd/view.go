/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package viewCmd

import (
	ut "github.com/adam-fraga/ratel/utils"
	"github.com/spf13/cobra"
)

// viewCmd represents the view command
var ViewCmd = &cobra.Command{
	Use:   "view",
	Short: "View commands for the project",
	Long: `The "view" command provides a set of subcommands to manage views within the project.
It encompasses all the commands related to handling and managing views. This command serves as a central hub 
for performing operations on views, including creating, listing...you can interact with various view components 
such as pages, components, layouts, partials, templates, and metadata and forms.
By encapsulating all view-related commands under a single entry point, the "view" command simplifies view 
management and streamlines the workflow for developers working on frontend components within the project.`,

	Run: func(cmd *cobra.Command, args []string) {
		ut.RunCommand("ratel", true, "view --help")
	},
}

func addViewSubCommands() {
	ViewCmd.AddCommand(createComponentCmd)
	ViewCmd.AddCommand(createLayoutCmd)
	ViewCmd.AddCommand(createMetadataCmd)
	ViewCmd.AddCommand(createPageCmd)
	ViewCmd.AddCommand(createPartialCmd)
	ViewCmd.AddCommand(createTemplateCmd)
	ViewCmd.AddCommand(createFormCmd)
	ViewCmd.AddCommand(listViewCmd)
}

func init() {
	addViewSubCommands()

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// viewCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// viewCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
