package tasks

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context, wg *sync.WaitGroup, jobs <-chan int) {
	defer wg.Done()
	for {
		select {
		case job, ok := <-jobs:
			if !ok {
				return
			}
			urls(job)
		case <-ctx.Done():
			fmt.Println("Worker stopped")
			return
		}
	}
}

func workerPool() {
	jobs := make(chan int)
	wg := &sync.WaitGroup{}
	maxJobs := 3
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*19)
	defer cancel()

	for i := 0; i < maxJobs; i++ {
		wg.Add(1)
		go worker(ctx, wg, jobs)
	}

	go func() {
		for i := 1; i <= 15; i++ {
			select {
			case jobs <- i:

			case <-ctx.Done():
				fmt.Println("Stopped adding jobs")
				break
			}
		}
		close(jobs)
	}()

	wg.Wait()
}

func urls(id int) {
	time.Sleep(time.Millisecond * 500)
	fmt.Println(id)
}

func ImportWorkerPool() {
	workerPool()
}
