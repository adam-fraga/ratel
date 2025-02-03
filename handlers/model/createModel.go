/*
Copyright © 2024 Adm FRG adam.fraga@live.fr
Package middleware contains handlers to execute the logic of the middleware system of ratel web framework
*/

package model

import (
	"fmt"
	"os"
	"path"

	"github.com/adam-fraga/ratel/internal/errors"
	ut "github.com/adam-fraga/ratel/utils"
)

type Model struct {
	Name string
	Path string
}

func (*Model) New() *Model {
	return &Model{}
}

// Createmodel Create a file model of type (Component, Page, Layout)
func (*Model) Create(mids []string) error {

	var m Model

	if len(mids) > 1 {
		ut.PrintInfoMsg(fmt.Sprintf("\n   Creating multiple models\n"))

		for _, mid := range mids {
			ut.PrintSuccessMsg(fmt.Sprintf("     %s", mid))
		}

		ut.PrintWarningMsg("\n   Confirm you to create the followings models (y/n):")
		var response string

		fmt.Scanln(&response)

		if response == "n" {
			m.Create(mids)
		}
	}

	for _, mid := range mids {
		m.Name = mid
		if m.Name == "" {
			return &errors.ModelError{
				Origin: "File: handlers/model/createModel.go => Func:CreateFile()",
				Msg:    "Failed to create Model, name cannot be empty",
				Err:    nil,
			}
		}

		if err := m.CreateFile(m); err != nil {
			return &errors.ModelError{
				Origin: "File: handlers/model/createModel.go => Func:CreateFile()",
				Msg:    "Failed to create Model file",
				Err:    err,
			}
		}
	}
	return nil
}

func (*Model) CreateFile(m Model) error {

	m.Path = "models/"

	file, err := os.Create(path.Join(m.Path, m.Name+".go"))
	defer file.Close()

	if err != nil {
		return &errors.ModelError{
			Origin: "File: handlers/model/createModel.go => Func:CreateFile()",
			Msg:    "Failed to create Model",
			Err:    err,
		}
	}

	if err := os.Chmod(m.Path+m.Name+".go", os.FileMode(0644)); err != nil {
		return &errors.ModelError{
			Origin: "File: handlers/model/createModel.go => Func:CreateFile()",
			Msg:    "Failed to set permission for Model file",
			Err:    err,
		}
	}

	ut.PrintSuccessMsg(fmt.Sprintf("     %s%s.go successfuly created", m.Path, m.Name))
	return nil
}
