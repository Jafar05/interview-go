package print_console

import "fmt"

type User struct {
	Name string
}

func MainPointer() {
	user := User{Name: "Олег"}
	fmt.Println("Имя до обновления:", user.Name) // Олег

	// В updateUser будет передавться переменная по значение.
	// Это будет просто копия структуры user
	updateUser(user)
	fmt.Println("Имя после обновления:", user.Name) // Олег
}

// Значение "u User" в updateUser будет копией со своим участком памяти
// "u User" будет локальный. Никак не связвн с user из функции main
func updateUser(u User) {
	u.Name = "Таненбаум"
	fmt.Println("Имя внутри функции [updateUser]", u.Name) // Таненбаум

	resetUser(&u)
	fmt.Println("Имя после вызова функции [resetUser] внутри функции [updateUser]", u.Name) // Таненбаум
}

func resetUser(u *User) {

	// Если писать так, то мы используем как раз тот участок памяти переменной
	// которую мы приняли в функцию resetUser
	//u.Name = "bla bla"

	// При такой записи "&User{Name: "Безымянный"}", мы перезатираем участок памяти
	// на который смотрела переменная u новым так как используем новую структуру &User{}
	u = &User{Name: "Безымянный"}
	fmt.Println("Имя внутри функции [resetUser]", u.Name) // Безымянный
}
