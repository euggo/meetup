package main

import (
	"errors"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(f)
}

// BGN OMIT
func f() error {
	time.Sleep(time.Millisecond * 3100)

	return errors.New("this should not be reached")
}

// END OMIT
