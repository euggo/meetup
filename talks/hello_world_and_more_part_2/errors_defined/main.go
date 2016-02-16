package main

import (
	"errors"
	"log"
	"time"
)

var (
	m = &murphy{}
)

type murphy struct{}

// START OMIT
func (m *murphy) law() (bool, error) {
	if time.Now().Second()%2 == 0 {
		return true, errEven
	}
	return true, errOdd
}

var (
	errEven = errors.New("i can even")
	errOdd  = errors.New("something is odd here")
)

func main() {
	if b, err := m.law(); err != nil {
		if err == errEven {
			log.Println(b, err)
		}

		if err == errOdd {
			log.Println(err, b)
		}
	}
}

// END OMIT
