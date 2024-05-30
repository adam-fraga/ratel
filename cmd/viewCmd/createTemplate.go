/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package viewCmd

import (
	"github.com/spf13/cobra"

	h "github.com/adam-fraga/ratel/handlers/views"
	ut "github.com/adam-fraga/ratel/utils"
)

var createTemplateCmd = &cobra.Command{
	Use:   "create-template",
	Short: "Create a new view template with go templ (.templ)",
	Long:  `Create a new view template with go templ (.templ) in the templates folder.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := h.CreateView("templates", args); err != nil {
			ut.PrintErrorMsg(err.Error())
		}
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createViewCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createViewCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
