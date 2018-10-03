package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// BGN1 OMIT
func main() { // HLargs
	paths, err := gzipFilePaths("./testdata") // HLpaths
	if err != nil {                           // HLpaths
		logFatalln(err) // HLpaths
	} // HLpaths
	done := make(chan struct{}) // HLdonechan
	closeOnTerm(done)           // close done channel if SIGTERM is received // HLdonechan

	fis, fisErr := fileInfos(done, 4, paths) // HLprocess
	for fi := range fis {                    // HLiterate
		fmt.Println(fi.path, fi.data, fi.err) // HLiterate
	} // HLiterate

	if err = fisErr(); err != nil { // HLreport
		logFatalln(err) // HLreport
	} // HLreport
} // HLargs

// END1 OMIT

func logFatalln(v ...interface{}) {
	fmt.Fprintln(os.Stderr, v...) //nolint
	os.Exit(1)
}

func closeOnTerm(done chan struct{}) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM)

	go func() {
		<-c
		close(done)
	}()
}
