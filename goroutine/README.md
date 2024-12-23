**Горутины** — это легковесные потоки выполнения в Go. Они позволяют выполнять функции параллельно или одновременно, но не создают полноценный поток (thread) на уровне операционной системы.

### Особенности горутин:
1. Легковесность:

- Горутины используют гораздо меньше памяти, чем потоки (стек горутины начинается с ~2 КБ).
- Можно запускать тысячи или даже миллионы горутин в одной программе.

2. Планирование:
- Горутины управляются планировщиком Go (runtime scheduler), который распределяет их выполнение между потоками операционной системы.
- Это абстрагирует управление потоками от разработчика.

3. Асинхронность:
- Горутины выполняются независимо от основного потока программы.
- Их работа не блокирует основную программу, если это не предусмотрено явно.

4. Запуск горутины:
- Для запуска достаточно перед функцией написать ключевое слово go.


#### Пример:
```go
package main

import (
    "fmt"
    "time"
)

func printNumbers() {
    for i := 1; i <= 5; i++ {
        fmt.Println(i)
        time.Sleep(500 * time.Millisecond)
    }
}

func main() {
    go printNumbers() // Запускаем функцию как горутину
    fmt.Println("Горутина запущена!")
    time.Sleep(3 * time.Second) // Ждём завершения горутины
    fmt.Println("Программа завершена.")
}
```

#### Что происходит:
1. Горутина printNumbers выполняется параллельно с основной программой.
2. Планировщик Go автоматически распределяет выполнение между горутиной и основным потоком.

### Преимущества горутин:
- Простота в использовании.
- Высокая производительность при работе с параллельными задачами.
- Эффективное использование системных ресурсов.

### Когда использовать:
- Для выполнения задач, которые можно параллелить (например, обработка запросов, выполнение вычислений).
- Для асинхронных операций, таких как работа с сетью или ввод/вывод.

---

**Изолированные горутины** — это горутины, которые не взаимодействуют напрямую с другими горутинами или общими данными. Они обычно выполняют свои задачи в полной изоляции, не влияя на состояние других горутин.

### Ключевые особенности:
- Минимальное использование общих ресурсов: Изолированные горутины не делят между собой данные или ресурсы, что минимизирует риски ошибок синхронизации (например, гонок данных).
- Безопасность: Так как они не используют общие данные, они исключают необходимость в синхронизации с помощью мьютексов или других механизмов блокировки.
- Примеры использования: Горутин, которые обрабатывают запросы в веб-сервисах, или выполняют задачи, не требующие взаимодействия с другими частями программы.

#### В Go управление выделением процессорных ядер осуществляется с помощью пакета runtime, в частности функции runtime.GOMAXPROCS.