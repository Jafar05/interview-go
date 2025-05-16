package linked_list

type ListNodeMiddleNode struct {
	Val  int
	Next *ListNodeMiddleNode
}

func MiddleNode(head *ListNodeMiddleNode) *ListNodeMiddleNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}
