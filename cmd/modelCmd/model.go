/*
Copyright © 2024 Adm FRG adam.fraga@live.fr
Package modelCmd provide a way to interact with the model system of the ratel web framework.
*/

package modelCmd

import (
	"fmt"

	// "github.com/adam-fraga/ratel/handlers"
	"github.com/spf13/cobra"
)

// middlewareCmd represents the middleware command
var ModelCmd = &cobra.Command{
	Use:   "model",
	Short: "model management commands for the project",
	Long:  `model management commands for the project.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("MODEL COMMAND CALLED")
	},
}

func addModelSubCommands() {
	ModelCmd.AddCommand(listModelCmd)
	ModelCmd.AddCommand(createModelCmd)
}

func init() {
	addModelSubCommands()
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// middlewareCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// middlewareCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
