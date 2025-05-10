package tasks

import "fmt"

// Функция zip объединяет два слайса
func zip(slice1, slice2 []int) [][2]int {
	// Определяем минимальную длину
	minLength := len(slice1)
	if len(slice2) < minLength {
		minLength = len(slice2)
	}

	// Создаем новый слайс для результата
	result := make([][2]int, minLength)

	// Заполняем результат парами значений
	for i := 0; i < minLength; i++ {
		result[i] = [2]int{slice1[i], slice2[i]}
	}

	return result
}

func main() {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5, 6, 7}

	zipped := zip(s1, s2)
	fmt.Println(zipped) // Вывод: [[1 4] [2 5] [3 6]]
}
