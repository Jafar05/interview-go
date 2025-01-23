package main

import (
	"fmt"
	"interview-go/tasks"
)

func main() {
	nums := []int{1, 3, 5, 7, 9, 11, 13, 15}

	// Тесты
	fmt.Println(tasks.BinarySearch(nums, 7))  // Ожидается: 3
	fmt.Println(tasks.BinarySearch(nums, 4))  // Ожидается: -1
	fmt.Println(tasks.BinarySearch(nums, 13)) // Ожидается: 6
}
