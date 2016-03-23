// START1 OMIT
package main

import (
	"net/http" // OMIT
	"time"     // OMIT
	// OMIT
	"github.com/daved/rpctime" // OMIT
	// "net/http", "time", "github.com/daved/rpctime"
)

type node struct {
	timeServer  *rpctime.Client
	multiplexer *http.ServeMux
}

func newNode(rpcPort string) (*node, error) {
	ts, err := rpctime.NewClient(rpcPort, time.Second*6)
	if err != nil {
		return nil, err
	}

	n := &node{}
	n.timeServer = ts
	n.multiplexer = n.mux()

	return n, nil
}

// END1 OMIT

// START2 OMIT
func (n *node) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	n.multiplexer.ServeHTTP(w, r)
}

func (n *node) mux() *http.ServeMux {
	m := http.NewServeMux()

	m.HandleFunc("/api/local", n.localHandler)
	m.HandleFunc("/api/remote/", n.remoteHandler)
	m.HandleFunc("/api/stats", n.statsHandler)

	return m
}

// END2 OMIT
