package addtwonumbers

func main() {

}

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// Get position based on decimal count [1,2,3] -> 1 num -> 21 -> 2 num -> 321 -> nums
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	resHead := &ListNode{Val: 0}
	curr := resHead

	// Initialize values for mem efficiency
	carry := 0
	ires := 0

	// while calculation still needs to be done keep looping
	for l1 != nil || l2 != nil || carry != 0 {
		l1v := 0
		l2v := 0

		// initialize the values that need to be added
		if l1 != nil {
			l1v = l1.Val
		}

		if l2 != nil {
			l2v = l2.Val
		}

		// The summation of the values
		ires = l1v + l2v + carry

		// Calculate the value to carry
		carry = ires / 10

		// Calculate the remainder and create a node
		curr.Next = &ListNode{Val: ires % 10}
		curr = curr.Next

		// Update node values
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	// Return the head.next because the head contains a dummy value
	return resHead.Next
}
