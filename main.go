package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	b := [3]int{1, 2, 3}
	c := [3]int{1, 2, 4}

	fmt.Println(a == b) // true (все элементы равны)
	fmt.Println(a == c) // false (различие в последнем элементе)
}
