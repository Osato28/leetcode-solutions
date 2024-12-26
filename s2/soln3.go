
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// Adds digits in place
	var carryThe1 bool = false
	var node *ListNode = &ListNode{0, nil}
	firstNode := node
	for {
		if l1 != nil {
			node.Val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			node.Val += l2.Val
			l2 = l2.Next
		}

		if node.Val >= 10 {
			node.Val = node.Val % 10
			carryThe1 = true
		}

		var initValue int
		if carryThe1 {
			initValue = 1
		} else {
			initValue = 0
		}
		carryThe1 = false

		if l1 == nil && l2 == nil && initValue == 0 {
			break
		}

		nextNode := &ListNode{initValue, nil}
		node.Next = nextNode
		node = nextNode
	}
	return firstNode
}