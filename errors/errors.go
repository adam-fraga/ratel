package errors

import (
	"fmt"
)

type Error struct {
	Type       string
	Msg        string
	Origin     string
	FileOrigin string
}

func (e *Error) Error() string {
	return fmt.Sprintf("Error of type: %s\nIn file: %s\n Returned by: %s\n Error Message: %s\n", e.Type, e.FileOrigin, e.Origin, e.Msg)
}
