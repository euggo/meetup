package main

import (
	"flag"
	"fmt"
	"os"
)

//START1 OMIT
var (
	// slow enables a slowing of the digest function and may help users to // OMIT
	// better understand the implementation of concurrency // OMIT
	slow = false
	// OMIT
	// width controls the amount of goroutines running the digest function // OMIT
	width = 8
)

//END1 OMIT

//START2 OMIT
func main() {
	// define and parse flags // OMIT
	flag.BoolVar(&slow, "slow", slow, `slow processing to clarify behavior`)
	flag.IntVar(&width, "width", width, `set concurrency width`)
	flag.Parse()
	//END2 OMIT

	// get fileInfoGroup
	//START3 OMIT
	fig, err := newFileInfoGroup("./testfiles")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	//END3 OMIT

	// get new conch and setup cleanup
	//START4 OMIT
	c := newConch(fig)
	defer close(c.done())

	// get fileOutput and error channels // OMIT
	fos, errs := c.run()
	//END4 OMIT

	// print file contents
	//START5 OMIT
	for fo := range fos {
		fmt.Println(fo.path, fo.data, fo.err)
	}

	// print error, if any
	select {
	case err := <-errs:
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	default:
	}
}

//END5 OMIT
