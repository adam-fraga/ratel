/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/adam-fraga/ratel/cmd/cacheCmd"
	"github.com/adam-fraga/ratel/cmd/dbCmd"
	"github.com/adam-fraga/ratel/cmd/handlerCmd"
	"github.com/adam-fraga/ratel/cmd/middlewareCmd"
	"github.com/adam-fraga/ratel/cmd/modelCmd"
	"github.com/adam-fraga/ratel/cmd/projectCmd"
	"github.com/adam-fraga/ratel/cmd/routerCmd"
	"github.com/adam-fraga/ratel/cmd/viewCmd"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ratel",
	Short: "Ratel is a cli tool for managing Anago web framework.",
	Long:  `Ratel is a cli tool for managing Anago web framework.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

func addCommands() {
	rootCmd.AddCommand(projectCmd.ProjectCmd)
	rootCmd.AddCommand(middlewareCmd.MiddlewareCmd)
	rootCmd.AddCommand(viewCmd.ViewCmd)
	rootCmd.AddCommand(dbCmd.DbCmd)
	rootCmd.AddCommand(cacheCmd.CacheCmd)
	rootCmd.AddCommand(routerCmd.RouterCmd)
	rootCmd.AddCommand(modelCmd.ModelCmd)
	rootCmd.AddCommand(handlerCmd.HandlerCmd)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	addCommands()

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ratel.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".ratel" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".ratel")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
