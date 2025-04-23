package Задачи

import "fmt"

func FindProgressCapacity() {
	slice := make([]int, 0) // Создаем пустой слайс
	prevCap := cap(slice)   // Предыдущее значение capacity

	for i := 0; i < 2000; i++ {
		slice = append(slice, i)
		currentCap := cap(slice) // Текущий capacity

		// Если capacity изменился, выводим информацию
		if currentCap != prevCap {
			growth := float64(currentCap-prevCap) / float64(prevCap) * 100
			fmt.Printf("Элементы: %d, Capacity: %d, Рост: %.2f%%\n", len(slice), currentCap, growth)
			prevCap = currentCap
		}
	}
}
