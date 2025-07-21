package linked_list

type ListNodeMiddleNodePalindrome struct {
	Val  int
	Next *ListNodeMiddleNodePalindrome
}

func IsPalindromeLinkedList(head *ListNodeMiddleNodePalindrome) bool {
	firstHalfEnd := middle(head)
	secondHalfStart := reverse(firstHalfEnd)

	p1 := head
	p2 := secondHalfStart

	for p1 != nil && p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1, p2 = p1.Next, p2.Next
	}
	return true
}

func middle(head *ListNodeMiddleNodePalindrome) *ListNodeMiddleNodePalindrome {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	return slow
}

func reverse(head *ListNodeMiddleNodePalindrome) *ListNodeMiddleNodePalindrome {
	curr := head
	var prev *ListNodeMiddleNodePalindrome
	for curr != nil {
		tmp := curr.Next
		curr.Next = prev
		prev = curr
		curr = tmp
	}
	return prev
}
