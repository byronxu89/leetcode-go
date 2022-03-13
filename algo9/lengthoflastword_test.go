package algo9

import "testing"

var lengthtests = []struct {
	s   string
	exp int
}{
	{
		"Hello World", 5,
	},
	{
		"   fly me   to   the moon  ", 4,
	}, {
		"luffy is still joyboy", 6,
	},
}

func TestILengthOfLastWord(t *testing.T) {
	for _, e := range lengthtests {
		res := lengthOfLastWord(e.s)
		if res != e.exp {
			t.Errorf("lengthOfLastWord(%s) = %d, expected %d",
				e.s, res, e.exp)
		}
	}
}
