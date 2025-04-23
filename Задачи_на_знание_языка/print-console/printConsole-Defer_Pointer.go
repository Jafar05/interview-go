package print_console

import "fmt"

type Account struct {
	Balance int
}

func MainDeferPointer() {
	initialBalance := 1000
	account := &Account{Balance: initialBalance}

	defer printBalance("Изначальный баланс", account.Balance) // 1000
	defer printBalance("Текущий баланс", account.Balance)     // 1000

	// Так как тут была скопирована копия указателя, defer printAccountBalance
	// на время регистрации имеет defer указатель на область памяти переменной
	// выше, а не переназначенной переменной account ниже
	defer printAccountBalance("Указатель на баланс", account) // 1300

	account.Balance += 500           // 1500
	updateBalance(account, 200)      // 1300
	account = &Account{Balance: 300} // Переназначаем указатель на новый аккаунт
}

func updateBalance(acc *Account, amount int) {
	acc.Balance -= amount
}

func printBalance(label string, balance int) {
	fmt.Printf("%s: %d\n", label, balance)
}

func printAccountBalance(label string, acc *Account) {
	fmt.Printf("%s: %d\n", label, acc.Balance)
}

// !ВАЖНО:
// В go все копируется, даже указатели копируются (в функции и тд)
// при копировании указателя, несколько переменных могут смотреть на одну область памяти
// указатель в переменной можно перетереть, поменять на другой указатель, с другой областью в памяти
