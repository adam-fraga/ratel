/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package viewCmd

import (
	"errors"
	"fmt"
	"os"
	s "strings"

	h "github.com/adam-fraga/ratel/handlers/view"
	er "github.com/adam-fraga/ratel/internal/errors"
	ut "github.com/adam-fraga/ratel/utils"
	"github.com/spf13/cobra"
)

// createViewCmd represents the createView command
var createPageCmd = &cobra.Command{
	Use:   "create-page",
	Short: "Create a new view page with a .templ file in the pages folder.",
	Long: `The "ratel view create-page" command is an essential part of the toolset provided by our web framework.
It simplifies the process of creating new pages by generating new .templ files in the views/pages directory.
You can create up to 10 pages at a time.`,

	Run: func(cmd *cobra.Command, args []string) {

		var v h.View
		page := v.New("pages")

		if len(args) == 0 {
			ut.PrintErrorMsg(fmt.Sprintf("You must provide a name for the %s", page.Type))
			os.Exit(1)
		} else if len(args) > 0 && len(args) < 10 {
			for i, arg := range args {
				args[i] = s.ToLower(arg)
				args[i] = s.ReplaceAll(args[i], "-", "_")
			}
			if err := v.Create(page, args); err != nil {
				var viewError *er.ViewError
				if errors.As(err, &viewError) {
					ut.PrintErrorMsg("Failed creating page component, error: " + viewError.Msg)
					return
				}
			}
		} else {
			if err := ut.RunCommand("ratel", true, "view create-page --help"); err != nil {
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
