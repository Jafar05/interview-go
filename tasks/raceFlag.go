package tasks

import (
	"fmt"
	"time"
)

func RaceFlag() {
	left := make(chan struct{})
	right := make(chan struct{})

	go func() {
		for {
			<-left
			time.Sleep(time.Millisecond * 500)
			fmt.Println("left")
			right <- struct{}{}
		}
	}()

	go func() {
		for {
			<-right
			time.Sleep(time.Millisecond * 500)
			fmt.Println("right")
			left <- struct{}{}
		}
	}()

	left <- struct{}{}
	select {}
}
