/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cacheCmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// cacheCmd represents the cache command
var CacheCmd = &cobra.Command{
	Use:   "cache",
	Short: "Cache commands",
	Long:  `Cache commands provide a way to interact with the cache system. of the web framework.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("CACHE COMMAND CALLED")
	},
}

func addCacheSubCommands() {
	CacheCmd.AddCommand(initCacheCmd)
	CacheCmd.AddCommand(runCacheCmd)
}

func init() {
	addCacheSubCommands()
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cacheCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cacheCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
