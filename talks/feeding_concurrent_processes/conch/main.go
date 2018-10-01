package main

import (
	"fmt"
	"os"
)

// BGN1 OMIT
func main() {
	paths, err := gzipFilePaths("./testfiles") // HLpaths
	if err != nil {                            // HLpaths
		logFatalln(err) // HLpaths
	} // HLpaths
	done := make(chan struct{}) // HLdonechan
	defer close(done)           // HLdonechan

	fis, fisErr := fileInfos(done, 4, paths) // HLprocess
	for fi := range fis {                    // HLiterate
		fmt.Println(fi.path, fi.data, fi.err) // HLiterate
	} // HLiterate

	if err = fisErr(); err != nil { // HLreport
		logFatalln(err) // HLreport
	} // HLreport
}

// END1 OMIT

func logFatalln(v ...interface{}) {
	fmt.Fprintln(os.Stderr, v...) //nolint
	os.Exit(1)
}
