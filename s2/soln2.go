/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func digitCount(number int) int {
	return len(strconv.Itoa(number))
}

// Converts number to an int, which results in a length limitation - numbers longer than int64 allows will overflow
func addTwoNumbers_short(l1 *ListNode, l2 *ListNode) *ListNode {
	// Convert linked lists to numbers
	var powerOf10 float64 = 0.0
	var num1, num2 int
	for {
		exponent := int(math.Pow(10, powerOf10))
		if l1 != nil {
			num1 += l1.Val * exponent
			l1 = l1.Next
		}
		if l2 != nil {
			num2 += l2.Val * exponent
			l2 = l2.Next
		}
		if l1 == nil && l2 == nil {
			break
		}
		powerOf10 += 1.0
	}

	// Sum the numbers
	var numOut int = num1 + num2

	// Convert the sum to a linked list, starting with the senior digit
	currentListNode := &ListNode{0, nil}
	digitCount := digitCount(numOut)
	powerOf10 = float64(digitCount) - 1
	for {
		exponent := int(math.Pow(10, powerOf10))
		digit := numOut / exponent
		currentListNode.Val = digit
		numOut -= digit * exponent
		powerOf10 -= 1.0
		if powerOf10 < 0.0 {
			break
		}
		currentListNode = &ListNode{0, currentListNode}
	}

	return currentListNode
}
