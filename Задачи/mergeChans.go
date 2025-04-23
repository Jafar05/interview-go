package Задачи

import (
	"fmt"
	"sync"
)

func mergeChan(chs ...chan int) chan int {
	res := make(chan int)
	wg := sync.WaitGroup{}

	for _, ch := range chs {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for val := range ch {
				res <- val
			}
		}()
	}

	go func() {
		wg.Wait()
		close(res)
	}()

	return res
}

func ImportMergeChan() {
	a := make(chan int)
	go func() {
		for i := 0; i < 3; i++ {
			a <- i
		}
		close(a)
	}()

	b := make(chan int)
	go func() {
		for i := 3; i < 6; i++ {
			b <- i
		}
		close(b)
	}()

	res := mergeChan(a, b)

	for val := range res {
		fmt.Println(val)
	}
}
