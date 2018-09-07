package main

import "github.com/euggo/meetup/talks/going_faast/client"

func main() {
	// BGN OMIT
	client.PrintCall(
		"https://us-central1-experimental-214818.cloudfunctions.net/faasfun_longechox",
		`"This may not work."`,
	)
	// END OMIT
}
