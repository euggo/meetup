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
	rpc.Register(rpctime.NewRPC()) // HL

	l, err := net.Listen("tcp", ":19876") // HL
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	rpc.Accept(l) // HL
}

// END1 OMIT
