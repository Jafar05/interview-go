package main

func main() {

}

// На вход подается канал значения из которого нужно перекладывать в оба канала
// Например если в канал записывается 1, 3, 5, то в канал a должно записаться 1, 3, 5 и канал b  должно записаться 1, 3, 5

func tee(in <-chan int) (chan any, chan any) {
	a := make(chan any, 1)
	b := make(chan any, 1)

	go func() {
		defer close(a)
		defer close(b)
		for val := range in {
			a <- val
			b <- val
		}
	}()

	return a, b
}
