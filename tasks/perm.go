package tasks

import (
	"fmt"
)

func perm(arr []rune, l, r int) {
	if l == r {
		fmt.Println(string(arr))
	} else {
		for i := l; i <= r; i++ {
			arr[i], arr[l] = arr[l], arr[i]
			perm(arr, l+1, r)
			arr[i], arr[l] = arr[l], arr[i]
		}
	}
}

func ImportPerm() {
	str := []rune("abc")
	perm(str, 0, len(str)-1)
}

//func main() {
//	input := "abc"
//	ch := make(chan string)
//	wg := &sync.WaitGroup{}
//
//	wg.Add(1)
//	go worker(input, wg, ch)
//
//	go func() {
//		wg.Wait()
//		close(ch)
//	}()
//
//	for combination := range ch {
//		fmt.Println(combination)
//	}
//}
//
//func worker(str string, wg *sync.WaitGroup, ch chan<- string) {
//	defer wg.Done()
//	permute([]rune(str), 0, ch)
//}
//
//func permute(rune []rune, start int, ch chan<- string) {
//	if start == len(rune)-1 {
//		ch <- string(rune)
//		return
//	}
//
//	for i := start; i < len(rune); i++ {
//		rune[i], rune[start] = rune[start], rune[i]
//		permute(rune, start+1, ch)
//		rune[i], rune[start] = rune[start], rune[i]
//	}
//}
