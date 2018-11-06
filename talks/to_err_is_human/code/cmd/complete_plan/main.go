package main

import (
	"fmt"
	"io"
	"os"
	"path"

	"github.com/euggo/meetup/talks/to_err_is_human/code/novel"
)

func init() {
	os.Args[0] = "someapp"
}

func main() {
	if err := run(); err != nil {
		cmd := path.Base(os.Args[0])
		fmt.Fprintf(os.Stderr, "%s: %s\n", cmd, err)
		os.Exit(1)
	}
}

func run() error {
	ds, err := novel.NewDataStore("INFORMATION", "Other", "Whatever")
	defer ds.Report()
	if err != nil {
		ds.RollBack()
		return err
	}

	of, err := novel.NewOutFile("out.txt")
	if err != nil {
		fmt.Printf("logging datastore:\n%s", ds.Lowercased())
		return err
	}
	defer cleanClose(of) // prints to stderr if returns error

	n, err := of.Write(ds.Lowercased())
	fmt.Printf("copied %d bytes\n", n)
	return err
}

func cleanClose(f io.Closer) {
	if err := f.Close(); err != nil {
		cmd := path.Base(os.Args[0])
		fmt.Fprintf(os.Stderr, "%s: %s\n", cmd, err)
	}
}
