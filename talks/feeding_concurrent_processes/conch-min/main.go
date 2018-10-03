package main

import (
	"fmt"
	"os"
)

// BGN1 OMIT
func main() { // HLargs
	paths, err := gzipFilePaths("./testdata") // HLpaths
	if err != nil {                           // HLpaths
		logFatalln(err) // HLpaths
	} // HLpaths

	fisc := fileInfos(4, paths) // HLprocess
	for fi := range fisc {      // HLiterate
		fmt.Println(fi.path, fi.data, fi.err) // HLiterate
	} // HLiterate
} // HLargs

// END1 OMIT

func logFatalln(v ...interface{}) {
	fmt.Fprintln(os.Stderr, v...) //nolint
	os.Exit(1)
}
