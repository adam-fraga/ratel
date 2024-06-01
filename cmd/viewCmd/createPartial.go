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

var createPartialCmd = &cobra.Command{
	Use:   "create-partial",
	Short: "Create a new view partial with go templ (.templ)",
	Long:  `Create a new view partial with go templ (.templ) in the partials folder.`,
	Run: func(cmd *cobra.Command, args []string) {

		var v h.View
		partial := v.New("partials")

		if len(args) == 0 {
			ut.PrintErrorMsg(fmt.Sprintf("You must provide a name for the %s", partial.Type))
			os.Exit(1)
		} else if len(args) > 0 && len(args) < 100 {
			for i, arg := range args {
				args[i] = s.ToLower(arg)
				args[i] = s.ReplaceAll(args[i], "-", "_")
			}
			if err := v.Create(partial, args); err != nil {
				ut.PrintErrorMsg(err.Error())
			}
		} else {
			ut.PrintErrorMsg(fmt.Sprintf("You cannot create more than 10 %s at once.", partial.Type))
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
