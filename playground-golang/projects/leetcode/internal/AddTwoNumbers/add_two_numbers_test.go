package addtwonumbers

import "testing"

func TestCase1(t *testing.T) {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}

	res := addTwoNumbers(l1, l2)

	if res == nil {
		
	}
}