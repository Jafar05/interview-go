package Задачи

import (
	"fmt"
	"math/big"
	"sync"
)

var (
	cache = make(map[int]int)
	mu    sync.RWMutex
)

func fibonacci(n int) int {
	if n < 2 {
		return n
	}

	mu.RLock()
	if val, ok := cache[n]; ok {
		mu.RUnlock()
		return val
	}
	mu.RUnlock()

	result := fibonacci(n-1) + fibonacci(n-2)

	mu.Lock()
	cache[n] = result
	mu.Unlock()

	return result
}

func fibonacciIterative(n int) *big.Int {
	if n < 2 {
		return big.NewInt(int64(n))
	}

	a := big.NewInt(0)
	b := big.NewInt(1)

	for i := 2; i <= n; i++ {
		a.Add(a, b)
		a, b = b, a
	}
	return b
}

func ImportFibonacci() {
	fmt.Println(fibonacci(6))
}
