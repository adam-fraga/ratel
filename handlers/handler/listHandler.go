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
		return &er.ClientError{Msg: fmt.Sprintf("Error getting the files to show: " + err.Error())}
	}

	m.printHandlersToStdout(&m)

	return nil
}

// printFilesToStdOut function to print the files to the stdout
func (*Handlers) printHandlersToStdout(m *Handlers) {
	m.totalFiles = 0
	var count uint8

	for _, mid := range m.handlers {
		count++
		if count == 1 {
			ut.PrintInfoMsg(fmt.Sprintf("\n   ***Handlers***\n"))
		}
		m.totalFiles++
		ut.PrintSuccessMsg(fmt.Sprintf("     %s%s", mid.Path, mid.Name))
	}

	ut.PrintInfoMsg("\n   TOTAL")
	ut.PrintSuccessMsg(fmt.Sprintf("     %d\n", m.totalFiles))
}

// getHandlerFiles function to get the handlers files from the directory
func (*Handlers) getHandlerFiles(m *Handlers) error {
	path := "handlers/"
	files, err := os.Open(path)
	defer files.Close()
	if err != nil {
		return &er.ClientError{Msg: fmt.Sprintf("Error opening the directory to get the handlers")}
	}
	for {
		file, err := files.Readdir(1)
		if err != nil {
			break
		}
		m.handlers = append(m.handlers, Handler{Name: file[0].Name(), Path: path})
	}
	return nil
}
