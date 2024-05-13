/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cacheCmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// runCacheCmd represents the runCache command
var runCacheCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a new cache instance",
	Long:  `Run a new cache instance for the web framework.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("runCache called")
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCacheCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCacheCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
