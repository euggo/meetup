package caselib

import "testing"

//START1 OMIT
func TestToWaveCase(t *testing.T) {
	//BGN2 OMIT
	tests := []struct {
		s, want string
	}{
		{"in a fantastic place.", "In A fAnTaStIc PlAcE."},
		{"å være midt i smørøyet.", "Å vÆrE mIdT i SmØrØyEt."},
	}
	//END2 OMIT
	//BGN4 OMIT
	for _, tt := range tests {
		got := ToWaveCase(tt.s)
		if got != tt.want {
			t.Errorf("got %q, want %q", got, tt.want)
		}
	}
	//END4 OMIT
}

//END1 OMIT

// Junk ...
func Junk() {
	// should match "tests" in TestToWaveCase OMIT
	//BGN3 OMIT
	tests := []struct { // HL
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
	_ = tests // OMIT
}
