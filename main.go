package main

import "fmt"

func main() {
	a := "Hello"
	fmt.Printf("a - %x, %q\n", &a, a)

	a = "world"
	fmt.Printf("a - %x, %q\n", &a, a)

	b := a[0:2]
	fmt.Printf("b - %x, %q\n", &b, b)
}
