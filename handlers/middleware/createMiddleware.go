package views

import (
	"fmt"
	"os"
	"path"

	"github.com/adam-fraga/ratel/internal/errors"
	ut "github.com/adam-fraga/ratel/utils"
)

type Middleware struct {
	Name        string
	Description string
	Path        string `default:"views/middlewares"`
}

// PromptCreateComponent prompts the user to enter the component configuration
func CreateMiddleware() error {

	var m Middleware

	os.Stdin.WriteString("Please enter the middleware name: ")
	fmt.Scanln(&m.Name)

	if m.Name == "" {
		return &errors.ClientError{Msg: "middleware name cannot be empty"}
	}

	os.Stdin.WriteString("Please enter the middleware description: ")
	fmt.Scanln(&m.Description)

	ut.PrintInfoMsg(fmt.Sprintf("\nMiddleware configuration:\n\nMiddleware Name: %s\nMiddleware Description: %s\n",
		m.Name, m.Description))

	os.Stdin.WriteString("Is the configuration correct? (y/n): ")
	var response string

	fmt.Scanln(&response)

	if response == "n" {
		CreateMiddleware()
	}

	err := CreateMiddlewareFile(m)
	if err != nil {
		return &errors.DevError{Msg: "Error creating middleware file: " + err.Error()}
	}

	return nil

}

func CreateMiddlewareFile(m Middleware) error {

	componentPath := path.Join(m.Path, m.Name)
	err := os.MkdirAll(componentPath, os.ModePerm)
	if err != nil {
		return &errors.DevError{
			Type:       "os.MkdirAll()",
			Origin:     "CreateMiddlewareFile()",
			FileOrigin: "createMiddleware.go",
			Msg:        "Error creating component directory: " + err.Error(),
		}
	}

	ut.PrintSuccessMsg("Middleware created !")
	return nil
}
