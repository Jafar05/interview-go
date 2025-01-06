package main

import "fmt"

func main() {
	// Инициализация массива чисел
	arr := [7]int{6, 8, 3, 5, 12, 9, 4}

	// Перебор элементов массива
	for _, num := range arr {
		if num > 10 {
			// Выводим число и выходим из цикла
			fmt.Println("Число больше 10 найдено:", num)
			break
		}
		fmt.Println("Текущее число:", num)
	}
}
