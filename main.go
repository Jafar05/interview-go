package main

import "fmt"

func main() {
	sl := make([]int, 0, 3)

	sl = append(sl, 1, 2)
	addSl(sl)

	fmt.Println(sl)
}

func addSl(sl []int) {
	sl = append(sl, 3)
}
