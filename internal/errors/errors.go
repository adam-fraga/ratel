package errors

import (
	"fmt"
	"github.com/fatih/color"
)

// DevError is a struct that represents an error that is thrown by the developer for the developer
type DevError struct {
	Type       string
	Msg        string
	Origin     string
	FileOrigin string
}

func (e *DevError) Error() string {
	printer := color.New(color.FgRed).SprintfFunc()
	return fmt.Sprintf(printer("Error of type: %s\nIn file: %s\nReturned by: %s\nError Message: %s\n",
		e.Type, e.FileOrigin, e.Origin, e.Msg))
}

// ClientError is a struct that represents an error that is thrown by the developer for the user
type ClientError struct {
	Msg string
}

func (e *ClientError) Error() string {
	printer := color.New(color.FgRed).SprintfFunc()
	return fmt.Sprintf(printer("Error: %s\n", e.Msg))
}
