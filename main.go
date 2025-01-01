package main

import "fmt"

func main() {
	var m map[string]int // Неинициализированная мапа (nil)

	// Попытка вставить элемент в неинициализированную мапу
	// Это приведет к панике!
	fmt.Println(m["apple"])
	m["apple"] = 5 // Паника: assignment to entry in nil map
	fmt.Println(m) // Выводит: map[]
}
