package main

import (
	"sync"
)

// BGN3 OMIT
func produce(paths []string) <-chan string { // HLargs
	psc := make(chan string) // HLchan

	// BGN31 OMIT
	go func() { // feed paths chan, close chan when complete // HLgoroutine
		defer close(psc) // HLclose

		for _, p := range paths { // HLiterate
			psc <- p // HLselect
		} // HLiterate
	}() // HLgoroutine
	// END31 OMIT

	return psc // HLreturn
} // HLargs

// END3 OMIT

// BGN5 OMIT
func digest(fisc chan<- *fileInfo, psc <-chan string) { // HLargs
	for p := range psc { // HLreceive
		fisc <- newFileInfo(p) // HLprocess
	} // HLreceive
} // HLargs

// END5 OMIT

// BGN4 OMIT
func consume(width int, psc <-chan string) <-chan *fileInfo { // HLargs
	fisc := make(chan *fileInfo) // HLchan

	// BGN41 OMIT
	go func() { // create consumers, close fileInfo channel when complete // HLgoroutine
		defer close(fisc) // HLclose

		var wg sync.WaitGroup // HLwg
		wg.Add(width)         // HLwg

		// BGN411 OMIT
		for i := 0; i < width; i++ { // HLwidth
			go func() { // run in goroutine to bind with wg.Done call // HLgogoroutine
				digest(fisc, psc) // HLdigest
				wg.Done()         // HLwg
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
func fileInfos(width int, paths []string) <-chan *fileInfo { // HLargs
	psc := produce(paths)       // HLproduce
	fisc := consume(width, psc) // HLconsume

	return fisc // HLreturn
} // HLargs

// END2 OMIT
