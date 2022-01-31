package algo3

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMegeK(t *testing.T) {

	tail1 := &ListNode{5, nil}
	prev1 := &ListNode{4, tail1}
	head1 := &ListNode{1, prev1}

	tail2 := &ListNode{4, nil}
	prev2 := &ListNode{3, tail2}
	head2 := &ListNode{1, prev2}

	tail3 := &ListNode{6, nil}
	head3 := &ListNode{2, tail3}
	testlist := []*ListNode{head1, head2, head3}

	merged := mergeKLists(testlist)

	expectedtail := &ListNode{6, nil}
	expectedprev1 := &ListNode{5, expectedtail}
	expectedprev2 := &ListNode{4, expectedprev1}
	expectedprev3 := &ListNode{4, expectedprev2}
	expectedprev4 := &ListNode{3, expectedprev3}
	expectedprev5 := &ListNode{2, expectedprev4}
	expectedprev6 := &ListNode{1, expectedprev5}
	expectedhead := &ListNode{1, expectedprev6}

	curr := merged
	for curr != nil {
		fmt.Println(*curr)
		curr = curr.Next
	}
	if !reflect.DeepEqual(merged, expectedhead) {
		t.Errorf("Expected List is not same as actual List")
	}

}
