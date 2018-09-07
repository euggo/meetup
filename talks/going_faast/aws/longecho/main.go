package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

var (
	errTooLate = errors.New("Sometimes, if you wait too long, it's too late. — Sarah Strohmeyer")
)

// BGN OMIT
func main() {
	lambda.Start(f)
}

func f(ctx context.Context, req *string) (string, error) {
	var toc <-chan time.Time
	if dl, ok := ctx.Deadline(); ok {
		toc = time.After(time.Until(dl.Add(-10 * time.Millisecond)))
	}

	select {
	case <-toc:
		return "", errTooLate
	case resp := <-randomlySlowMessage(*req):
		return resp, nil
	}
}

// END OMIT

func randomlySlowMessage(in string) <-chan string {
	c := make(chan string)

	go func() {
		if time.Now().UnixNano()%2 == 0 {
			time.Sleep(3001 * time.Millisecond) // should be just above default lambda timeout
		}

		c <- fmt.Sprintf("Your input: %s", in)
	}()

	return c
}
