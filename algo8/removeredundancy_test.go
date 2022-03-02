package algo8

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMinimal(t *testing.T) {
	edges := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}}
	redundantedge := findRedundantConnection(edges)
	fmt.Println(redundantedge)
	if !reflect.DeepEqual(redundantedge, []int{1, 4}) {
		t.Errorf("#1 The computed redundant edge is incorrect")
	}
}
