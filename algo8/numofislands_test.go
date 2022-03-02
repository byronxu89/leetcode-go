package algo8

import (
	"testing"
)

func TestNumOfIslands(t *testing.T) {

	grids := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	if numIslands(grids) != 3 {
		t.Errorf("#1 The computed num of islands is incorrect")
	}

}
