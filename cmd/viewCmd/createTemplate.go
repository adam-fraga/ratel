/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package viewCmd

import (
	"fmt"
	"os"
	s "strings"

	"github.com/spf13/cobra"

	h "github.com/adam-fraga/ratel/handlers/views"
	ut "github.com/adam-fraga/ratel/utils"
)

var createTemplateCmd = &cobra.Command{
	Use:   "create-template",
	Short: "Create a new template file with a .templ file in the templates folder.",
	Long: `The create-template command is an essential part of the toolset provided by our web framework.
It simplifies the process of creating new template files by generating new .templ files in the views/templates directory.
You can create up to 10 templates at a time.`,

	Run: func(cmd *cobra.Command, args []string) {

		var v h.View
		template := v.New("templates")

		if len(args) == 0 {
			ut.PrintErrorMsg(fmt.Sprintf("You must provide a name for the %s", template.Type))
			os.Exit(1)
		} else if len(args) > 0 && len(args) < 10 {
			for i, arg := range args {
				args[i] = s.ToLower(arg)
				args[i] = s.ReplaceAll(args[i], "-", "_")
			}
			if err := v.Create(template, args); err != nil {
				ut.PrintErrorMsg(err.Error())
			}
		} else {
			if err := ut.RunCommandWithOutput("ratel", "view create-template --help"); err != nil {
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
