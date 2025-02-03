/*
Copyright © 2024 Adm FRG adam.fraga@live.fr
Package errors contains the custom error types that are used in the ratel web framework
*/

package errors

import (
	"github.com/fatih/color"
)

type ProjectError struct {
	Origin string
	Msg    string
	Err    error
}

func (e *ProjectError) Error() string {
	printer := color.New(color.FgRed).SprintfFunc()
	return printer(e.Msg)
}

func (e *ProjectError) Unwrap() error {
	return e.Err
}

type HandlerError struct {
	Origin string
	Msg    string
	Err    error
}

func (e *HandlerError) Error() string {
	printer := color.New(color.FgRed).SprintfFunc()
	return printer(e.Msg)
}

func (e *HandlerError) Unwrap() error {
	return e.Err
}

type MiddlewareError struct {
	Origin string
	Msg    string
	Err    error
}

func (e *MiddlewareError) Error() string {
	printer := color.New(color.FgRed).SprintfFunc()
	return printer(e.Msg)
}

func (e *MiddlewareError) Unwrap() error {
	return e.Err
}

type ModelError struct {
	Origin string
	Msg    string
	Err    error
}

func (e *ModelError) Error() string {
	printer := color.New(color.FgRed).SprintfFunc()
	return printer(e.Msg)
}

func (e *ModelError) Unwrap() error {
	return e.Err
}

type ViewError struct {
	Origin string
	Msg    string
	Err    error
}

func (e *ViewError) Error() string {
	printer := color.New(color.FgRed).SprintfFunc()
	return printer(e.Msg)
}

func (e *ViewError) Unwrap() error {
	return e.Err
}

type DBError struct {
	Origin string
	Msg    string
	Err    error
}

func (e *DBError) Error() string {
	printer := color.New(color.FgRed).SprintfFunc()
	return printer(e.Msg)
}

func (e *DBError) Unwrap() error {
	return e.Err
}
