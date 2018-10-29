package main

import (
	"fmt"

	"github.com/euggo/meetup/talks/to_err_is_human/code/novel"
)

func main() {
	ds, _ := novel.NewDataStore("INFORMATION", "Other", "Whatever")
	of, _ := novel.NewOutFile("out.txt")

	n, _ := of.Write(ds.Lowercased())
	fmt.Printf("copied %d bytes\n", n)

	of.Close()
	ds.Report()
}
