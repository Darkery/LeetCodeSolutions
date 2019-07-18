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

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	cur := new(ListNode)
	cur.Next = head
	result := head.Next

	for cur.Next != nil && cur.Next.Next != nil {
		cur.Next, cur.Next.Next, cur.Next.Next.Next = cur.Next.Next, cur.Next.Next.Next, cur.Next
		cur = cur.Next.Next
	}
	return result
}

func main() {

	head := initList(5)

	printLinkedList(head)

	println("----------")

	printLinkedList(swapPairs(head))
}