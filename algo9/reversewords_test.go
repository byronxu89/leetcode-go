package algo9

import "testing"

var reversewordstests = []struct {
	s   string
	exp string
}{
	{
		"  hello world  ", "world hello",
	},
	{
		"the sky is blue", "blue is sky the",
	}, {
		"a good   example", "example good a",
	}, {
		"  Bob    Loves  Alice   ", "Alice Loves Bob",
	},
}

func TestReverseWords(t *testing.T) {
	for _, e := range reversewordstests {
		res := reverseWords(e.s)
		if res != e.exp {

			t.Errorf("reverseWords(%s) = %s, expected %s",
				e.s, res, e.exp)
		}
	}
}
