package main

func main() {
	// Объявляем переменную мапы, но не инициализируем ее
	var m map[string]int

	// Попытка записи в неинициализированную (nil) мапу
	m["key"] = 42 // Здесь произойдет runtime panic: "assignment to entry in nil map"
}
