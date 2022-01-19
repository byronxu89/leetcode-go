package main

import (
	. "app/algo3"
	"fmt"
)

func main() {
	tail1 := &ListNode{5, nil}
	prev1 := &ListNode{4, tail1}
	head1 := &ListNode{1, prev1}

	tail2 := &ListNode{4, nil}
	prev2 := &ListNode{3, tail2}
	head2 := &ListNode{1, prev2}

	tail3 := &ListNode{6, nil}
	head3 := &ListNode{2, tail3}
	testlist := []*ListNode{head1, head2, head3}

	merged := MergeKLists(testlist)
	curr := merged
	for curr != nil {
		fmt.Println(*curr)
		curr = curr.Next
	}

	testslice := []int{1, 1, 2, 2, 3}
	ans := PermuteUnique(testslice)
	fmt.Println(ans)

}
