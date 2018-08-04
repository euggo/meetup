package main

import (
	"github.com/euggo/meetup/talks/strings_and_things/a_unit_test/caselib"
	"github.com/euggo/meetup/talks/strings_and_things/runtest"
)

func main() {
	GoTest := func() { runtest.Run(caselib.TestToWaveCase) }
	//START1 OMIT
	GoTest() // helper function for demo use
	//END1 OMIT
}
