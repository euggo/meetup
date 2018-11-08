package main

import (
	"fmt"
	"os"

	"github.com/euggo/meetup/talks/to_err_is_human/code/novel"
)

func main() {
	var retry bool
	err := novel.ErrDisconnected

	for i := 0; i < 1; i++ {
		// BGN OMIT
		if err != nil {
			if err == novel.ErrDisconnected {
				retry = true
				break
			}
			fmt.Fprint(os.Stderr, err)
		}
		// END OMIT
	}

	fmt.Printf("retry was set to %t\n", retry)
}
