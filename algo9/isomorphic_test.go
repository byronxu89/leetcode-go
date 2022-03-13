package algo9

import "testing"

var tests = []struct {
	s   string
	t   string
	exp bool
}{
	{
		"egg", "add", true,
	}, {
		"abdc", "abab", false,
	},
	{
		"foo", "bar", false,
	}, {
		"paper", "title", true,
	},
}

func TestIsomorphic(t *testing.T) {
	for _, e := range tests {
		res := isIsomorphic(e.s, e.t)
		if res != e.exp {
			t.Errorf("isIsomorphic(%v, %s) = %t, expected %t",
				e.s, e.t, res, e.exp)
		}
	}
}
