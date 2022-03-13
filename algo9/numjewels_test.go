package algo9

import "testing"

func TestNumJewels(t *testing.T) {

	res := numJewelsInStones("aA", "aAAbbbb")
	if res != 3 {
		t.Errorf("numJewelsInStones(%s, %s) = %d, expected %d",
			"aA", "aAAbbbb", res, 3)

	}
}
