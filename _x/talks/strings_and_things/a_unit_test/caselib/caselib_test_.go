package caselib

import "testing"

//START1 OMIT
func TestToWaveCase(t *testing.T) {
	//BGN2 OMIT
	ds := []struct {
		s, want string
	}{
		{"in a fantastic place.", "In A fAnTaStIc PlAcE."},
		{"å være midt i smørøyet.", "Å vÆrE mIdT i SmØrØyEt."},
	}
	//END2 OMIT
	//BGN4 OMIT
	for _, d := range ds {
		got := ToWaveCase(d.s)
		if got != d.want {
			t.Errorf("got %q, want %q", got, d.want)
		}
	}
	//END4 OMIT
}

//END1 OMIT

// Junk ...
func Junk() {
	// should match "ds" in TestToWaveCase OMIT
	//BGN3 OMIT
	ds := []struct { // HL
		// start of anonymous struct definition
		s    string // HL
		want string // HL
		// end of anonymous struct definition
	}{ // HL
		// start of slice literal
		// // struct literals
		{"in a fantastic place.", "In A fAnTaStIc PlAcE."},     // HL
		{"å være midt i smørøyet.", "Å vÆrE mIdT i SmØrØyEt."}, // HL
		// end of slice literal
	} // HL
	//END3 OMIT
	_ = ds // OMIT
}
