package algo5

import "testing"

func TestEating(t *testing.T) {
	piles := []int{30, 11, 23, 4, 20}
	H := 6
	if minEatingSpeed(piles, H) != 23 {
		t.Errorf("#1 The expected eating speed  is not same as the actual speed")
	}
}
