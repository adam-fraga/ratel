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
	Use: "ratel",
	Short: `Ratel is a versatile web framework designed to streamline web development with Go, offering a range 
  of powerful tools and functionalities to simplify project setup, management, and deployment.`,
	Long: `Ratel is a versatile web framework designed to streamline web development with Go. Developed by [Your Name], Ratel is an integral part of their day-to-day workflow, providing a robust set of tools and functionalities to simplify project setup, management, and deployment.

Continuously updated with new features and improvements, Ratel empowers developers to build robust and scalable 
web applications with ease. Its intuitive command-line interface and comprehensive toolset make it a 
go-to choice for web development projects.

Key Features:
- Project Management: Ratel offers a range of commands for managing various aspects of your web project, including project creation, initialization, configuration, and deployment.
- Middleware Support: The framework includes middleware commands for managing authentication, logging, error handling, and request/response modification.
- Database Integration: Ratel provides database commands for setup, migration, seeding, querying, and administration.
- View Handling: The framework offers commands for handling views within the project, facilitating creation, listing, updating, and deletion of views and components.

List of Commands:
- project: Project-related commands for setup, management, and deployment.
- middleware: Middleware commands for managing middleware functionalities.
- db: Database commands for setup, migration, seeding, querying, and administration.
- view: View commands for handling views within the project.
- create-dev-container: Command for creating a database container for the project using Docker.
- list: Command for listing views within the project.
- more soon...

Ratel is more than just a framework—it's a powerful ally in your web development journey, 

Developed by Adam Fraga, Ratel is an essential tool in modern web development workflows, 
constantly evolving to meet the needs of developers.
`,
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
