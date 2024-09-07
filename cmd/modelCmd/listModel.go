package modelCmd

import (
	h "github.com/adam-fraga/ratel/handlers/model"
	ut "github.com/adam-fraga/ratel/utils"
	"github.com/spf13/cobra"
)

var listModelCmd = &cobra.Command{
	Use:   "list",
	Short: "List all models",
	Long:  `List all models created in the project.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			if err := h.List(); err != nil {
				ut.PrintErrorMsg("Error listing the models: " + err.Error())
			}
		} else {
			if err := ut.RunCommand("ratel", true, "model list --help"); err != nil {
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
