package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/euggo/meetup/talks/to_err_is_human/code/novel"
)

func main() {
	var err error
	err = &novel.OutFileError{
		Name: "out.txt",
		Err:  io.ErrShortWrite,
	}

	for i := 0; i < 1; i++ {
		// BGN OMIT
		if oferr, ok := novel.AsOutFileError(err); ok {
			if oferr.Temporary() {
				time.Sleep(time.Second)
				continue
			}
			fmt.Fprintln(os.Stderr, oferr)
			break
		}
		// END OMIT
	}
}
