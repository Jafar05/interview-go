package Avito

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Есть функция, работающая неопределённо долго и возвращающая число.
// Её тело нельзя изменять (представим, что внутри сетевой запрос).
func unpredictableFunc() int64 {
	rnd := rand.Int63n(5000)
	time.Sleep(time.Duration(rnd) * time.Millisecond)
	return rnd
}

// Нужно изменить функцию обёртку, которая будет работать с заданным таймаутом (например, 1 секунду).
// Если "длинная" функция отработала за это время — отлично, возвращаем результат.
// Если нет — возвращаем ошибку. Результат работы в этом случае нам не важен.
//
// Дополнительно нужно измерить, сколько выполнялась эта функция (просто вывести в лог).
// Сигнатуру функцию обёртки менять можно.
func predictableFunc() int64 {
	return unpredictableFunc()
}

// Решение

//func predictableFunc() int64 {
//	resultChan := make(chan int64, 1)
//	start := time.Now()
//
//	go func() {
//		result := unpredictableFunc()
//		resultChan <- result
//	}()
//
//	select {
//	case r := <-resultChan:
//		fmt.Println("chan res", time.Since(start))
//		return r
//	case <-time.After(1 * time.Second):
//		fmt.Println(time.Since(start))
//		return -1
//	}
//}

// Решение 2
//func predictableFunc(ctx context.Context) (int64, error) {
//	res := make(chan int64, 1)
//
//	go func() {
//		res <- unpredictableFunc()
//		close(res)
//	}()
//
//	select {
//	case <-ctx.Done():
//		return -1, fmt.Errorf("timeout error")
//	case val := <-res:
//		return val, nil
//	}
//}
//
//func main() {
//	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
//	defer cancel()
//
//	num, err := predictableFunc(ctx)
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	fmt.Println(num)
//}
