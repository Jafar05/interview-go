package tasks

import (
	"fmt"
	"sync"
)

func splitChanks(sl []int, chIn chan []int) {
	sizeChunks := 3
	for start := 0; start < len(sl); start += sizeChunks {
		end := start + sizeChunks
		if end > len(sl) {
			end = len(sl)
		}

		chIn <- sl[start:end]
	}
	close(chIn)
}

func workerpool(wg *sync.WaitGroup, chIn chan []int) {
	defer wg.Done()
	for val := range chIn {
		for _, nums := range val {
			fmt.Println(nums * 2)
		}
	}
}

func main() {
	sl := []int{1, 2, 3, 4, 5, 6, 7}

	chIn := make(chan []int)
	go splitChanks(sl, chIn)

	wg := &sync.WaitGroup{}
	maxGo := 3

	for i := 0; i < maxGo; i++ {
		wg.Add(1)
		go workerpool(wg, chIn)
	}

	wg.Wait()

}
