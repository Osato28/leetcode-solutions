/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// First try: misinterpreted the problem description, solved the wrong kind of problem
func addTwoNumbers_mistake(l1 *ListNode, l2 *ListNode) *ListNode {
	curr1 := l1
	curr2 := l2

	startOut := &ListNode{0, nil}
	var currOut *ListNode = startOut
	for {
		if curr1 != nil {
			currOut.Val += curr1.Val
			curr1 = curr1.Next
		}
		if curr2 != nil {
			currOut.Val += curr2.Val
			curr2 = curr2.Next
		}
		if curr1 == nil && curr2 == nil {
			break
		}
		currOut.Next = &ListNode{0, nil}
		currOut = currOut.Next
	}
	return startOut
}