/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cacheCmd

import (
	"fmt"

	"github.com/adam-fraga/ratel/cmd"
	"github.com/spf13/cobra"
)

// cacheCmd represents the cache command
var cacheCmd = &cobra.Command{
	Use:   "cache",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Cache command called")
	},
}

func addCacheSubCommands() {
	cacheCmd.AddCommand(initCacheCmd)
	cacheCmd.AddCommand(runCacheCmd)
}

func init() {
	cmd.RootCmd.AddCommand(cacheCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cacheCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cacheCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
