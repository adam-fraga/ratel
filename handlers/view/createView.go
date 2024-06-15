/*
Copyright © 2024 Adm FRG adam.fraga@live.fr
Package views contains handlers to execute the logic of the views system of ratel web framework
*/

package view

import (
	"fmt"
	"os"
	"path"

	"github.com/adam-fraga/ratel/internal/errors"
	ut "github.com/adam-fraga/ratel/utils"
)

// View represent a view
type View struct {
	Name string
	Path string
	Type string
}

// New Create a new view
func (*View) New(viewType string) *View {
	return &View{
		Type: viewType,
	}
}

// Create view of type (Component, Page, Layout, Partial, Template or Metadata)
func (*View) Create(v *View, files []string) error {

	if len(files) > 1 {
		ut.PrintInfoMsg(fmt.Sprintf("\n   Creating multiple %s\n", v.Type))
		var response string

		for _, file := range files {
			ut.PrintSuccessMsg(fmt.Sprintf("     %s", file))
		}

		ut.PrintWarningMsg(fmt.Sprintf("\n   Confirm you to create the followings %s (y/n):", v.Type))

		fmt.Scanln(&response)

		if response == "n" {
			v.Create(v, files)
		}
	}

	for _, file := range files {
		v.Name = file
		if v.Name == "" || v.Type == "" {
			return &errors.ClientError{Msg: fmt.Sprintf("%s name cannot be empty", v.Type)}
		}

		if err := v.CreateFile(v); err != nil {
			return &errors.DevError{Msg: fmt.Sprintf("Error creating %s file :" + err.Error())}
		}
	}

	return nil

}

// CreateFile Create a file view of type (Component, Page, Layout, forms, Partial, Template or Metadata) in the appropriate folder
func (*View) CreateFile(v *View) error {

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

	ut.PrintSuccessMsg(fmt.Sprintf("     %s%s.go successfuly created", v.Path, v.Name))
	return nil
}
