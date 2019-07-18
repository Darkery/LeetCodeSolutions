package main

type ListNode struct {
	Val int
	Next *ListNode
}

func initList(size int) *ListNode{

	i := 1
	node := new(ListNode)
	head := node

	for i <= size {
		node.Val = i
		if i == size {
			break
		}
		node.Next = new(ListNode)
		node = node.Next
		i++
	}

	return head
}

func printLinkedList(node *ListNode)  {
	temp := node
	for temp != nil{
		println("temp.Val = ", temp.Val)
		temp = temp.Next
	}
}

func reverseList(head *ListNode, len int) (h *ListNode, t *ListNode) {

	if head == nil {
		return nil, nil
	}

	curr := head
	var result *ListNode = nil

	for i := 0; i < len; i++ {
		println("i = ", i)
		curr.Next, result, curr = result, curr, curr.Next
	}

	head.Next = curr
	return result, head
}

func isGroupMemLessThanK(head *ListNode, k int) bool {
	temp := head
	for i := 0; i < k; i++ {
		if temp == nil{
			return true
		}
		temp = temp.Next
	}
	return false
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if isGroupMemLessThanK(head, k) || k <= 1 {
		return head
	}

	cur := new(ListNode)
	cur.Next = head
	result := new(ListNode)

	firstTime := true
	for !isGroupMemLessThanK(cur.Next, k) {
		groupHead, groupTail := reverseList(cur.Next, k)
		cur.Next = groupHead
		cur = groupTail
		if firstTime {
			result = groupHead
			firstTime = false
		}
	}

	return result
}


func main() {

	head := initList(7)

	printLinkedList(head)

	println("----------")

	printLinkedList(reverseKGroup(head, 3))
}