package views

import (
	"fmt"
	"os"

	"github.com/adam-fraga/ratel/internal/errors"
	model "github.com/adam-fraga/ratel/models/views"
	ut "github.com/adam-fraga/ratel/utils"
)

func PromptLinkLayoutToPages(layout *model.Layout) error {

	var p model.Page

	os.Stdin.WriteString("Please enter the page name: ")
	fmt.Scanln(&p.Name)

	os.Stdin.WriteString("Please enter the page description: ")
	fmt.Scanln(&p.Description)

	ut.PrintInfoMsg(fmt.Sprintf("\nPage configuration:\n\nPage Name: %s\nPage Description: %s\n",
		p.Name, p.Description))

	//Ask the user if the configuration is correct
	os.Stdin.WriteString("Is the configuration correct? (y/n): ")
	var response string

	fmt.Scanln(&response)

	if response == "n" {
		PromptLinkLayoutToPages(layout)
	}

	err := p.Create()
	if err != nil {
		return &errors.DevError{
			Type:       "p.Create()",
			FileOrigin: "createPage.go",
			Origin:     "PromptLinkLayoutToPages()",
			Msg:        "Error creating page: " + err.Error(),
		}
	}

	createPageFile(p)

	return nil
}

func createPageFile(p model.Page) error {

	pagePath := p.Path
	err := os.MkdirAll(pagePath, os.ModePerm)
	if err != nil {
		return &errors.DevError{
			Type:       "os.MkdirAll()",
			Origin:     "createPageFile()",
			FileOrigin: "createPage.go",
			Msg:        "Error creating page directory: " + err.Error(),
		}
	}

	ut.PrintSuccessMsg("Page created !")
	return nil
}
