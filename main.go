package main

import (
	"github.com/adam-fraga/commands"
)

func main() {
	mainCommands := commands.rootCmd()
	for _, mainCmd := range mainCommands {
		commands.setSubCommands(mainCmd)
	}
	commands.subCmd.Flags().StringP("name", "n", "default", "name of the sub command")
	commands.rootCmd.AddCommand(commands.subCmd)
	commands.rootCmd.AddCommand()
	commands.rootCmd.Execute()
}
