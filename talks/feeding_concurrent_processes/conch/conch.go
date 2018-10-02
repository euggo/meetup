package main

import (
	"errors"
	"fmt"
	"sync"
)

// BGN3 OMIT
func produce(done <-chan struct{}, paths []string) (<-chan string, <-chan error) { // HLargs
	psc := make(chan string)   // HLchan
	esc := make(chan error, 1) // HLchan

	// BGN31 OMIT
	go func() { // feed paths chan, close chan when complete // HLgoroutine
		defer close(psc) // HLclose
		defer close(esc) // HLclose

		for _, p := range paths {
			select { // HLselect
			case psc <- p: // HLselect
			case <-done: // HLselect
				esc <- errors.New("canceled") // HLselect
				return                        // HLselect
			} // HLselect
		}
	}() // HLgoroutine
	// END31 OMIT

	return psc, esc // HLreturn
} // HLargs

// END3 OMIT

// BGN5 OMIT
func digest(done <-chan struct{}, fisc chan<- *fileInfo, psc <-chan string) { // HLargs
	for p := range psc { // HLreceive
		select { // HLselect
		case fisc <- newFileInfo(p): // HLselect
		case <-done: // HLselect
			return // HLselect
		} // HLselect
	} // HLreceive
} // HLargs

// END5 OMIT

// BGN4 OMIT
func consume(done <-chan struct{}, width int, psc <-chan string) <-chan *fileInfo { // HLargs
	fisc := make(chan *fileInfo) // HLchan

	// BGN41 OMIT
	go func() { // create consumers, close fileInfo channel when complete // HLgoroutine
		defer close(fisc) // HLclose

		var wg sync.WaitGroup // HLwg
		wg.Add(width)         // HLwg

		// BGN411 OMIT
		for i := 0; i < width; i++ { // HLwidth
			go func() { // run in goroutine to bind with wg.Done call // HLgogoroutine
				digest(done, fisc, psc) // HLdigest
				wg.Done()               // HLwg
			}() // HLgogoroutine
		} // HLwidth
		// END411 OMIT

		wg.Wait() // HLwg
	}() // HLgoroutine
	// END41 OMIT

	return fisc // HLreturn
} // HLargs

// END4 OMIT

// BGN2 OMIT
func fileInfos(done <-chan struct{}, width int, paths []string) (<-chan *fileInfo, func() error) { // HLargs
	psc, esc := produce(done, paths)  // HLproduce
	fisc := consume(done, width, psc) // HLconsume

	errFn := fileInfosErrorFunc(esc) // HLerror

	return fisc, errFn // HLreturn
} // HLargs

// END2 OMIT

// BGN21 OMIT
func fileInfosErrorFunc(esc <-chan error) func() error { // HLargs
	var last error // HLstate

	return func() error {
		if err := <-esc; err != nil { // HLerrorchan
			last = fmt.Errorf("cannot handle fileInfos: %s", err) // HLstate
			return last                                           // HLstate
		} // HLerrorchan

		return last // HLstate
	}
} // HLargs

// END21 OMIT
