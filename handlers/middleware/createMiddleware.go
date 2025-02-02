/*
Copyright © 2024 Adm FRG adam.fraga@live.fr
Package middleware contains handlers to execute the logic of the middleware system of ratel web framework
*/

package middleware

import (
	"fmt"
	"os"
	"path"

	"github.com/adam-fraga/ratel/internal/errors"
	ut "github.com/adam-fraga/ratel/utils"
)

type Middleware struct {
	Name string
	Path string
}

func (*Middleware) New() *Middleware {
	return &Middleware{}
}

// CreateView Create a file view of type (Component, Page, Layout)
func (*Middleware) Create(mids []string) error {

	var m Middleware

	if len(mids) > 1 {
		ut.PrintInfoMsg(fmt.Sprintf("\n   Creating multiple middlewares\n"))

		for _, mid := range mids {
			ut.PrintSuccessMsg(fmt.Sprintf("     %s", mid))
		}

		ut.PrintWarningMsg("\n   Confirm you to create the followings middleware (y/n):")
		var response string

		fmt.Scanln(&response)

		if response == "n" {
			m.Create(mids)
		}
	}

	for _, mid := range mids {
		m.Name = mid
		if m.Name == "" {
			return &errors.ClientError{Msg: fmt.Sprintf("%s name cannot be empty", m.Name)}
		}

		if err := m.CreateFile(m); err != nil {
			return &errors.DevError{Msg: fmt.Sprintf("Error creating %s file :" + err.Error())}
		}
	}

	return nil

}

func (*Middleware) CreateFile(m Middleware) error {

	m.Path = "middlewares/"

	file, err := os.Create(path.Join(m.Path, m.Name+".go"))
	defer file.Close()

	if err != nil {
		return &errors.DevError{
			Type:       "Creation view file error",
			Origin:     "createViewFile()",
			FileOrigin: "handlers/middlewares/createMiddleware.go",
			Msg:        err.Error() + fmt.Sprintf("Error creating %v file\n", file)}
	}

	if err := os.Chmod(m.Path+m.Name+".go", os.FileMode(0644)); err != nil {
		return &errors.DevError{
			Type:       "Creation view file error",
			Origin:     "createViewFile()",
			FileOrigin: "handlers/views/createView.go",
			Msg:        err.Error() + fmt.Sprintf("Error changing permissions for %v file\n", file)}
	}

	ut.PrintSuccessMsg(fmt.Sprintf("     %s%s.go successfuly created", m.Path, m.Name))
	return nil
}
