**Структура** в Go — это составной тип данных, который позволяет группировать разные данные (поля) в единое целое. Каждый элемент структуры называется полем и может иметь свой тип данных (например, int, string, bool и т.д.).

### Основные особенности структуры:
- Группировка различных типов данных: Структуры позволяют объединять разные типы данных в одну логическую единицу.
- Определение полей: Каждое поле в структуре имеет имя и тип.
- Поле структуры может быть любым типом данных, включая другие структуры, массивы, слайсы и даже функции.

### Синтаксис для определения структуры:
```go
type Person struct {
    Name    string
    Age     int
    Address string
}
```

#### В данном примере мы определили структуру Person, которая содержит три поля:

- Name — строка (string),
- Age — целое число (int),
- Address — строка (string).

#### Пример создания и использования структуры:
```go
package main

import "fmt"

type Person struct {
    Name    string
    Age     int
    Address string
}

func main() {
    // Создание экземпляра структуры
    p := Person{
        Name:    "Alice",
        Age:     30,
        Address: "123 Main St",
    }

    // Доступ к полям структуры
    fmt.Println(p.Name)    // Вывод: Alice
    fmt.Println(p.Age)     // Вывод: 30
    fmt.Println(p.Address) // Вывод: 123 Main St
}
```
### Важные моменты:
1. Именованные поля: Каждое поле в структуре имеет имя, которое используется для доступа к его значению.

2. Создание экземпляра структуры:
- Вы можете создать экземпляр структуры, указав значения для ее полей, как это показано в примере выше.
- Если вы хотите создать структуру с нулевыми значениями, можно просто объявить переменную структуры: var p Person.

3. Типы полей: Поля могут быть любых типов, включая примитивные типы, срезы, карты и другие структуры. Структуры в Go могут быть вложенными.

4. Указатели на структуру: Структуры могут быть переданы в функции по значению или по указателю. Это важно для производительности, так как структуры могут занимать много памяти.

#### Пример передачи структуры по указателю:
```go
func changeName(p *Person) {
    p.Name = "Bob" // Изменяет значение Name через указатель
}

func main() {
    p := Person{Name: "Alice"}
    changeName(&p)
    fmt.Println(p.Name) // Вывод: Bob
}
```

### Применение структур:
- Моделирование объектов: Структуры часто используются для моделирования объектов реального мира, таких как люди, продукты, заказы и т.д.
- Возврат нескольких значений из функций: Структуры удобно использовать для возврата нескольких значений из функции.
- Организация данных: Структуры помогают организовать и группировать связанные данные.

### Заключение:
Структура в Go — это способ упаковки различных типов данных в один единый блок. Это мощный инструмент для работы с данными, который позволяет легко представлять сложные объекты и манипулировать ими.

---

## zero value
- Zero value: структура, где каждый элемент имеет своё zero value.
- Например, для числовых полей будет 0, для строк — пустая строка, для указателей — nil.