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

func reverseList(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	curr := head
	var result *ListNode = nil

	for curr != nil {
		curr.Next, result, curr = result, curr, curr.Next
	}

	return result
}

func main() {

	head := initList(5)

	printLinkedList(head)

	println("----------")

	printLinkedList(reverseList(head))
}