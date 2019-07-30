package Add_Two_Numbers

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var head *ListNode = new(ListNode)
	cur := head
	carry := 0

	for l1 != nil || l2 != nil {
		v1 := 0
		v2 := 0
		if l1 != nil {
			v1 = l1.Val
		}

		if l2 != nil {
			v2 = l2.Val
		}

		sum := carry + v1 + v2
		carry = sum / 10

		cur.Next = new(ListNode)
		cur = cur.Next
		cur.Val = sum % 10

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	if carry > 0 {
		cur.Next = new(ListNode)
		cur = cur.Next
		cur.Val = carry
	}

	return head.Next
}
