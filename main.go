package main

import (
	"sync"
)

func main() {
	jobs := make(chan int)
	res := make(chan int)
	workerPool(10, square, jobs, res)
}

func workerPool(id int, f func(int) int, jobs chan int, result chan int) {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		for val := range jobs {
			result <- f(val)
		}
	}()
	wg.Wait()
}

func square(num int) int {
	return num * num
}
