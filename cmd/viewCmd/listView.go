package viewCmd

import (
	h "github.com/adam-fraga/ratel/handlers/views"
	ut "github.com/adam-fraga/ratel/utils"
	"github.com/spf13/cobra"
)

var listViewCmd = &cobra.Command{
	Use:   "list",
	Short: "List all views: pages, components, layouts, partials, templates, meta or all",
	Long: `List all views: pages, components, layouts, partials, templates, meta or all.
  Usage: "ratel view list [pages|components|layouts|partials|templates|meta] or ratel view list for all views"
  `,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			h.ListViews("")
		} else if len(args) == 1 {
			h.ListViews(args[0])
		} else {
			ut.PrintErrorMsg("Invalid number of arguments: ratel view list [pages|components|layouts|partials|templates|meta] or no arguments to list all the views")
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
