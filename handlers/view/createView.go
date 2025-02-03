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
			return &errors.ViewError{
				Origin: "File: handlers/views/createView.go => Func: CreateFile()",
				Msg:    "Failed create Project's view, view name cannot be empty",
				Err:    nil,
			}
		}

		if err := v.CreateFile(v); err != nil {
			return &errors.ViewError{
				Origin: "File: handlers/views/createView.go => Func: CreateFile()",
				Msg:    "Failed create Project's view, error creating view file",
				Err:    nil,
			}
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
		return &errors.ViewError{
			Origin: "File: handlers/views/createView.go => Func: CreateFile()",
			Msg:    "Failed create Project's view",
			Err:    err,
		}
	}

	if err := os.Chmod(v.Path+v.Name+".templ", os.FileMode(0644)); err != nil {
		return &errors.ViewError{
			Origin: "File: handlers/views/createView.go => Func: CreateFile()",
			Msg:    "Failed create Project's view, error setting permission for the view file",
			Err:    err}

	}

	ut.PrintSuccessMsg(fmt.Sprintf("     %s%s.go successfuly created", v.Path, v.Name))
	return nil
}
