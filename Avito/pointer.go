package Avito

import "fmt"

type Person struct {
	Name string
}

// В функции ChangeName ты передаёшь указатель person *Person,
// но внутри функции создаёшь новый объект Person и присваиваешь его локальной переменной person.
// Это затеняет исходный указатель, переданный в main, и не изменяет оригинальный объект.
func ChangeName(person *Person) {
	person = &Person{
		Name: "Alice",
	}
}

func pointer() {
	person := &Person{
		Name: "Bob",
	}

	fmt.Println(person.Name) // Bob
	ChangeName(person)
	fmt.Println(person.Name) // Bob
}

// Можно поменять так

// 1
/*
func ChangeName(person *Person) {
	person.Name = "Alice"
}
*/

// 2
/*
func ChangeName(person **Person) {
	*person = &Person{
		Name: "Alice",
	}
}

func main() {
	person := &Person{
		Name: "Bob",
	}

	fmt.Println(person.Name) // Bob
	ChangeName(&person)
	fmt.Println(person.Name) // Bob
}
*/

//3
/*
func ChangeName(person *Person) {
	*person = Person{
		Name: "Alice",
	}
}

func main() {
	person := Person{
		Name: "Bob",
	}

	fmt.Println(person.Name) // Bob
	ChangeName(&person)
	fmt.Println(person.Name) // Bob
}
*/
