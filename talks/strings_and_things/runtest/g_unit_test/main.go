package main

import (
	"github.com/euggo/meetup/_x/talks/strings_and_things/g_unit_test/caselib"
	"github.com/euggo/meetup/_x/talks/strings_and_things/runtest"
)

func main() {
	GoTest := func() { runtest.Run(caselib.TestToWaveCase) }
	//START1 OMIT
	GoTest() // helper function for demo use
	//END1 OMIT
}
