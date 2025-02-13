package projectCmd

import (
	"errors"
	"github.com/adam-fraga/ratel/handlers/project"
	er "github.com/adam-fraga/ratel/internal/errors"
	ut "github.com/adam-fraga/ratel/utils"
	"github.com/spf13/cobra"
)

var (
	echoFlag     bool
	fiberFlag    bool
	postgresFlag bool
	sqliteFlag   bool
	mongoFlag    bool
)

// initProjectCmd represents the init command
var initProjectCmd = &cobra.Command{
	Use:   "init [repository name]",
	Short: "Initialize a new Go project with optional frameworks using Ratel",
	Long: `
The "init" command helps you initialize a new Go project with the Ratel framework. 
  You need to provide the name of your GitHub repository as an argument with this format:
  
  "github/your-name/repository"

### Optional Flags:

Frameworks:

  You can choose a web framework by providing one of the following flags:
    - **--fiber**  → Initialize the project with the Fiber framework.
    - **--echo**   → Initialize the project with the Echo framework.

Database Providers:

  Select a database provider using one of these flags:
    - **--postgres**  → Set up the project with PostgreSQL.
    - **--mongo**     → Set up the project with MongoDB.
    - **--sqlite**    → Set up the project with SQLite.`,

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
		var dbProvider string

		if echoFlag {
			framework = "Echo"
		} else if fiberFlag {
			framework = "Fiber"
		} else {
			framework = "" // Default go "net/http" package
		}

		if echoFlag {
			dbProvider = "postgres"
		} else if fiberFlag {
			dbProvider = "mongodb"
		} else {
			dbProvider = "sqlite" // Default go "net/http" package
		}

		if err := project.InitProject(repoName, framework, dbProvider); err != nil {
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
	initProjectCmd.Flags().BoolVar(&echoFlag, "postgres", false, "Initialize the project with postgres connector")
	initProjectCmd.Flags().BoolVar(&fiberFlag, "mongo", false, "Initialize the project with mongo connector")
}
