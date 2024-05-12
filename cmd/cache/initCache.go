/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cacheCmd

import (
	"github.com/adam-fraga/ratel/cmd"
	"github.com/adam-fraga/ratel/handlers"
	"github.com/spf13/cobra"
)

// initCacheCmd represents the initCache command
var initCacheCmd = &cobra.Command{
	Use:         "init",
	Short:       "init a new cache",
	Long:        "Create and initialize structure for a new middleware with the framework",
	Annotations: map[string]string{"category": "cache"},
	ValidArgs:   []string{"name"},
	Run: func(cmd *cobra.Command, args []string) {
		handlers.InitCache(args[0])
	},
}

func init() {
	cmd.RootCmd.AddCommand(initCacheCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCacheCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCacheCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
