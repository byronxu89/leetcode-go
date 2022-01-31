package algo2

import (
	"fmt"
	"testing"
)

func TestShortestSubarray(t *testing.T) {
	nums := []int{1, 2, 2, 3, 1, 4, 2}
	actualnum := findShortestSubArray(nums)
	exptectnum := 6
	fmt.Println(actualnum)
	if exptectnum != actualnum {
		t.Errorf("Expected shortest sub array result is not same as the actual result")
	}
}
