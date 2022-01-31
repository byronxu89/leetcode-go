package algo

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	actualresult := PlusOne([]int{9, 9, 9})
	expectedresult := []int{1, 0, 0, 0}
	if !reflect.DeepEqual(actualresult, expectedresult) {
		t.Errorf("Expected result is not same as the actual result")
	}
}
