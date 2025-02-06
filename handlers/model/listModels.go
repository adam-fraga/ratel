package model

import (
	"fmt"
	"os"

	er "github.com/adam-fraga/ratel/internal/errors"
	ut "github.com/adam-fraga/ratel/utils"
)

// ViewFiles struct to hold the view files
type Models struct {
	totalFiles uint16
	models     []Model
}

// Model struct to hold the middleware files
func List() error {

	var m Models

	if err := m.getModelFiles(&m); err != nil {
		return &er.ModelError{
			Origin: "File: handlers/model/listModels.go => Func: List()",
			Msg:    "Error getting project's models, error: " + err.Error(),
			Err:    err,
		}
	}

	m.printmodelsToStdout(&m)

	return nil
}

// printFilesToStdOut function to print the files to the stdout
func (*Models) printmodelsToStdout(m *Models) {
	m.totalFiles = 0
	var count uint8

	for _, mid := range m.models {
		count++
		if count == 1 {
			ut.PrintInfoMsg(fmt.Sprintf("\n ***Models***\n"))
		}
		m.totalFiles++
		ut.PrintSuccessMsg(fmt.Sprintf("  %s%s", mid.Path, mid.Name))
	}

	ut.PrintInfoMsg("\n  TOTAL")
	ut.PrintSuccessMsg(fmt.Sprintf("  %d\n", m.totalFiles))
}

// getModelFiles function to get the middleware files from the directory
func (*Models) getModelFiles(m *Models) error {
	path := "src/models/"
	files, err := os.Open(path)
	defer files.Close()
	if err != nil {
		return &er.ModelError{
			Origin: "File: handlers/model/listModels.go => Func: List()",
			Msg:    "Error getting project's models, error: " + err.Error(),
			Err:    err,
		}
	}
	for {
		file, err := files.Readdir(1)
		if err != nil {
			break
		}
		m.models = append(m.models, Model{Name: file[0].Name(), Path: path})
	}
	return nil
}
