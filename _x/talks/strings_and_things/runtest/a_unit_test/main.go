package main

import (
	"github.com/euggo/meetup/_x/talks/strings_and_things/a_unit_test/caselib"
	"github.com/euggo/meetup/_x/talks/strings_and_things/runtest"
)

func main() {
	GoTest := func() { runtest.Run(caselib.TestToWaveCase) }
	//START1 OMIT
	GoTest() // This is a helper function for demonstration only.
	//END1 OMIT
}
