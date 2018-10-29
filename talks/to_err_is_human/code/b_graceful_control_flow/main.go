package main

import (
	"errors"
	"fmt"
	"os"
	"path"
)

func init() {
	os.Args[0] = "myapp"
}

// BGN OMIT
func main() {
	if err := run(); err != nil {
		cmd := path.Base(os.Args[0])
		fmt.Fprintf(os.Stderr, "%s: %s\n", cmd, err)
		os.Exit(1)
	}
}

func run() error {
	defer fmt.Println("hello")

	if err := alwaysFail(); err != nil {
		return err
	}
	return nil
}

// END OMIT
func alwaysFail() error {
	return errors.New("this always fails")
}
