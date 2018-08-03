package caselib

import "testing"

func TestToWaveCase(t *testing.T) {
	tests := []struct {
		s, want string
	}{
		{"in a fantastic place.", "In A fAnTaStIc PlAcE."},
		{"å være midt i smørøyet.", "Å vÆrE mIdT i SmØrØyEt."},
	}

	for _, tt := range tests {
		got := ToWaveCase(tt.s)
		if got != tt.want {
			t.Errorf("got %q, want %q", got, tt.want)
		}
	}
}
