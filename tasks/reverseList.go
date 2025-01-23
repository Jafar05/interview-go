package tasks

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func reverseList(head *Node) *Node {
	var prev *Node

	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}

func printList(head *Node) {
	for head != nil {
		fmt.Printf("%d -> ", head.Value)
		head = head.Next
	}
	fmt.Println("nil")
}

func PrintReverse() {
	// Создаём односвязный список: 1 -> 2 -> 3 -> 4 -> nil
	head := &Node{Value: 1}
	head.Next = &Node{Value: 2}
	head.Next.Next = &Node{Value: 3}
	head.Next.Next.Next = &Node{Value: 4}

	fmt.Println("Исходный список:")
	printList(head)

	// Разворачиваем список
	reversed := reverseList(head)

	fmt.Println("Развёрнутый список:")
	printList(reversed)
}
