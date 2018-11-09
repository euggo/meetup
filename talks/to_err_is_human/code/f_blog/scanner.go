package main

import (
	"bufio"
	"bytes"
)

func scanner() {
	token := ""
	input := &bytes.Buffer{}
	// BGN OMIT
	sc := bufio.NewScanner(input)
	for sc.Scan() {
		token = sc.Text()
		// process token
	}
	if err := sc.Err(); err != nil {
		// process the error
	}
	// END OMIT
	_ = token
}
