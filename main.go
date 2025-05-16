package main

import (
	"fmt"
	"interview/algorithms/linked_list"
)

func main() {

	head := &linked_list.ListNodeMiddleNode{Val: 1}
	head.Next = &linked_list.ListNodeMiddleNode{Val: 2}
	head.Next.Next = &linked_list.ListNodeMiddleNode{Val: 3}
	head.Next.Next.Next = &linked_list.ListNodeMiddleNode{Val: 4}
	head.Next.Next.Next.Next = &linked_list.ListNodeMiddleNode{Val: 5}

	middle1 := linked_list.MiddleNode(head)
	fmt.Print("Original list: ", middle1.Val)

}
