package slice

import "fmt"

func ModifySlice() {
	data := []int{10, 20, 30, 40}

	fmt.Println("Изначальный слайс: ", data)     // 10, 20, 30, 40
	modify(data[:2])                             // 10, 20 индексы [0, 1] - не включая 2
	fmt.Println("Слайс после изменений: ", data) // 10, 20, 50, 60

}

func modify(slice []int) { // 10, 20, len(2), cap(4)
	slice = append(slice, 50, 60)                       // 10, 20, 50, 60
	fmt.Println("Слайс в функции модификации: ", slice) // 10, 20, 50, 60
}
