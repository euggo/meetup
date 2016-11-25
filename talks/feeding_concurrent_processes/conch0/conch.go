package main

import (
	"errors"
	"path/filepath"
	"sync"
	"time"
)

// conch holds a file info group, and done channel.
//START1 OMIT
type conch struct {
	fig *fileInfoGroup
	d   chan struct{}
}

//END1 OMIT

// newConch returns a pointer to a new conch with channel setup.
//START2 OMIT
func newConch(fig *fileInfoGroup) *conch {
	return &conch{
		fig: fig,
		d:   make(chan struct{}),
	}
}

//END2 OMIT

// done returns the done channel.
//START3 OMIT
func (c *conch) done() chan struct{} {
	return c.d
}

//END3 OMIT

// produce is a generator that produces paths for processing. If canceled, an
// error is produced.
//START5 OMIT
func (c *conch) produce() (<-chan string, <-chan error) {
	paths := make(chan string)
	errs := make(chan error, 1)

	go func() {
		defer close(paths)

		for _, v := range c.fig.fsi {
			select {
			case paths <- filepath.Join(c.fig.dir, v.Name()):
			case <-c.d:
				errs <- errors.New("canceled")
				return
			}
		}
	}()

	return paths, errs
}

//END5 OMIT

// digest processes the file located at the currently provided path, and sends
// out a new result.
//START7 OMIT
func (c *conch) digest(paths <-chan string, fos chan<- *fileOutput) {
	for p := range paths {
		fo := newFileOutput(p)

		if slow {
			select {
			case <-time.After(time.Second):
			case <-c.d:
				return
			}
		}

		select {
		case fos <- fo:
		case <-c.d:
			return
		}
	}
}

//END7 OMIT

// consume sets-up digest goroutines according to width. Each digest goroutine
// waits for data from the paths channel and the entire function collapses when
// completed.
//START6 OMIT
func (c *conch) consume(paths <-chan string) <-chan *fileOutput {
	fos := make(chan *fileOutput)

	go func() {
		var wg sync.WaitGroup
		wg.Add(width)

		for i := 0; i < width; i++ {
			go func() {
				c.digest(paths, fos)
				wg.Done()
			}()
		}

		wg.Wait()
		close(fos)
	}()

	return fos
}

//END6 OMIT

// run calls the path generator, calls the consume func using the returned
// paths channel, and returns the outs and errs channels.
//START4 OMIT
func (c *conch) run() (<-chan *fileOutput, <-chan error) {
	paths, errs := c.produce()
	fos := c.consume(paths)

	return fos, errs
}

//END4 OMIT
