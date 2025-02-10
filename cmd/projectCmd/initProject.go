package projectCmd

import (
	"errors"
	"github.com/adam-fraga/ratel/handlers/project"
	er "github.com/adam-fraga/ratel/internal/errors"
	ut "github.com/adam-fraga/ratel/utils"
	"github.com/spf13/cobra"
)

var (
	echoFlag  bool
	fiberFlag bool
)

// initProjectCmd represents the init command
var initProjectCmd = &cobra.Command{
	Use:   "init [repository name]",
	Short: "Initialize a new Go project with optional frameworks using Ratel",
	Long: `
The "init" command helps you initialize a new Go project with the Ratel framework. 
  You need to provide the name of your GitHub repository as an argument with this format:
  
  "github/your-name/repository"

You can optionally specify a framework for your project:
  - Fiber: Use the --fiber flag to initialize with the Fiber framework.
  - Echo: Use the --echo flag to initialize with the Echo framework.

If no framework flag is provided, the project will be initialized without any framework.
`,

	Annotations: map[string]string{"category": "project"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			if err := ut.RunCommand("ratel", true, "project init --help"); err != nil {
				ut.PrintErrorMsg("Error running the command: " + err.Error())
			}
			return
		}

		repoName := args[0]

		var framework string
		if echoFlag {
			framework = "Echo"
		} else if fiberFlag {
			framework = "Fiber"
		} else {
			framework = "" // Default go "net/http" package
		}

		if err := project.InitProject(repoName, framework); err != nil {
			var projectError *er.ProjectError
			if errors.As(err, &projectError) {
				ut.PrintErrorMsg("Failed to initialize the project " + projectError.Msg)
			}
		}
	},
}

func init() {
	initProjectCmd.Flags().BoolVar(&echoFlag, "echo", false, "Initialize the project with Echo framework")
	initProjectCmd.Flags().BoolVar(&fiberFlag, "fiber", false, "Initialize the project with Fiber framework")
}
