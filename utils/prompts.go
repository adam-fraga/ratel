package utils

import (
	"fmt"
	"os"

	dt "github.com/adam-fraga/ratel/internal/datatypes"
	"github.com/adam-fraga/ratel/internal/errors"
)

// PromptDbConfig prompts the user to enter the database configuration
func PromptDbConfig(dbConfig *dt.DbUserConfig) (*dt.DbUserConfig, error) {

	os.Stdin.WriteString("Please enter the database user: ")
	fmt.Scanln(&dbConfig.DbUser)

	os.Stdin.WriteString("Please enter the database password: ")
	fmt.Scanln(&dbConfig.DbPassword)

	os.Stdin.WriteString("Please confirm password: ")
	var passwordConfirm string
	fmt.Scanln(&passwordConfirm)

	if dbConfig.DbPassword != passwordConfirm {
		err := &errors.ClientError{Msg: "Sorry your passwords do not match try again"}
		PrintErrorMsg(err.Error())
		PromptDbConfig(dbConfig)
	}

	os.Stdin.WriteString("Please enter the database name: ")
	fmt.Scanln(&dbConfig.DbName)

	PrintInfoMsg(fmt.Sprintf("\nDatabase configuration:\n\nDB Port: %s\nDB User: %s\nDB Name: %s\n",
		dbConfig.DbPort, dbConfig.DbUser, dbConfig.DbName))

	//Ask the user if the configuration is correct
	os.Stdin.WriteString("Is the configuration correct? (y/n): ")
	var response string

	fmt.Scanln(&response)

	if response == "n" {
		PromptDbConfig(dbConfig)
	}

	return dbConfig, nil
}
