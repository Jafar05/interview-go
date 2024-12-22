package main

import (
	"fmt"
	"time"
)

func main() {
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}

	mergeSl := make([]int, len(slice1)+len(slice2))

	for i, char := range slice1 {
		mergeSl[i] = char
	}

	for i, char := range slice2 {
		mergeSl[len(slice1)+i] = char
	}
	time.After(1)

	fmt.Println(mergeSl)

}
