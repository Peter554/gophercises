package lib

import "testing"

func TestCaeserCipher(t *testing.T) {
	cases := []struct {
		s string
		k int
		w string
	}{
		{s: "middle-Outz", k: 2, w: "okffng-Qwvb"},
	}

	for _, c := range cases {
		got := CaeserCipher(c.s, c.k)
		if got != c.w {
			t.Errorf("CaeserCipher(\"%s\", %d) got \"%s\", wanted \"%s\"", c.s, c.k, got, c.w)
		}
	}
}
