// START1 OMIT
package rpctime

import (
	"net"     // OMIT
	"net/rpc" // OMIT
	"time"    // OMIT
	// "net", "net/rpc", "time"
)

type Client struct {
	conn *rpc.Client // HL
}

func NewClient(addr string, timeout time.Duration) (*Client, error) {
	conn, err := net.DialTimeout("tcp", addr, timeout) // HL
	if err != nil {
		return nil, err
	}

	c := &Client{
		conn: rpc.NewClient(conn), // HL
	}

	return c, nil
}

// END1 OMIT

// START2 OMIT
func (c *Client) Time(zone string) (string, error) {
	var curTime string

	err := c.conn.Call("RPC.Time", zone, &curTime) // HL

	return curTime, err
}

func (c *Client) Stats() (uint64, error) {
	var reqs uint64

	err := c.conn.Call("RPC.Stats", false, &reqs) // HL

	return reqs, err
}

// END2 OMIT
