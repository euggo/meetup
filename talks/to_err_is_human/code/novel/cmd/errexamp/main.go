package main

import (
	"fmt"

	"github.com/euggo/meetup/talks/to_err_is_human/code/novel"
)

func main() {
	var err error
	// BGN OMIT
	err = &novel.BadValueError{
		"3",
		"must contain only letters",
	}

	fmt.Println(err)
	// END OMIT
}
