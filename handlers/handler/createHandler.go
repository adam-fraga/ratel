/*
Copyright © 2024 Adm FRG adam.fraga@live.fr
Package handler contains handlers to execute the logic of the handler system of the ratel web framework
*/

package handler

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/adam-fraga/ratel/internal/errors"
	ut "github.com/adam-fraga/ratel/utils"
)

// Handler represent a handler
type Handler struct {
	Name string
	Path string
}

// New Create a new handler
func (*Handler) New() *Handler {
	return &Handler{}
}

// Create a new handler
func (*Handler) Create(handlers []string) error {

	var h Handler

	if len(handlers) > 1 {
		ut.PrintInfoMsg("\n 🏗️ Creating multiple handlers\n")

		for _, handler := range handlers {
			ut.PrintSuccessMsg(fmt.Sprintf("    📌 %s", handler))
		}

		ut.PrintWarningMsg("\n ⚠️ Confirm (Y/N):")
		var response string

		fmt.Scanln(&response)

		if strings.ToLower(response) != "y" {
			ut.PrintWarningMsg("\n ❌ Operation canceled.\n")
			return nil
		}
	}

	for _, handler := range handlers {
		h.Name = handler
		if h.Name == "" {
			return &errors.HandlerError{
				Origin: "File: handlers/handler/createHandler.go => Func: Create() ",
				Msg:    "Handler name cannot be empty",
				Err:    nil,
			}
		}

		if err := h.CreateFile(h); err != nil {
			return &errors.HandlerError{
				Origin: "File: handlers/handler/createHandler.go => Func: Create()",
				Msg:    "Failed to create the handler file, error: " + err.Error(),
				Err:    err,
			}
		}
	}
	return nil
}

// CreateFile create a new handler file in the handlers directory
func (*Handler) CreateFile(h Handler) error {

	h.Path = "src/handlers/"

	file, err := os.Create(path.Join(h.Path, h.Name+".go"))
	defer file.Close()

	if err != nil {
		return &errors.HandlerError{
			Origin: "File: handlers/handlers/createHandler.go => Func: CreateFile()",
			Msg:    fmt.Sprintf("Failed to create %v file, error: %s\n", file, err.Error()),
			Err:    err,
		}
	}

	if err := os.Chmod(h.Path+h.Name+".go", os.FileMode(0644)); err != nil {
		return &errors.HandlerError{
			Origin: "File: handlers/handlers/createhandler.go => Func: CreateFile()",
			Msg:    fmt.Sprintf("Error changing permissions for %v file, error: %s\n", file, err.Error()),
			Err:    err,
		}
	}

	ut.PrintSuccessMsg(fmt.Sprintf(" ✅ %s%s.go successfuly created", h.Path, h.Name))
	return nil
}
