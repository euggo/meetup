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
	go func() { // feed paths chan, close chan when complete
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
	}()
	// END31 OMIT

	return psc, esc // HLreturn
}

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
}

// END5 OMIT

// BGN4 OMIT
func consume(done <-chan struct{}, width int, psc <-chan string) <-chan *fileInfo { // HLargs
	fisc := make(chan *fileInfo) // HLchan

	// BGN41 OMIT
	go func() { // create consumers, close fileInfo channel when complete
		defer close(fisc) // HLclose

		var wg sync.WaitGroup // HLwg
		wg.Add(width)         // HLwg

		// BGN411 OMIT
		for i := 0; i < width; i++ { // HLwidth
			go func() { // run in goroutine to bind with wg.Done call
				digest(done, fisc, psc) // HLdigest
				wg.Done()               // HLwg
			}()
		} // HLwidth
		// END411 OMIT

		wg.Wait() // HLwg
	}()
	// END41 OMIT

	return fisc // HLreturn
}

// END4 OMIT

// BGN2 OMIT
func fileInfos(done <-chan struct{}, width int, paths []string) (<-chan *fileInfo, func() error) { // HLargs
	psc, esc := produce(done, paths)  // HLproduce
	fisc := consume(done, width, psc) // HLconsume

	var last error          // HLerror
	errFn := func() error { // HLerror
		if err := <-esc; err != nil { // HLerror
			last = fmt.Errorf("cannot handle fileInfos: %s", err) // HLerror
			return last                                           // HLerror
		} // HLerror

		return last // HLerror
	} // HLerror

	return fisc, errFn // HLreturn
}

// END2 OMIT
