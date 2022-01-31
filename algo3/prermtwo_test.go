package algo3

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPerm(t *testing.T) {
	testslice := []int{1, 1, 2, 2, 3}
	realval := permuteUnique(testslice)
	expectedval := [][]int{{1, 1, 2, 2, 3}, {1, 1, 2, 3, 2}, {1, 1, 3, 2, 2},
		{1, 2, 1, 2, 3}, {1, 2, 1, 3, 2}, {1, 2, 2, 1, 3}, {1, 2, 2, 3, 1}, {1, 2, 3, 1, 2},
		{1, 2, 3, 2, 1}, {1, 3, 1, 2, 2}, {1, 3, 2, 1, 2}, {1, 3, 2, 2, 1}, {2, 1, 1, 2, 3},
		{2, 1, 1, 3, 2}, {2, 1, 2, 1, 3}, {2, 1, 2, 3, 1}, {2, 1, 3, 1, 2}, {2, 1, 3, 2, 1},
		{2, 2, 1, 1, 3}, {2, 2, 1, 3, 1}, {2, 2, 3, 1, 1}, {2, 3, 1, 1, 2}, {2, 3, 1, 2, 1},
		{2, 3, 2, 1, 1}, {3, 1, 1, 2, 2}, {3, 1, 2, 1, 2}, {3, 1, 2, 2, 1}, {3, 2, 1, 1, 2},
		{3, 2, 1, 2, 1}, {3, 2, 2, 1, 1}}
	fmt.Println(realval)
	if !reflect.DeepEqual(expectedval, realval) {
		t.Errorf("Expected permutation is not same as actual permutation")
	}

}
