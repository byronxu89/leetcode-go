package algo4

import (
	"reflect"
	"testing"
)

// test function
func TestSolve(t *testing.T) {
	/*
		actualString := ReturnGeeks()
		expectedString := "geeks"
		if actualString != expectedString {
			t.Errorf("Expected String(%s) is not same as"+
				" actual string (%s)", expectedString, actualString)
		}
	*/

	actualboard := [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}}

	solve(actualboard)

	expectedboard := [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'O', 'X', 'X'}}

	if !reflect.DeepEqual(actualboard, expectedboard) {
		t.Errorf("Expected board is not same as actual board ")
	}

}
