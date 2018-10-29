package main

import (
	"fmt"
	"os"
	"path"

	"github.com/euggo/meetup/talks/to_err_is_human/code/novel"
)

func init() {
	os.Args[0] = "improvedapp"
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
		fmt.Println("logging datastore:", ds.Lowercased())
		return err
	}

	n, err := of.Write(ds.Lowercased())
	fmt.Printf("copied %d bytes\n", n)
	if err != nil {
		return err
	}

	return of.Close()
}
