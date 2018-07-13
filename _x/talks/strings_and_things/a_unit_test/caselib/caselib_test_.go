package caselib

import "testing"

//START1 OMIT
func TestToWaveCase(t *testing.T) {
	ds := []struct {
		s, want string
	}{
		{"in a fantastic place.", "In A fAnTaStIc PlAcE."},
		{"å være midt i smørøyet.", "Å vÆrE mIdT i SmØrØyEt."},
	}
	for _, d := range ds {
		got := ToWaveCase(d.s)
		if got != d.want {
			t.Errorf("got %q, want %q", got, d.want)
		}
	}
}

//END1 OMIT
