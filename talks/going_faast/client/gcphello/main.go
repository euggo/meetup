package main

import "github.com/euggo/meetup/talks/going_faast/client"

func main() {
	// BGN OMIT
	client.PrintCall(
		"https://us-central1-experimental-214818.cloudfunctions.net/faasfun_hellox",
		`{"name": "EUGGO"}`,
	)
	// END OMIT
}
