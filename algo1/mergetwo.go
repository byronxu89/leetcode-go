package algo

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	cur := head
	for list1 != nil || list2 != nil {
		if list1 != nil && list2 != nil {
			if list1.Val < list2.Val {
				cur.Next = list1
				list1 = list1.Next
			} else {
				cur.Next = list2
				list2 = list2.Next
			}
			cur = cur.Next
		} else if list1 != nil {
			cur.Next = list1
			break
		} else {
			cur.Next = list2
			break
		}
	}
	return head.Next

}

var MergeTwoLists = mergeTwoLists
