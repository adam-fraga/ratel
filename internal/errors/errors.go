/*
Copyright © 2024 Adm FRG adam.fraga@live.fr
Package errors contains the custom error types that are used in the ratel web framework
*/

package errors

import (
	"fmt"

	"github.com/fatih/color"
)

type ProjectError struct {
	Origin string
	Msg    string
	Err    error
}

func (e *ProjectError) Error() string {
	printer := color.New(color.FgRed).SprintfFunc()
	return fmt.Sprintf(printer(e.Msg))
}

func (e *ProjectError) Unwrap() error {
	return e.Err
}

type DBError struct {
	Origin string
	Msg    string
	Err    error
}

func (e *DBError) Error() string {
	printer := color.New(color.FgRed).SprintfFunc()
	return fmt.Sprintf(printer(e.Msg))
}

func (e *DBError) Unwrap() error {
	return e.Err
}
