package main

import (
	"errors"
	"fmt"
	"os"
	"path"
	"time"
)

func init() {
	os.Args[0] = "myapp"
}

// BGN OMIT
func main() {
	if code, err := run(); err != nil {
		cmd := path.Base(os.Args[0])
		fmt.Fprintf(os.Stderr, "%s: %s\n", cmd, err)
		os.Exit(code)
	}
}

func run() (int, error) {
	defer fmt.Println("hello")

	if err := mightFail(); err != nil {
		return 2, err
	}
	if err := alwaysFail(); err != nil {
		return 1, err
	}
	return 0, nil
}

// END OMIT

func mightFail() error {
	if time.Now().UnixNano()%2 == 0 {
		return errors.New("this might fail")
	}
	return nil
}

func alwaysFail() error {
	return errors.New("this always fails")
}
