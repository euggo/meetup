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
	//BGN3 OMIT
	for _, d := range ds {
		got := ToWaveCase(d.s)
		if got != d.want {
			t.Errorf("got %q, want %q", got, d.want)
		}
	}
	//END3 OMIT
}

//END1 OMIT
