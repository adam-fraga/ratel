/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package projectCmd

import (
	"fmt"

	"github.com/spf13/cobra"
	// "github.com/adam-fraga/ratel/handlers"
)

// listProjectDependenciesCmd represents the listProjectDependencies command
var listProjectDependenciesCmd = &cobra.Command{
	Use:   "list-dependencies",
	Short: "List all javascript and golang project dependencies",
	Long:  `List all project dependencies available in the project.`,
	Run: func(cmd *cobra.Command, args []string) {
		//ADD LOGIC THAT EXTRACT DEPENDENCIES FROM PACKAGE.JSON AND GO.MOD THEN PRINT THEM
		fmt.Println("listProjectDependencies called")
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listProjectDependenciesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listProjectDependenciesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
