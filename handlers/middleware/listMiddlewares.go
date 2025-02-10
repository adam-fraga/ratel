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
		return &er.MiddlewareError{
			Origin: "File: handlers/middleware/listMiddlewares.go => Func: List()",
			Msg:    "Failed to list project's middlewares, error: " + err.Error(),
			Err:    err,
		}
	}
	m.printMiddlewaresToStdout()
	return nil
}

func (m *Middlewares) printMiddlewaresToStdout() {
	ut.PrintListToStdout("Middlewares", m.middlewares, func(mid Middleware) string {
		return fmt.Sprintf("  📂 %s%s", mid.Path, mid.Name)
	})
}

// getMiddlewareFiles function to get the middleware files from the directory
func (*Middlewares) getMiddlewareFiles(m *Middlewares) error {
	path := "src/middlewares/"
	files, err := os.Open(path)
	defer files.Close()
	if err != nil {
		return &er.MiddlewareError{
			Origin: "File: handlers/middleware/listMiddlewares.go => Func: List()",
			Msg:    "Failed get project's middlewares from directory, error: " + err.Error(),
			Err:    err,
		}
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
