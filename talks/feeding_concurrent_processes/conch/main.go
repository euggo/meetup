package main

import (
	"fmt"
	"os"
)

// BGN1 OMIT
func main() {
	paths, err := gzipFilePaths("./testfiles")
	if err != nil {
		fmt.Printf("cannot get paths: %s\n", err)
		os.Exit(1)
	}
	c := newConch() // HLcreate

	fis, runErr := c.run(4, paths) // HLrun
	for fi := range fis {          // HLiterate
		fmt.Println(fi.path, fi.data, fi.err) // HLiterate
	} // HLiterate

	if err = runErr(); err != nil { // HLreport
		fmt.Printf("run error: %s\n", err) // HLreport
	} // HLreport
}

// END1 OMIT
