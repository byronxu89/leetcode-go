package algo5

import "testing"

func TestFinish(t *testing.T) {
	weights := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	days := 5
	if shipWithinDays(weights, days) != 15 {
		t.Errorf("#1 The expected ways  is not same as actual ways of climbing stairs")
	}
}
