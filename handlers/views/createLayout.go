package views

import (
	"fmt"
	"os"

	"github.com/adam-fraga/ratel/internal/db"
	"github.com/adam-fraga/ratel/internal/errors"
	model "github.com/adam-fraga/ratel/models/views"
	ut "github.com/adam-fraga/ratel/utils"
)

func PromptCreateLayout() error {

	var l model.Layout
	var db db.Db

	if err := db.Connect(); err != nil {
		return &errors.DevError{Msg: "Error connecting to the database: " + err.Error()}
	}
	defer db.Close()

	os.Stdin.WriteString("Please enter the layout name: ")
	fmt.Scanln(&l.Name)

	os.Stdin.WriteString("Please enter the layout description: ")
	fmt.Scanln(&l.Description)

	ut.PrintInfoMsg(fmt.Sprintf("\nLayout configuration:\n\nLayout Name: %s\nLayout Description: %s\n",
		l.Name, l.Description))

	//Ask the user if the configuration is correct
	os.Stdin.WriteString("Is the configuration correct? (y/n): ")
	var response string

	fmt.Scanln(&response)

	if response == "n" {
		PromptCreateLayout()
	}

	l.Create()
	createLayoutFile(l)

	return nil
}

func createLayoutFile(l model.Layout) error {

	layoutPath := l.Path
	err := os.MkdirAll(layoutPath, os.ModePerm)
	if err != nil {
		return &errors.DevError{
			Type:       "os.MkdirAll()",
			Origin:     "createLayoutFile()",
			FileOrigin: "createLayout.go",
			Msg:        "Error creating layout directory: " + err.Error(),
		}
	}

	ut.PrintSuccessMsg("Layout created !")
	return nil
}
