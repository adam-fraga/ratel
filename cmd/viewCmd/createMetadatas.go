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

var createMetadataCmd = &cobra.Command{
	Use:   "create-metadata",
	Short: "Create a new metadatas file with a .templ file in the metadatas folder.",
	Long: `The "ratel view create-metadata" command is an essential part of the toolset provided by our web framework.
It simplifies the process of creating new metadata files by generating new .templ files in the views/metadatas directory.
You can create up to 20 metadatas at a time.`,
	Run: func(cmd *cobra.Command, args []string) {

		var v h.View
		metadata := v.New("metadatas")

		if len(args) == 0 {
			ut.PrintErrorMsg(fmt.Sprintf("You must provide a name for the %s", metadata.Type))
			os.Exit(1)
		} else if len(args) > 0 && len(args) < 20 {
			for i, arg := range args {
				args[i] = s.ToLower(arg)
				args[i] = s.ReplaceAll(args[i], "-", "_")
			}
			if err := v.Create(metadata, args); err != nil {
				ut.PrintErrorMsg(err.Error())
			}
		} else {
			if err := ut.RunCommand("ratel", true, "view create-metadata --help"); err != nil {
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
