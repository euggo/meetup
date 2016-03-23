// START1 OMIT
package rpctime

import (
	"errors"
	"sync"
	"time"

	// importing externally defined "internal" packages not permitted at compile time
	"github.com/daved/rpctime/internal/zones"
)

type RPC struct {
	mu   sync.Mutex
	reqs uint64
}

func NewRPC() *RPC {
	return &RPC{}
}

// END1 OMIT

// START2 OMIT
func (r *RPC) Time(zone string, curTime *time.Time) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.reqs++

	loc, err := zones.LocByZone(zone)
	if err != nil {
		return errors.New("zone not found")
	}

	*curTime = time.Now().In(loc)
	return nil
}

func (r *RPC) Stats(_ bool, reqs *uint64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	*reqs = r.reqs
	return nil
}

// END2 OMIT
