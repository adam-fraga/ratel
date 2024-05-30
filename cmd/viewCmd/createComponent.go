/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package viewCmd

import (
	"os"
	s "strings"

	"github.com/spf13/cobra"

	h "github.com/adam-fraga/ratel/handlers/views"
	ut "github.com/adam-fraga/ratel/utils"
)

// createViewCmd represents the createView command
var createComponentCmd = &cobra.Command{
	Use:   "create-component",
	Short: "Create a new view component with go templ (.templ)",
	Long:  `Create a new view component with go templ (.templ) in the component folder.`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			ut.PrintErrorMsg("You must provide a name for the component")
			os.Exit(1)
		} else if len(args) > 0 && len(args) < 100 {
			for i, arg := range args {
				args[i] = s.ToLower(arg)
				args[i] = s.ReplaceAll(args[i], "-", "_")
			}
			if err := h.CreateView("components", args); err != nil {
				ut.PrintErrorMsg(err.Error())
			}
		} else {
			ut.PrintErrorMsg("You cannot create more than 100 components at once.")
			os.Exit(1)
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
