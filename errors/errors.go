package errors

type Error struct {
	Type string
	Msg  string
}

func (e *Error) Error() string {
	return e.Type + ":  " + e.Msg
}
