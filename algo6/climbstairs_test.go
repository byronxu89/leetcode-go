package algo6

import "testing"

func TestClimb(t *testing.T) {
	n := 30
	if climbStairs(n) != 1346269 {
		t.Errorf("#1 The expected ways  is not same as actual ways of climbing stairs")
	}
}
