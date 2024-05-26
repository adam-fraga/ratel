package views

import (
	"fmt"
	"os"
	"path"

	"github.com/adam-fraga/ratel/internal/errors"
	model "github.com/adam-fraga/ratel/models/views"
	ut "github.com/adam-fraga/ratel/utils"
)

// PromptCreateComponent prompts the user to enter the component configuration
func PromptCreateComponent() error {

	var c model.Component

	os.Stdin.WriteString("Please enter the component name: ")
	fmt.Scanln(&c.Name)

	os.Stdin.WriteString("Please enter the component description: ")
	fmt.Scanln(&c.Description)

	ut.PrintInfoMsg(fmt.Sprintf("\nComponent configuration:\n\nComponent Name: %s\nComponent Description: %s\n",
		c.Name, c.Description))

	os.Stdin.WriteString("Is the configuration correct? (y/n): ")
	var response string

	fmt.Scanln(&response)

	if response == "n" {
		PromptCreateComponent()
	}

	err := createComponentFile(c)
	if err != nil {
		return &errors.DevError{Msg: "Error creating component file: " + err.Error()}
	}

	return nil

}

func createComponentFile(c model.Component) error {

	componentPath := path.Join(c.Path, c.Name)
	err := os.MkdirAll(componentPath, os.ModePerm)
	if err != nil {
		return &errors.DevError{
			Type:       "os.MkdirAll()",
			Origin:     "createComponentFile()",
			FileOrigin: "createComponent.go",
			Msg:        "Error creating component directory: " + err.Error(),
		}
	}

	ut.PrintSuccessMsg("Component created !")
	return nil
}
