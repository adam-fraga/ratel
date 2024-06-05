/*
Copyright © 2024 Adm FRG adam.fraga@live.fr
Package errors contains the custom error types that are used in the ratel web framework
*/

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
	return fmt.Sprintf(printer("APP Error of type: %s\nIn file: %s\nReturned by: %s\nError Message: %s\n",
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

type DbError struct {
	Query  string
	Msg    string
	Action string
}

func (e *DbError) Error() string {
	printer := color.New(color.FgRed).SprintfFunc()
	return fmt.Sprintf(printer("DB Error on action: %s\nQuery: %s\nError Message: %s\n", e.Action, e.Query, e.Msg))
}
