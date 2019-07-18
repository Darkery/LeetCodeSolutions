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

func hasCycle(head *ListNode) bool {
	record := make(map[*ListNode]byte)
	temp := head

	for temp != nil {
		_, isPresent := record[temp]
		if isPresent {
			return true
		}
		record[temp] = 0
		temp = temp.Next
	}
	return false
}

func main() {

	head := initList(5)

	printLinkedList(head)
	head.Next.Next = head

	println("----------")

	println(hasCycle(head))
}