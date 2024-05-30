/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package viewCmd

import (
	h "github.com/adam-fraga/ratel/handlers/views"
	ut "github.com/adam-fraga/ratel/utils"
	"github.com/spf13/cobra"
)

// createViewCmd represents the createView command
var createPageCmd = &cobra.Command{
	Use:   "create-page",
	Short: "Create a new view page with go templ (.templ)",
	Long:  `Create a new view page with go templ (.templ) in the views folder.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := h.CreateView("pages", args); err != nil {
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
