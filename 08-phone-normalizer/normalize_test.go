package main

import "testing"

func TestNormalize(t *testing.T) {
	cases := []struct {
		s string
		w string
	}{
		{s: "1 2 3 4 5", w: "12345"},
		{s: "123_45", w: "12345"},
		{s: "12-34-5", w: "12345"},
	}

	for _, c := range cases {
		got, e := normalize(c.s)
		if e != nil {
			t.Error(e)
		}
		if got != c.w {
			t.Errorf("normalize(%s) got \"%s\", wanted \"%s\"", c.s, got, c.w)
		}
	}
}
