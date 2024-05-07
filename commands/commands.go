package commands

import (
	"github.com/spf13/cobra"
)

/*
This function loop on the root commands and add sub commands to them by Looping on root commands and adding
sub commands to them ( you should pass the appropriate sub commands to the function) this function should be
called by looping on a list of sub commands and passing each sub command to it.
*/
func SetSubCommands(rootCmds []*cobra.Command, subCmd *cobra.Command) {
	for _, cmd := range rootCmds {
		switch cmd.Use {
		case "db":
			if subCmd.Annotations["category"] == "db" {
				cmd.AddCommand(subCmd)
			}
		case "project":
			if subCmd.Annotations["category"] == "project" {
				cmd.AddCommand(subCmd)
			}
		case "middleware":
			if subCmd.Annotations["category"] == "middleware" {
				cmd.AddCommand(subCmd)
			}
		case "cache":
			if subCmd.Annotations["category"] == "cache" {
				cmd.AddCommand(subCmd)
			}
		}
	}
}
