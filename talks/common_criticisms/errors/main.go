package main

import (
	"errors"
	"log"
	"time"
)

type murphy struct{}

func newMurphy() *murphy {
	return &murphy{}
}

func (m *murphy) law() (string, error) {
	s := time.Now().Second()
	if s%2 == 0 {
		return "", errWillGoWrong
	}
	if s%3 == 0 || s%5 == 0 || s%7 == 0 {
		return "", errCanGoWrong
	}

	return "inconceivable!", nil
}

var (
	errWillGoWrong = errors.New("this is fine")
	errCanGoWrong  = errors.New("we knew it would")
)

// START OMIT
func main() {
	m := newMurphy()

	l, err := m.law()
	if err != nil {
		if err == errWillGoWrong {
			log.Println("will error:", err)
			log.Fatalln("mod 2 / l len is", len(l))
		}

		if err == errCanGoWrong {
			log.Println("mod 3,5,7 / l len is", len(l))
			log.Fatalln("can error:", err)
		}
	}

	log.Println(l)
}

// END OMIT
