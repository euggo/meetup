package work

import "testing"

func TestAdd(t *testing.T) {
	cases := []struct {
		augend int
		addend int
		sum    int
	}{
		{1, 2, 3},
		{4, 5, 9},
		{100, 222, 322},
	}

	for _, v := range cases {
		got := add(v.augend, v.addend)
		if got != v.sum {
			t.Errorf("got %d, want %d", got, v.sum)
		}
	}
}
