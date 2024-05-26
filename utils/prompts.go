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

// PromptCreateLayout prompts the user to enter the layout configuration
// func PromptCreateLayout(*dt.Layout) error {
//
// 	var layout dt.Layout
//
// 	os.Stdin.WriteString("Please enter the layout name: ")
// 	fmt.Scanln(&layout.Name)
//
// 	os.Stdin.WriteString("Please enter the layout description: ")
// 	fmt.Scanln(&layout.Description)
//
// 	PrintInfoMsg(fmt.Sprintf("\nLayout configuration:\n\nLayout Name: %s\nLayout Description: %s\n",
// 		layout.Name, layout.Description))
//
// 	//Ask the user if the configuration is correct
// 	os.Stdin.WriteString("Is the configuration correct? (y/n): ")
// 	var response string
//
// 	fmt.Scanln(&response)
//
// 	if response == "n" {
// 		PromptCreateLayout(&layout)
// 	}
//
// 	return nil
// }
//
// // PromptLinkLayoutToPages prompts the user to enter the page configuration
// func PromptLinkLayoutToPages(layout *dt.Layout) error {
//
// 	var page dt.Page
//
// 	os.Stdin.WriteString("Please enter the page name: ")
// 	fmt.Scanln(&page.Name)
//
// 	os.Stdin.WriteString("Please enter the page description: ")
// 	fmt.Scanln(&page.Description)
//
// 	PrintInfoMsg(fmt.Sprintf("\nPage configuration:\n\nPage Name: %s\nPage Description: %s\n",
// 		page.Name, page.Description))
//
// 	//Ask the user if the configuration is correct
// 	os.Stdin.WriteString("Is the configuration correct? (y/n): ")
// 	var response string
//
// 	fmt.Scanln(&response)
//
// 	if response == "n" {
// 		PromptLinkLayoutToPages(layout)
// 	}
//
// 	return nil
// }
//
// // PromptCreateComponent prompts the user to enter the component configuration
// func PromptCreateComponent(*dt.Component) error {
//
// 	var component dt.Component
//
// 	os.Stdin.WriteString("Please enter the component name: ")
// 	fmt.Scanln(&component.Name)
//
// 	os.Stdin.WriteString("Please enter the component description: ")
// 	fmt.Scanln(&component.Description)
//
// 	PrintInfoMsg(fmt.Sprintf("\nComponent configuration:\n\nComponent Name: %s\nComponent Description: %s\n",
// 		component.Name, component.Description))
//
// 	//Ask the user if the configuration is correct
// 	os.Stdin.WriteString("Is the configuration correct? (y/n): ")
// 	var response string
//
// 	fmt.Scanln(&response)
//
// 	if response == "n" {
// 		PromptCreateComponent(&component)
// 	}
//
// 	return nil
// }
