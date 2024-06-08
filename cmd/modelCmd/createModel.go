package modelCmd

import (
	"os"

	s "strings"

	h "github.com/adam-fraga/ratel/handlers/model"
	ut "github.com/adam-fraga/ratel/utils"
	"github.com/spf13/cobra"
)

var createModelCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new model",
	Long: `The "ratel model create: command is an essential part of our web framework toolset. 
  It simplifies creating new handlers by generating files in the handlers directory. 
  You can create up to 20 model files at a time`,
	Run: func(cmd *cobra.Command, args []string) {
		var m h.Model

		if len(args) == 0 {
			ut.PrintErrorMsg("You must provide a name for the model")
			os.Exit(1)
		} else if len(args) > 0 && len(args) < 20 {
			for i, arg := range args {
				args[i] = s.ToLower(arg)
				args[i] = s.ReplaceAll(args[i], "-", "_")
			}
			if err := m.Create(args); err != nil {
				ut.PrintErrorMsg(err.Error())
			}
		} else {
			if err := ut.RunCommandWithOutput("ratel", "model create --help"); err != nil {
				ut.PrintErrorMsg("Error running the command: " + err.Error())
			}
		}
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// middlewareCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// middlewareCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
