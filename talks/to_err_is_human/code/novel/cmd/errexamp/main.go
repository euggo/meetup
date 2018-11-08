package main

import (
	"fmt"

	"github.com/euggo/meetup/talks/to_err_is_human/code/novel"
)

func main() {
	// BGN2 OMIT
	err := &novel.BadValueError{
		"3",
		"must contain only letters",
	}

	fmt.Println(err)
	// END2 OMIT
}
