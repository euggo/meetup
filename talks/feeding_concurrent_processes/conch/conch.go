package main

import (
	"errors"
	"sync"
)

// BGN1 OMIT
type conch struct {
	d chan struct{}
}

func newConch() *conch {
	return &conch{
		d: make(chan struct{}),
	}
}

func (c *conch) done() chan struct{} {
	return c.d
}

// END1 OMIT

// BGN3 OMIT
func (c *conch) produce(paths []string) (<-chan string, <-chan error) {
	psc := make(chan string)
	ec := make(chan error, 1)

	go func() { // feed paths chan, close chan when complete
		defer close(psc)

		for _, p := range paths {
			select {
			case psc <- p:
			case <-c.d:
				ec <- errors.New("canceled")
				return
			}
		}
	}()

	return psc, ec
}

// END3 OMIT

// BGN5 OMIT
func (c *conch) digest(fisc chan<- *fileInfo, psc <-chan string) {
	for p := range psc {
		fi := newFileInfo(p)
		select {
		case fisc <- fi:
		case <-c.d:
			return
		}
	}
}

// END5 OMIT

// BGN4 OMIT
func (c *conch) consume(width int, psc <-chan string) <-chan *fileInfo {
	fisc := make(chan *fileInfo)

	go func() { // create consumers, close fileInfo channel when complete
		var wg sync.WaitGroup
		wg.Add(width)

		for i := 0; i < width; i++ {
			go func() { // run consumer in goroutine to bind with wg.Done call
				c.digest(fisc, psc)
				wg.Done()
			}()
		}

		wg.Wait()
		close(fisc)
	}()

	return fisc
}

// END4 OMIT

// BGN2 OMIT
func (c *conch) run(width int, paths []string) (<-chan *fileInfo, func() error) {
	psc, ec := c.produce(paths)   // HLproduce
	fisc := c.consume(width, psc) // HLconsume

	errFn := func() error { // HLerror
		select { // HLerror
		case err := <-ec: // HLerror
			return err // HLerror
		default: // HLerror
			return nil // HLerror
		} // HLerror
	} // HLerror

	return fisc, errFn
}

// END2 OMIT
