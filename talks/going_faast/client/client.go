package client

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	ctype = "application/json"
)

// PrintCall ...
func PrintCall(url, payload string) {
	resp, err := http.Post(url, ctype, bytes.NewBufferString(payload))
	trip(err)

	bs, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close() // nolint
	trip(err)

	fmt.Printf("Code: %d\n", resp.StatusCode)
	fmt.Printf("Msg : %s\n", string(bs))
}

func trip(err error) {
	if err != nil {
		panic(err)
	}
}
