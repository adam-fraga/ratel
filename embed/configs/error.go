package errors

import (
	"fmt"
)

type CustomError struct {
	Message string
	Status  int
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("Error: %s", e.Message)
}

func New(message string, status int) *CustomError {
	return &CustomError{Message: message, Status: status}
}
