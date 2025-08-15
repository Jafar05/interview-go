package main

import "fmt"

func main() {

}

type Custom interface {
	Say()
}

type Base struct {
	name string
}

type Child struct {
	lastName string
	*Base
}

func (b *Base) Say() {
	fmt.Println("Hello ", b.name)
}

func (b *Child) Say() {
	fmt.Println("Hello ", b.lastName)
}

func init() {
	b1 := &Base{
		name: "Parent",
	}
	c1 := &Child{Base: &Base{name: "Child"}, lastName: "Inherited"}

	arr := []Custom{b1, c1}

	for _, val := range arr {
		val.Say()
	}
}

func NewObject(t interface{}) interface{} {
	var newObject any
	if t == "Child" {
		newObject = &Child{}
	} else if t == "Base" {
		newObject = &Base{}
	}

	return newObject
}

func GenerateBase() {

}

/*
1 Сделать структуры Base и Child.
2 Структура Base должна содержать строковое поле name.
3 Структура Child должна содержать строковое поле lastName.
4 Сделать функцию Say у структуры Base, которая распечатывает на экране: Hello, %name!
5 Пронаследовать Child от Base.
6 Инициализировать экземпляр b1 Base.
- присвоить name значение Parent
7 Инициализировать экземпляр c1 Сhild.
- присвоить name значение Child
- присвоить lastName значение Inherited
8 Вызвать у обоих экземпляров метод Say.
9 Переопределить метод Say для стурктуры Child, чтобы он выводил на экран: Hello, %lastName %name!
10 Сделать массив, содержащий b1 и c1.
11 Вызвать Say у всех элементов массива из шага 10.
12 Сделать метод NewObject для создания экземпляров Base и Child в зависимости от входного параметра.
13 Написать юнит тесты для метода NewObject.
14 Заменить инициализацию переменных b1 и c1 на использование метода NewObject.
15 Сделать генератор объетов Base и Child такой, чтобы:
- объекты Base создавались в фоновом потоке с задержкой 1 секунда;
- объекты Child создавались в фоновом потоке с задержкой 2 секунды;
- общее время генерации объектов не превышало 11 секунд;
16 Сделать асинхронный обработчик сгенерированных объектов такой, чтобы:
- метод Say вызывался в порядке генерации объектов;
- не приводил к утечкам памяти;
*/
