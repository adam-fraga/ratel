package middleware

import (
	"fmt"
	"os"

	er "github.com/adam-fraga/ratel/internal/errors"
	ut "github.com/adam-fraga/ratel/utils"
)

// ViewFiles struct to hold the view files
type Middlewares struct {
	totalFiles  uint16
	middlewares []Middleware
}

// Middleware struct to hold the middleware files
func List() error {

	var m Middlewares

	if err := m.getMiddlewareFiles(&m); err != nil {
		return &er.ClientError{Msg: fmt.Sprintf("Error getting the files to show: " + err.Error())}
	}

	m.printMiddlewaresToStdout(&m)

	return nil
}

// printFilesToStdOut function to print the files to the stdout
func (*Middlewares) printMiddlewaresToStdout(m *Middlewares) {
	m.totalFiles = 0
	var count uint8

	for _, mid := range m.middlewares {
		count++
		if count == 1 {
			ut.PrintInfoMsg(fmt.Sprintf("\n   ***Middlewares***\n"))
		}
		m.totalFiles++
		ut.PrintSuccessMsg(fmt.Sprintf("     %s%s", mid.Path, mid.Name))
	}

	ut.PrintInfoMsg("\n   TOTAL")
	ut.PrintSuccessMsg(fmt.Sprintf("     %d\n", m.totalFiles))
}

// getMiddlewareFiles function to get the middleware files from the directory
func (*Middlewares) getMiddlewareFiles(m *Middlewares) error {
	path := "middlewares/"
	files, err := os.Open(path)
	defer files.Close()
	if err != nil {
		return &er.ClientError{Msg: fmt.Sprintf("Error opening the directory to get the middlewares")}
	}
	for {
		file, err := files.Readdir(1)
		if err != nil {
			break
		}
		m.middlewares = append(m.middlewares, Middleware{Name: file[0].Name(), Path: path})
	}
	return nil
}
