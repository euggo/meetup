// START1 OMIT
package main

import (
	"fmt"
	"net"
	"net/rpc"
	"os"

	"github.com/daved/rpctime"
)

func main() {
	rpc.Register(rpctime.NewRPC())

	l, err := net.Listen("tcp", ":19876")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	rpc.Accept(l)
}

// END1 OMIT
