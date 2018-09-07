package main

import "github.com/euggo/meetup/talks/going_faast/client"

func main() {
	// BGN OMIT
	client.PrintCall(
		"https://lambda.aws.euggo.org/faasfun/longecho",
		`"Will it work?"`,
	)
	// END OMIT
}
