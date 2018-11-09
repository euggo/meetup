package main

import (
	"fmt"

	"github.com/euggo/meetup/talks/to_err_is_human/code/novel"
)

func process(rune) error {
	return &novel.BadValueError{
		"3",
		"must contain only letters",
	}
}

// BGN OMIT
func parse(data string) error {
	for _, r := range data {
		if err := process(r); err != nil {
			return fmt.Errorf("cannot parse data: %s", err)
		}
	}

	return nil
}

// END OMIT

func main() {
	fmt.Println(parse("abc"))
}
