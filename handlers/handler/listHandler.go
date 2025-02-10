package handler

import (
	"fmt"
	"os"

	er "github.com/adam-fraga/ratel/internal/errors"
	ut "github.com/adam-fraga/ratel/utils"
)

// Handler represent a handler in the project
type Handlers struct {
	totalFiles uint16
	handlers   []Handler
}

// List function to list the handlers
func List() error {

	var m Handlers

	if err := m.getHandlerFiles(&m); err != nil {
		return &er.HandlerError{
			Origin: "File: handlers/handler/listHandler.go => Func: List()",
			Msg:    fmt.Sprintf("Failed to list handlers in the project"),
			Err:    err,
		}
	}

	m.printHandlersToStdout(&m)

	return nil
}

// printFilesToStdOut function to print the files to the stdout
func (*Handlers) printHandlersToStdout(h *Handlers) {
	ut.PrintListToStdout("Handlers", h.handlers, func(hdl Handler) string {
		return fmt.Sprintf("  📂 %s%s", hdl.Path, hdl.Name)
	})
}

// getHandlerFiles function to get the handlers files from the directory
func (*Handlers) getHandlerFiles(m *Handlers) error {
	path := "src/handlers/"
	files, err := os.Open(path)
	defer files.Close()
	if err != nil {
		return &er.HandlerError{
			Origin: "File: handlers/handler/listHandler.go => Func: List()",
			Msg:    "Failed to get the handlers from project directory, error: " + err.Error(),
			Err:    err,
		}
	}
	for {
		file, err := files.Readdir(1)
		if err != nil {
			ut.PrintErrorMsg(fmt.Sprintf("Failed to read handler %v, error: %s", file, err.Error()))
			break
		}
		m.handlers = append(m.handlers, Handler{Name: file[0].Name(), Path: path})
	}
	return nil
}
