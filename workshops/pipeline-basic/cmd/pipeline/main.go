package main

import (
	"context"
	"flag"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/codemodus/sigmon/v2"
)

func main() {
	var (
		slow  bool
		width = 2
	)

	flag.BoolVar(&slow, "slow", slow, "visualize/emphasize concurrency")
	flag.IntVar(&width, "width", width, "concurrency width")
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sm := sigmon.New(func(s *sigmon.State) {
		fmt.Println(s.Signal())
		cancel()
	})
	sm.Start()

	infos := orchestrate(ctx, width, slow, data())

	for i := range infos {
		fmt.Println(i)
	}
}

// Info ...
type Info struct {
	routine int
	s       string
}

func newInfo(n int, s string) *Info {
	return &Info{
		routine: n,
		s:       strings.ToUpper(s),
	}
}

func orchestrate(ctx context.Context, width int, slow bool, d []string) <-chan *Info {
	ssc := produce(ctx, d)
	infos := consume(width, slow, ssc)

	return infos
}

func produce(ctx context.Context, d []string) <-chan string {
	ssc := make(chan string)

	go func() {
		defer close(ssc)

		for _, s := range d {
			select {
			case ssc <- s:
			case <-ctx.Done():
				fmt.Println("done!!!")
				return
			}
		}
	}()

	return ssc
}

func consume(width int, slow bool, ssc <-chan string) <-chan *Info {
	infos := make(chan *Info)

	go func() {
		defer func() {
			close(infos)
		}()

		var wg sync.WaitGroup
		wg.Add(width)

		for iter := 0; iter < width; iter++ {
			go func(n int) {
				defer wg.Done()

				for s := range ssc {
					if slow {
						time.Sleep(time.Second * 2)
					}
					infos <- newInfo(n, s)
				}
			}(iter)
		}

		wg.Wait()
	}()

	return infos
}
