package work

import "testing"

func TestAdd(t *testing.T) {
	d := []struct {
		augend int
		addend int
		sum    int
	}{
		{1, 2, 3},
		{4, 5, 9},
		{100, 222, 322},
	}

	for _, v := range d {
		got := add(v.augend, v.addend)
		if v.sum != got {
			t.Errorf("want %d, got %d", v.sum, got)
		}
	}
}
