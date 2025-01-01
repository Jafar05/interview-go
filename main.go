package main

import "fmt"

func main() {
	src := []int{1, 2, 3, 4}
	dst := make([]int, 5) // Целевой срез длиной 5

	copy(dst, src) // Копируем из src в dst
	fmt.Println("Source:", src)
	fmt.Println("Destination:", dst)
}
