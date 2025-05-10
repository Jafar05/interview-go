package tasks

import (
	"fmt"
	"sync"
	"time"
)

type Semaphore struct {
	sema chan struct{}
}

func NewSema(maxJobs int) *Semaphore {
	return &Semaphore{
		sema: make(chan struct{}, maxJobs),
	}
}

func (s *Semaphore) Inc() {
	s.sema <- struct{}{}
}

func (s *Semaphore) Dec() {
	<-s.sema
}

func semaphore() {
	const maxJobs = 3
	wg := &sync.WaitGroup{}
	s := NewSema(maxJobs)

	for i := 1; i <= 15; i++ {
		s.Inc()
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer func() { s.Dec() }()
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
