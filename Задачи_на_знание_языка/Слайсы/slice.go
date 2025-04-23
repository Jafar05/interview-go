package Слайсы

import "fmt"

func modifyArray(arr [3]int) {
	arr[0] = 10
	fmt.Println("Inside modifyArray", arr) // 10 2 3
}

func modifySlice(slice []int) {
	slice[0] = 10
	fmt.Println("Inside modifySlice", slice) // 10 2 3
}

func MainSlice() {
	array := [3]int{1, 2, 3}

	// Тут слайс(срез) создается из массива array
	// поинтер(указатель) смотрит на массив который
	// находится в MainSlice и изменения в modifyArray никак не
	// затрагивают этот массив
	slice := array[:]

	fmt.Println("Before modify array", array) // 1 2 3
	modifyArray(array)

	// Почему array не поменялся?
	// Так как в go все передается через копирование, то это значит,
	// что в функцию modifyArray мы передали копию массива из MainSlice.
	// В modifyArray лежит уже свой массив, который никак не влияет на массив в этой функции
	// Поэтому тут так и остается массив со значениями 1 2 3
	fmt.Println("After modify array", array) // 1 2 3

	fmt.Println("Before modify slice", slice) // 1 2 3
	modifySlice(slice)

	// Почему slice поменялся?
	// Слайс это структура у которой есть поинтер(указатель) на массив.
	// Так как в modifySlice мы передали копию slice, в modifySlice
	// создался свой слайс, но с поинтером(указателем) на тот же
	// участок памяти с массивом. Поэтому когда мы поменяли значене в
	// modifySlice, мы по сути поменяли значение у массива, на который смотрят
	// оба слайса и в modifySlice и в MainSlice, поэтому пока в modifySlice
	// capacity не превышает исходный массив(а иначе там создасться новый массив и
	// слайс в modifySlice уже будет иметь поинтер(указатель) на новый массив
	// значения будут меняться в обоих функциях
	fmt.Println("After modify slice", array) // 10 2 3
	fmt.Println("Final array", array)        // 10 2 3
}
