package algo6

import "testing"

func TestMinimal(t *testing.T) {
	triangle := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	miniumsum := minimumTotal(triangle)
	if miniumsum != 11 {
		t.Errorf("#1 The minnum of tht total path sum is unequal to the actual sum")
	}
}
