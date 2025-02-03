/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package viewCmd

import (
	"fmt"
	"os"
	s "strings"

	"github.com/spf13/cobra"

	h "github.com/adam-fraga/ratel/handlers/view"
	ut "github.com/adam-fraga/ratel/utils"
)

// createViewCmd represents the createView command
var createComponentCmd = &cobra.Command{
	Use:   "create-component",
	Short: "Create a new view component with go templ (.templ) in the component folder.",
	Long: `The "ratel view create component" command is an essential part of the toolset provided by our web framework.
It simplifies the process of creating new view components by generating new .templ files in the views/components directory.
You can create up to 20 components at a time.`,
	Run: func(cmd *cobra.Command, args []string) {

		var v h.View
		component := v.New("components")

		if len(args) == 0 {
			ut.PrintErrorMsg(fmt.Sprintf(" You must provide a name for the %s", component.Type))
			os.Exit(1)
		} else if len(args) > 0 && len(args) < 20 {
			for i, arg := range args {
				args[i] = s.ToLower(arg)
				args[i] = s.ReplaceAll(args[i], "-", "_")
			}
			if err := v.Create(component, args); err != nil {
				ut.PrintErrorMsg(err.Error())
			}
		} else {
			if err := ut.RunCommand("ratel", true, "view create-component --help"); err != nil {
				ut.PrintErrorMsg("Error running the command: " + err.Error())
			}
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
