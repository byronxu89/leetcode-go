package algo3

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	if n == 0 {
		return nil
	}
	if n == 1 {
		return lists[0]
	}
	previous, current := lists[0], lists[1]
	for i := 1; i < n; i++ {
		current = lists[i]
		previous = mergeTwoList(previous, current)
	}
	return previous
}
func mergeTwoList(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	current := head
	for l1 != nil || l2 != nil {
		if l1 == nil {
			current.Next = l2
			break
		}
		if l2 == nil {
			current.Next = l1
			break
		}
		if l1.Val < l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}
	return head.Next

}
