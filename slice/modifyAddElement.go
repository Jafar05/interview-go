package slice

import "fmt"

func modifyElement(slice []int) {
	slice[1] = 999
}

func addElement(slice []int) { // len = 3, cap = 3

	// Тут произошла алокация памяти, поэтому тут уже новый массив с len = 4, cap = 6
	slice = append(slice, 100) // len = 4, cap = 6
	slice[0] = 888
	fmt.Println("slice внутри addElement", slice) // 888, 999, 30, 100
}

func MainAddElement() {
	originSlice := []int{10, 20, 30}

	fmt.Println("До modifyElement", originSlice) // 10, 20, 30
	modifyElement(originSlice)

	// Значение поменяется на 10, 999, 30, так как в функции modifyElement у струтуры
	// которую туда передали(в modifyElement уже будет копия) будет ссылка на тот же массив
	// так как не было алокации памяти
	fmt.Println("После modifyElement", originSlice) // 10, 999, 30

	fmt.Println("До addElement", originSlice) // 10, 999, 30
	addElement(originSlice)
	fmt.Println("После addElement", originSlice) // 10, 999, 30
}
