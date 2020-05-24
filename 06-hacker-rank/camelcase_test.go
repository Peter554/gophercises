package lib

import "testing"

func TestCamelCase(t *testing.T) {
	cases := []struct {
		s string
		w int
	}{
		{s: "foo", w: 1},
		{s: "fooBarBaz", w: 3},
		{s: "", w: 0},
	}

	for _, c := range cases {
		got := CamelCase(c.s)
		if got != c.w {
			t.Errorf("%s got %d, wanted %d", c.s, got, c.w)
		}
	}
}
