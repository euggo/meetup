package main

import (
	"fmt"
	"io"
	"log"

	"github.com/euggo/meetup/talks/to_err_is_human/code/novel"
)

func main() {
	ds, err := novel.NewDataStore("INFORMATION", "Other", "Whatever")
	defer ds.Report()
	if err != nil {
		log.Fatalln(err)
	}

	of, err := novel.NewOutFile("out.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer cleanClose(of) // log if of.Close returns error

	n, err := of.Write(ds.Lowercased())
	fmt.Printf("copied %d bytes\n", n)
	if err != nil {
		log.Fatalln(err)
	}
}

func cleanClose(f io.Closer) {
	if err := f.Close(); err != nil {
		log.Println(err)
	}
}
