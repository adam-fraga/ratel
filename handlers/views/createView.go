package views

import (
	"fmt"
	"os"
	"path"

	"github.com/adam-fraga/ratel/internal/errors"
	m "github.com/adam-fraga/ratel/models"
	ut "github.com/adam-fraga/ratel/utils"
)

// CreateView Create a file view of type (Component, Page, Layout)
func CreateView(viewType string, files []string) error {

	var v m.View
	v.Type = viewType

	for _, name := range files {
		ut.PrintInfoMsg(fmt.Sprintf("Creating %s %s", v.Type, name))
	}

	if len(files) > 1 {
		os.Stdin.WriteString("Create the following files ? (y/n): ")
		var response string

		fmt.Scanln(&response)

		if response == "n" {
			CreateView(viewType, files)
		}
	}

	for _, file := range files {
		v.Name = file
		if v.Name == "" || v.Type == "" {
			return &errors.ClientError{Msg: fmt.Sprintf("%s name cannot be empty", v.Type)}
		}

		if err := createViewFile(v); err != nil {
			return &errors.DevError{Msg: fmt.Sprintf("Error creating %s file :" + err.Error())}
		}
	}

	return nil

}

func createViewFile(v m.View) error {

	v.Path = "views/" + v.Type + "/"

	file, err := os.Create(path.Join(v.Path, v.Name+".templ"))
	defer file.Close()

	if err != nil {
		return &errors.DevError{
			Type:       "Creation view file error",
			Origin:     "createViewFile()",
			FileOrigin: "handlers/views/createView.go",
			Msg:        err.Error() + fmt.Sprintf("Error creating %v file\n", file)}
	}

	if err := os.Chmod(v.Path+v.Name+".templ", os.FileMode(0755)); err != nil {
		return &errors.DevError{
			Type:       "Creation view file error",
			Origin:     "createViewFile()",
			FileOrigin: "handlers/views/createView.go",
			Msg:        err.Error() + fmt.Sprintf("Error changing permissions for %v file\n", file)}
	}

	ut.PrintSuccessMsg(fmt.Sprintf("Creating %s %s successfuly...\n", v.Type, v.Name))
	return nil
}
