package main

import (
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

// BGN OMIT
func main() {
	lambda.Start(f)
}

type request struct {
	Name string `json:"name"`
}

type response struct {
	Msg  string `json:"msg"`
	Time string `json:"time"`
}

func f(req request) (response, error) {
	msg := fmt.Sprintf("Hello from AWS, %s. Have a great day.", req.Name)

	resp := response{
		Msg:  msg,
		Time: time.Now().String(),
	}

	return resp, nil
}

// END OMIT
