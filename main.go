package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {
	chat()
}

func msg(ch chan string) {
	message := []string{
		"Hello",
		"Hi",
		"How are you",
		"Fine",
	}

	for {
		ch <- message[rand.IntN(len(message))]
		time.Sleep(time.Second)
	}
	close(ch)
}

func chat() {
	ch := make(chan string)
	go msg(ch)
	for val := range ch {
		fmt.Println(val)
	}
}
