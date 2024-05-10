package errors

import (
	"fmt"
)

type Error struct {
	Type string
	Msg  string
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s Error: %s", e.Type, e.Msg)
}
