package main

import (
	"encoding/json" // OMIT
	"fmt"
	"net/http" // OMIT
	"os"
	"time" // OMIT

	"github.com/codemodus/parth"
	"github.com/daved/rpctime"
	// "encoding/json", "net/http", "time"
)

type timeJSON struct {
	Time time.Time `json:"time"`
}

func (n *node) localHandler(w http.ResponseWriter, r *http.Request) {
	t := &timeJSON{time.Now()}

	b, err := json.Marshal(t)
	if err != nil {
		http.Error(w, http.StatusText(500), 500) // use http.StatusInternalServerError
		return
	}

	w.Write([]byte(b))
}

// START1 OMIT
func (n *node) remoteHandler(w http.ResponseWriter, r *http.Request) {
	zoneID, err := parth.SubSegToString(r.URL.Path, "remote")
	if err != nil || zoneID == "" {
		zoneID = "GMT"
	}

	res, err := n.timeServer.Time(zoneID)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	t := &timeJSON{res}

	b, err := json.Marshal(t)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	w.Write(b)
}

// END1 OMIT

// START2 OMIT
type statsJSON struct {
	RPCCount uint64 `json:"rpc_count"`
}

func (n *node) statsHandler(w http.ResponseWriter, r *http.Request) {
	res, err := n.timeServer.Stats()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	s := &statsJSON{res}

	b, err := json.Marshal(s)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	w.Write(b)
}

// END2 OMIT

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

func main() {
	n, err := newNode(":19876")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	http.ListenAndServe(":29876", n)
}
