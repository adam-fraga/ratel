package moduleCmd

import (
	"os"

	"github.com/adam-fraga/ratel/utils"
	"github.com/spf13/cobra"
)

// ModuleCmd represents the module command
var ModuleCmd = &cobra.Command{
	Use:   "module",
	Short: "Manage backend modules",
	Long: `Manage and generate backend modules for your project. Available modules:

  - auth: User authentication system (login, register, JWT, etc.).
  - oauth: OAuth 2.0 authentication (Google, GitHub, etc.).
  - user: User management with roles and permissions.
  - mailer: Email sending and template management.
  - logger: Structured logging setup.
  - payment: Payment system integration with Stripe.
  - task: Background task processing.
  - storage: File uploads and storage handling.
  - cache: Caching system (Redis, in-memory, etc.).
  - websocket: Real-time communication with WebSockets.
  
Example usage:
  ratel module add --name auth
  ratel module add --name payment`,
}

var addModuleCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new backend module",
	Long:  `Creates a new backend module and populates the necessary template files for it.`,
	Run: func(cmd *cobra.Command, args []string) {
		moduleName, _ := cmd.Flags().GetString("name")
		if moduleName == "" {
			utils.PrintErrorMsg("❌ Error: You must provide a module name using --name flag")
			os.Exit(1)
		}
		utils.PrintCyanInfoMsg("📦 Creating module: " + moduleName)
		// Here you can add logic to generate the module structure
		utils.PrintSuccessMsg("✅ Module successfully created!")
	},
}

func init() {
	addModuleCmd.Flags().String("name", "", "Specify the name of the module")
	ModuleCmd.AddCommand(addModuleCmd)
}
