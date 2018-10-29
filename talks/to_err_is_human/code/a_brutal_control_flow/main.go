package main

import (
	"errors"
	"fmt"
	"os"
)

// BGN OMIT
func main() {
	defer fmt.Println("hello")

	if err := alwaysFail(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// END OMIT

func alwaysFail() error {
	return errors.New("this always fails")
}
