package Задачи

import "fmt"

func PipelineOfNumbers() {
	in := make(chan int)
	out := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			in <- i
		}
		close(in)
	}()

	go func() {
		defer close(out)
		for {
			num, ok := <-in
			if !ok {
				return
			}
			out <- num * num
		}
	}()

	for num := range out {
		fmt.Println(num)
	}
}
