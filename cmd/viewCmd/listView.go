package viewCmd

import (
	h "github.com/adam-fraga/ratel/handlers/view"
	ut "github.com/adam-fraga/ratel/utils"
	"github.com/spf13/cobra"
)

var listViewCmd = &cobra.Command{
	Use:   "list",
	Short: "List all views: pages, components, layouts, partials, templates, meta or all",
	Long: `The "ratel view list" command allows you to view all views within your project, including:

- No value: all views
- pages
- components
- layouts
- partials
- templates
- metadatas
- forms

If you enter "view list", it will display all views located in the default directories defined for each view type.
These directories are predefined within the project structure and include pages, components, layouts, partials,
templates, and metadata. If you specify a particular view type, for example, "view list components",
it will list all views within the components directory. 
Similarly, you can specify other view types such as pages, layouts, partials, templates, or metadata to view 
the contents of those specific directories. 
This command provides a convenient way to explore the views within your project,
helping you to understand the organization and structure of your project's frontend components.
`,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			h.ListViews("")
		} else if len(args) == 1 {
			h.ListViews(args[0])
		} else {
			ut.RunCommand("ratel", true, "view list --help")
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
