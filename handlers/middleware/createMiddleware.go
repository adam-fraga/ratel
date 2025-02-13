/*
Copyright © 2024 Adm FRG adam.fraga@live.fr
Package middleware contains handlers to execute the logic of the middleware system of ratel web framework
*/

package middleware

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/adam-fraga/ratel/internal/errors"
	ut "github.com/adam-fraga/ratel/utils"
)

type Middleware struct {
	Name string
	Path string
}

func (Middleware) New() Middleware {
	return Middleware{}
}

// CreateView Create a file view of type (Component, Page, Layout)
func (*Middleware) Create(mids []string) error {

	var m Middleware

	if len(mids) > 1 {
		ut.PrintInfoMsg(fmt.Sprintf("\n 🏗️ Creating multiple middlewares\n"))

		for _, mid := range mids {
			ut.PrintSuccessMsg(fmt.Sprintf("  📌 %s", mid))
		}

		ut.PrintWarningMsg("\n ⚠️ Confirm (Y/N):")
		var response string

		fmt.Scanln(&response)

		if strings.ToLower(response) != "y" {
			ut.PrintWarningMsg("\n ❌ Operation canceled.\n")
			return nil
		}
	}

	for _, mid := range mids {
		m.Name = mid
		if m.Name == "" {
			return &errors.MiddlewareError{
				Origin: "File: handlers/middleware/createMiddleware.go => Func: Create()",
				Msg:    fmt.Sprintf("Failed to create Middleware %s name cannot be empty", m.Name),
				Err:    nil,
			}
		}

		if err := m.CreateFile(m); err != nil {
			return &errors.MiddlewareError{
				Origin: "File: handlers/middleware/createMiddleware.go => Func: Create()",
				Msg:    "Failed to create Middleware, error: " + err.Error(),
				Err:    err,
			}
		}
	}
	return nil
}

func (*Middleware) CreateFile(m Middleware) error {

	m.Path = "src/middlewares/"

	file, err := os.Create(path.Join(m.Path, m.Name+".go"))
	defer file.Close()

	if err != nil {
		return &errors.MiddlewareError{
			Origin: "File: handlers/middleware/createMiddleware.go => Func: CreateFile()",
			Msg:    "Failed to create Middleware, error: " + err.Error(),
			Err:    err,
		}
	}

	if err := os.Chmod(m.Path+m.Name+".go", os.FileMode(0644)); err != nil {

		return &errors.MiddlewareError{
			Origin: "File: handlers/middleware/createMiddleware.go => Func: CreateFile()",
			Msg:    "Failed to set permission for the Middleware file, error: " + err.Error(),
			Err:    err,
		}
	}

	ut.PrintSuccessMsg(fmt.Sprintf(" ✅ %s%s.go successfuly created", m.Path, m.Name))
	return nil
}
