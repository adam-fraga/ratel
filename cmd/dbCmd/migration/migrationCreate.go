/*
Copyright © 2024 Admtechlabs adam.fraga@admtechlabs.com
*/

package dbMigrationCmd

import (
	"fmt"
	"regexp"
	// "strconv"
	"strings"

	"github.com/adam-fraga/ratel/utils"
	"github.com/spf13/cobra"
)

// runMigrationCreateCmd represents the create subcommand for database migration
var runMigrationCreateCmd = &cobra.Command{
	Use:   "create [name]",
	Short: "Create and apply a new database migration",
	Long: `Generate a new database migration file and apply it to update the schema.
This ensures that the database structure is in sync with the application models.`,
	Annotations: map[string]string{"category": "db"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			utils.PrintErrorMsg(fmt.Errorf("Failed to create migration file, commands except one arg").Error())
			return
		}
		migrationName := args[0]
		fmt.Println(migrationName)
		migrationNameValidator, err := regexp.MatchString(`^([1-9][0-9]{0,2})_([a-zA-Z0-9]+)$`, migrationName)
		if err != nil {
			utils.PrintErrorMsg("Failed to verify migration name, error: " + err.Error())
			return
		}
		fmt.Println(migrationNameValidator)
		if !migrationNameValidator {
			utils.PrintErrorMsg(fmt.Errorf("Failed to create migration file, migration name should start by a number lower than 1000 use \"_\" as a separator and only alphanul character for your migration name").Error())
			return
		}

		fmt.Println(migrationName)
		migrationInfos := strings.Split(migrationName, "_")
		fmt.Println(migrationInfos)

		// if len(migrationInfos) > 1 {
		// 	if migrationID, err := strconv.Atoi(migrationInfos[0]); err != nil {
		// 		utils.PrintErrorMsg(fmt.Errorf("Failed to create migration file, migration must start with XXX_MigrationName").Error())
		// 		return
		// 	}
		// }
	},
}

func init() {
	// Define flags and configuration settings here.
}
