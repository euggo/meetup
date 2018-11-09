package main

import "errors"

// ValidationError ...
type ValidationError struct {
	val string
}

func (e *ValidationError) Error() string {
	return "you screwed up: " + e.val
}

// Err{Name} ...
var (
	ErrInvalid = errors.New("not valid")
)

func main() {}

func valid() bool { return true }

// BGN1 OMIT
func yep() error {
	if !valid() {
		return ErrInvalid
	}
	return nil
}

// END1 OMIT

// BGN2 OMIT
func nope() error {
	var err *ValidationError = nil
	if !valid() {
		err = &ValidationError{"for reasons"}
	}
	return err // err != nil ... always
}

// END2 OMIT
