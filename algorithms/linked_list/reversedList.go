package linked_list

type ListReverseNode struct {
	Val  int
	Next *ListReverseNode
}

func ReversedList(head *ListReverseNode) *ListReverseNode {
	prev := &ListReverseNode{}
	curr := head
	for curr != nil {
		tmp := curr
		curr = curr.Next
		tmp.Next = prev
		prev = tmp
	}
	return prev
}
