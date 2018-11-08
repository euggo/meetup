package novel

import "errors"

// BGN1 OMIT
// BadValueError manages the scope and message related to a bad input value.
type BadValueError struct {
	Val string
	Msg string
}

// Error satisfies the error interface.
func (e *BadValueError) Error() string {
	return "bad value '" + e.Val + "': " + e.Msg
}

// END1 OMIT

// BGN2 OMIT
var (
	ErrExplicitInput = errors.New("system will not handle explicit input")
	ErrDisconnected  = errors.New("system dependency disconnected unexpectedly")
)

// END2 OMIT

// BGN3 OMIT
// OutFileError manages data related to an OutFile handling failure.
type OutFileError struct {
	Name string
	Err  error
	Temp bool
}

// Error satisfies the error interface.
func (e *OutFileError) Error() string {
	return "failure handling " + e.Name + ": " + e.Err.Error()
}

// Temporary indicates whether the error is temporary and can be tried later.
func (e *OutFileError) Temporary() bool {
	return e.Temp
}

// END3 OMIT

// BGN4 OMIT
// AsOutFileError is a convenience assertion function.
func AsOutFileError(err error) (*OutFileError, bool) {
	if err == nil {
		return nil, false
	}
	oferr, ok := err.(*OutFileError)
	return oferr, ok
}

// END4 OMIT
