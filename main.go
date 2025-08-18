package main

import "fmt"

func main() {
	s := "abcdef"
	r := []rune(s)
	sub := string(r[2:5])
	fmt.Println(sub) // cde
}
