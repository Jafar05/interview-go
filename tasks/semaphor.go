package tasks

import (
	"fmt"
	"sync"
	"time"
)

func semaphore() {
	wg := &sync.WaitGroup{}
	sem := make(chan struct{}, 3)

	for i := 1; i <= 15; i++ {
		sem <- struct{}{}
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer func() {
				<-sem
			}()
			urlsSema(i)
		}()
	}

	wg.Wait()
}

func urlsSema(id int) {
	time.Sleep(time.Millisecond * 500)
	fmt.Println(id)
}

func ImportSema() {
	semaphore()
}
