package main

import (
	. "app/algo1"
	"fmt"
)

func main() {

	result := PlusOne([]int{1, 2, 3})
	fmt.Println(result)

	node1 := ListNode{Val: 4, Next: nil}
	node2 := ListNode{Val: 2, Next: &node1}
	node3 := ListNode{Val: 1, Next: &node2}

	node4 := ListNode{Val: 4, Next: nil}
	node5 := ListNode{Val: 3, Next: &node4}
	node6 := ListNode{Val: 1, Next: &node5}

	head := MergeTwoLists(&node3, &node6)

	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}

}
